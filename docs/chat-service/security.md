# Руководство по безопасности Chat Service

## Обзор

В этом документе описаны меры безопасности, применяемые в Chat Service для защиты данных и обеспечения безопасной работы сервиса.

## Аутентификация

### JWT токены

1. Генерация токена:
```go
func generateToken(user *User) (string, error) {
    claims := jwt.MapClaims{
        "user_id": user.ID,
        "exp":     time.Now().Add(24 * time.Hour).Unix(),
        "iat":     time.Now().Unix(),
    }
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(config.JWTSecret))
}
```

2. Валидация токена:
```go
func validateToken(tokenString string) (*jwt.Token, error) {
    return jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
        }
        return []byte(config.JWTSecret), nil
    })
}
```

### Middleware аутентификации

```go
func AuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        token := r.Header.Get("Authorization")
        if token == "" {
            http.Error(w, "Unauthorized", http.StatusUnauthorized)
            return
        }

        token = strings.TrimPrefix(token, "Bearer ")
        if _, err := validateToken(token); err != nil {
            http.Error(w, "Invalid token", http.StatusUnauthorized)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

## Авторизация

### Проверка прав доступа

1. Проверка владельца чата:
```go
func checkChatOwnership(chatID, userID string) bool {
    chat, err := db.GetChat(chatID)
    if err != nil {
        return false
    }
    return chat.OwnerID == userID
}
```

2. Проверка участника чата:
```go
func checkChatParticipant(chatID, userID string) bool {
    chat, err := db.GetChat(chatID)
    if err != nil {
        return false
    }
    for _, p := range chat.Participants {
        if p.UserID == userID {
            return true
        }
    }
    return false
}
```

### Middleware авторизации

```go
func ChatAccessMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        userID := getUserIDFromContext(r.Context())
        chatID := chi.URLParam(r, "chatID")

        if !checkChatParticipant(chatID, userID) {
            http.Error(w, "Forbidden", http.StatusForbidden)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

## Защита данных

### Шифрование

1. Шифрование сообщений:
```go
func encryptMessage(content string) (string, error) {
    block, err := aes.NewCipher([]byte(config.EncryptionKey))
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return "", err
    }

    ciphertext := gcm.Seal(nonce, nonce, []byte(content), nil)
    return base64.StdEncoding.EncodeToString(ciphertext), nil
}
```

2. Расшифровка сообщений:
```go
func decryptMessage(encrypted string) (string, error) {
    data, err := base64.StdEncoding.DecodeString(encrypted)
    if err != nil {
        return "", err
    }

    block, err := aes.NewCipher([]byte(config.EncryptionKey))
    if err != nil {
        return "", err
    }

    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return "", err
    }

    nonceSize := gcm.NonceSize()
    if len(data) < nonceSize {
        return "", errors.New("ciphertext too short")
    }

    nonce, ciphertext := data[:nonceSize], data[nonceSize:]
    plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)
    if err != nil {
        return "", err
    }

    return string(plaintext), nil
}
```

### Хеширование

1. Хеширование паролей:
```go
func hashPassword(password string) (string, error) {
    bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
    return string(bytes), err
}
```

2. Проверка паролей:
```go
func checkPasswordHash(password, hash string) bool {
    err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
    return err == nil
}
```

## Защита от атак

### Rate Limiting

```go
func RateLimitMiddleware(next http.Handler) http.Handler {
    limiter := rate.NewLimiter(rate.Every(time.Second), 100)
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        if !limiter.Allow() {
            http.Error(w, "Too Many Requests", http.StatusTooManyRequests)
            return
        }
        next.ServeHTTP(w, r)
    })
}
```

### CORS

```go
func CORSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("Access-Control-Allow-Origin", config.AllowedOrigins)
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
        w.Header().Set("Access-Control-Max-Age", "86400")

        if r.Method == "OPTIONS" {
            w.WriteHeader(http.StatusOK)
            return
        }

        next.ServeHTTP(w, r)
    })
}
```

### XSS Protection

```go
func XSSMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.Header().Set("X-XSS-Protection", "1; mode=block")
        w.Header().Set("X-Content-Type-Options", "nosniff")
        w.Header().Set("X-Frame-Options", "DENY")
        next.ServeHTTP(w, r)
    })
}
```

## Безопасность WebSocket

### Проверка соединения

```go
func (h *Hub) handleConnection(conn *websocket.Conn) {
    // Проверка токена
    token := conn.Request().Header.Get("Authorization")
    if token == "" {
        conn.Close()
        return
    }

    // Валидация токена
    if _, err := validateToken(token); err != nil {
        conn.Close()
        return
    }

    // Установка таймаутов
    conn.SetReadDeadline(time.Now().Add(60 * time.Second))
    conn.SetPongHandler(func(string) error {
        conn.SetReadDeadline(time.Now().Add(60 * time.Second))
        return nil
    })

    // Обработка сообщений
    for {
        messageType, message, err := conn.ReadMessage()
        if err != nil {
            break
        }
        h.handleMessage(conn, messageType, message)
    }
}
```

### Защита от DoS

```go
func (h *Hub) handleMessage(conn *websocket.Conn, messageType int, message []byte) {
    // Проверка размера сообщения
    if len(message) > 4096 {
        conn.WriteJSON(map[string]string{
            "error": "Message too large",
        })
        return
    }

    // Проверка частоты сообщений
    if !h.rateLimiter.Allow() {
        conn.WriteJSON(map[string]string{
            "error": "Too many messages",
        })
        return
    }

    // Обработка сообщения
    h.processMessage(conn, messageType, message)
}
```

## Мониторинг безопасности

### Логирование

```go
func SecurityLogger(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        start := time.Now()
        next.ServeHTTP(w, r)
        duration := time.Since(start)

        logger.Info("security",
            zap.String("method", r.Method),
            zap.String("path", r.URL.Path),
            zap.String("ip", r.RemoteAddr),
            zap.Duration("duration", duration),
            zap.Int("status", w.(*responseWriter).status),
        )
    })
}
```

### Алерты

```go
func sendSecurityAlert(alert *SecurityAlert) error {
    // Отправка в Slack
    if err := slack.SendMessage(alert.Format()); err != nil {
        return err
    }

    // Отправка по email
    if err := email.SendAlert(alert); err != nil {
        return err
    }

    // Запись в лог
    logger.Error("security_alert",
        zap.String("type", alert.Type),
        zap.String("message", alert.Message),
        zap.String("severity", alert.Severity),
    )

    return nil
}
```

## Рекомендации

### 1. Общие рекомендации
- Регулярно обновляйте зависимости
- Используйте последние версии Go
- Следуйте принципу наименьших привилегий
- Регулярно проводите аудит безопасности

### 2. Конфигурация
- Храните секреты в переменных окружения
- Используйте разные ключи для разных окружений
- Ограничивайте доступ к конфигурации
- Регулярно ротируйте ключи

### 3. Мониторинг
- Настройте алерты на подозрительную активность
- Регулярно проверяйте логи
- Мониторьте использование ресурсов
- Отслеживайте попытки взлома 