# Установка и запуск Chat Service

## Требования
- Go 1.19 или выше
- MongoDB 4.4 или выше

## Установка

1. Перейдите в директорию chat-service:
```bash
cd chat-service
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте файл `.env` в корне chat-service директории:
```env
# Настройки приложения
APP_PORT=8082
APP_ENV=development

# Настройки MongoDB
MONGODB_URI=mongodb://localhost:27017
MONGODB_DB=chat_service

# JWT настройки
JWT_SECRET=your_jwt_secret

# WebSocket настройки
WS_PING_INTERVAL=30s
WS_PONG_WAIT=60s
WS_WRITE_WAIT=10s
```

4. Создайте базу данных в MongoDB:
```javascript
use chat_service
```

## Запуск

### Режим разработки
```bash
go run cmd/main.go
```
Сервис будет доступен по адресу: http://localhost:8082

### Сборка
```bash
go build -o chat-service cmd/main.go
```

### Запуск собранного бинарника
```bash
./chat-service
```

## Структура проекта

```
chat-service/
├── cmd/              # Точки входа
├── config/          # Конфигурация
├── controllers/     # HTTP контроллеры
├── database/        # Работа с БД
├── middleware/      # Middleware функции
├── models/          # Модели данных
├── routes/          # Маршруты
├── utils/           # Вспомогательные функции
└── ws/             # WebSocket обработчики
```

## API Endpoints

### Чаты
- `GET /api/chats` - Получение списка чатов
- `GET /api/chats/:id` - Получение информации о чате
- `POST /api/chats` - Создание нового чата
- `DELETE /api/chats/:id` - Удаление чата

### Сообщения
- `GET /api/chats/:id/messages` - Получение сообщений чата
- `POST /api/messages` - Отправка сообщения
- `DELETE /api/messages/:id` - Удаление сообщения

### Пользователи
- `GET /api/users/search` - Поиск пользователей

### WebSocket
- `ws://localhost:8082/ws` - WebSocket endpoint для real-time сообщений

## WebSocket Events

### Клиент -> Сервер
- `message` - Отправка сообщения
- `typing` - Индикатор набора текста
- `read` - Отметка о прочтении

### Сервер -> Клиент
- `message` - Новое сообщение
- `typing` - Индикатор набора текста
- `read` - Отметка о прочтении
- `error` - Ошибка

## Разработка

### Стиль кода
- Следуйте Go Code Review Comments
- Используйте gofmt для форматирования кода
- Пишите тесты для нового функционала

### Тестирование
```bash
# Запуск всех тестов
go test ./...

# Запуск тестов с покрытием
go test -cover ./...
```

### Линтинг
```bash
# Установка линтера
go install golang.org/x/lint/golint@latest

# Запуск линтера
golint ./...
```

## Деплой

### Подготовка к деплою
1. Обновите переменные окружения для продакшена
2. Соберите бинарник:
```bash
GOOS=linux GOARCH=amd64 go build -o chat-service cmd/main.go
```

### Деплой на сервер
1. Скопируйте бинарник и .env файл на сервер
2. Настройте systemd сервис:
```ini
[Unit]
Description=Chat Service
After=network.target mongodb.service

[Service]
Type=simple
User=app
WorkingDirectory=/opt/chat-service
ExecStart=/opt/chat-service/chat-service
Restart=always
Environment=APP_ENV=production

[Install]
WantedBy=multi-user.target
```

3. Запустите сервис:
```bash
sudo systemctl enable chat-service
sudo systemctl start chat-service
```

## Мониторинг

### Метрики
- Количество активных соединений
- Количество сообщений в секунду
- Время отклика API
- Использование памяти

### Логирование
- Уровни логирования: DEBUG, INFO, WARN, ERROR
- Формат: JSON
- Ротация логов: ежедневная
- Максимальный размер файла: 100MB 