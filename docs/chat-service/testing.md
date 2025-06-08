# Руководство по тестированию Chat Service

## Обзор

В этом документе описаны подходы и практики тестирования Chat Service. Мы используем различные типы тестов для обеспечения качества кода и функциональности.

## Типы тестов

### 1. Модульные тесты

#### Запуск
```bash
# Запуск всех модульных тестов
go test ./... -short

# Запуск тестов с покрытием
go test ./... -short -cover

# Запуск тестов конкретного пакета
go test ./controllers -short
```

#### Примеры

1. Тест контроллера чата:
```go
func TestChatController_GetChat(t *testing.T) {
    // Arrange
    ctrl := NewChatController(mockDB, mockCache)
    req := httptest.NewRequest("GET", "/api/chats/123", nil)
    w := httptest.NewRecorder()

    // Act
    ctrl.GetChat(w, req)

    // Assert
    assert.Equal(t, http.StatusOK, w.Code)
    // ... дополнительные проверки
}
```

2. Тест WebSocket обработчика:
```go
func TestMessageHandler_HandleMessage(t *testing.T) {
    // Arrange
    handler := NewMessageHandler(mockHub)
    msg := &Message{
        Type: "message",
        Data: map[string]interface{}{
            "chat_id": "123",
            "content": "test",
        },
    }

    // Act
    err := handler.HandleMessage(msg)

    // Assert
    assert.NoError(t, err)
    // ... дополнительные проверки
}
```

### 2. Интеграционные тесты

#### Запуск
```bash
# Запуск интеграционных тестов
go test ./... -tags=integration

# Запуск с покрытием
go test ./... -tags=integration -cover
```

#### Примеры

1. Тест создания чата:
```go
func TestIntegration_CreateChat(t *testing.T) {
    // Arrange
    app := setupTestApp()
    req := httptest.NewRequest("POST", "/api/chats", strings.NewReader(`{
        "name": "Test Chat",
        "type": "group",
        "participant_ids": ["1", "2"]
    }`))
    req.Header.Set("Authorization", "Bearer "+testToken)

    // Act
    resp, err := app.Test(req)

    // Assert
    assert.NoError(t, err)
    assert.Equal(t, http.StatusCreated, resp.StatusCode)
    // ... дополнительные проверки
}
```

2. Тест WebSocket соединения:
```go
func TestIntegration_WebSocketConnection(t *testing.T) {
    // Arrange
    app := setupTestApp()
    wsURL := "ws://localhost:8082/ws"
    ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)

    // Act & Assert
    assert.NoError(t, err)
    defer ws.Close()

    // Отправка сообщения
    err = ws.WriteJSON(map[string]interface{}{
        "type": "message",
        "data": map[string]interface{}{
            "chat_id": "123",
            "content": "test",
        },
    })
    assert.NoError(t, err)

    // Получение ответа
    var response map[string]interface{}
    err = ws.ReadJSON(&response)
    assert.NoError(t, err)
    assert.Equal(t, "message", response["type"])
}
```

### 3. Нагрузочные тесты

#### Запуск
```bash
# Запуск нагрузочных тестов
go test ./... -tags=load

# Запуск с параметрами
go test ./... -tags=load -v -timeout 30m
```

#### Примеры

1. Тест производительности WebSocket:
```go
func TestLoad_WebSocketPerformance(t *testing.T) {
    // Arrange
    app := setupTestApp()
    clients := make([]*websocket.Conn, 1000)
    messages := make(chan int, 10000)

    // Act
    for i := 0; i < 1000; i++ {
        ws, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
        assert.NoError(t, err)
        clients[i] = ws
        go sendMessages(ws, messages)
    }

    // Assert
    time.Sleep(30 * time.Second)
    assert.Greater(t, len(messages), 10000)
}
```

2. Тест производительности API:
```go
func TestLoad_APIPerformance(t *testing.T) {
    // Arrange
    app := setupTestApp()
    req := httptest.NewRequest("GET", "/api/chats", nil)
    req.Header.Set("Authorization", "Bearer "+testToken)

    // Act & Assert
    for i := 0; i < 1000; i++ {
        start := time.Now()
        resp, err := app.Test(req)
        duration := time.Since(start)

        assert.NoError(t, err)
        assert.Equal(t, http.StatusOK, resp.StatusCode)
        assert.Less(t, duration, 100*time.Millisecond)
    }
}
```

## Моки и стабы

### 1. Мок базы данных
```go
type MockDB struct {
    mock.Mock
}

func (m *MockDB) GetChat(id string) (*Chat, error) {
    args := m.Called(id)
    return args.Get(0).(*Chat), args.Error(1)
}
```

### 2. Мок кэша
```go
type MockCache struct {
    mock.Mock
}

func (m *MockCache) Get(key string) (interface{}, error) {
    args := m.Called(key)
    return args.Get(0), args.Error(1)
}
```

### 3. Мок WebSocket хаба
```go
type MockHub struct {
    mock.Mock
}

func (m *MockHub) Broadcast(msg *Message) error {
    args := m.Called(msg)
    return args.Error(0)
}
```

## Тестовое окружение

### 1. Docker Compose
```yaml
version: '3'
services:
  mongodb:
    image: mongo:4.4
    ports:
      - "27017:27017"
    volumes:
      - mongodb_data:/data/db

  redis:
    image: redis:6.0
    ports:
      - "6379:6379"
    volumes:
      - redis_data:/data

volumes:
  mongodb_data:
  redis_data:
```

### 2. Тестовые данные
```go
func setupTestData() {
    // Очистка базы данных
    db.DropDatabase()

    // Создание тестовых пользователей
    users := []User{
        {ID: "1", Username: "user1"},
        {ID: "2", Username: "user2"},
    }
    db.Users.InsertMany(users)

    // Создание тестовых чатов
    chats := []Chat{
        {ID: "1", Name: "Test Chat 1"},
        {ID: "2", Name: "Test Chat 2"},
    }
    db.Chats.InsertMany(chats)
}
```

## CI/CD интеграция

### 1. GitHub Actions
```yaml
name: Tests

on: [push, pull_request]

jobs:
  test:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v2
      - name: Set up Go
        uses: actions/setup-go@v2
        with:
          go-version: 1.19
      - name: Run tests
        run: |
          go test ./... -short
          go test ./... -tags=integration
          go test ./... -tags=load
```

### 2. Отчеты о покрытии
```bash
# Генерация отчета
go test ./... -coverprofile=coverage.out

# Просмотр отчета
go tool cover -html=coverage.out
```

## Рекомендации

### 1. Написание тестов
- Пишите тесты до реализации (TDD)
- Используйте табличные тесты
- Тестируйте граничные случаи
- Проверяйте обработку ошибок

### 2. Организация тестов
- Группируйте связанные тесты
- Используйте подтесты
- Добавляйте комментарии
- Следуйте соглашениям об именовании

### 3. Поддержка тестов
- Регулярно обновляйте тесты
- Удаляйте устаревшие тесты
- Рефакторите тестовый код
- Документируйте изменения 