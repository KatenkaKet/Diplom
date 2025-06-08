# Chat Service API Documentation

## Общая информация

Базовый URL: `http://localhost:8082`

Все запросы должны содержать заголовок авторизации:
```
Authorization: Bearer <jwt_token>
```

## Endpoints

### Чаты

#### Получение списка чатов
```http
GET /api/chats
```

Query параметры:
- `page` (int, optional) - номер страницы (по умолчанию 1)
- `limit` (int, optional) - количество чатов на странице (по умолчанию 20)

Ответ:
```json
{
  "chats": [
    {
      "id": "string",
      "name": "string",
      "type": "string",
      "last_message": {
        "id": "string",
        "content": "string",
        "created_at": "string",
        "sender_id": "string"
      },
      "participants": [
        {
          "id": "string",
          "username": "string",
          "avatar": "string"
        }
      ],
      "created_at": "string",
      "updated_at": "string"
    }
  ],
  "total": 0,
  "page": 1,
  "limit": 20
}
```

#### Получение информации о чате
```http
GET /api/chats/{chat_id}
```

Ответ:
```json
{
  "id": "string",
  "name": "string",
  "type": "string",
  "participants": [
    {
      "id": "string",
      "username": "string",
      "avatar": "string"
    }
  ],
  "created_at": "string",
  "updated_at": "string"
}
```

#### Создание нового чата
```http
POST /api/chats
```

Тело запроса:
```json
{
  "name": "string",
  "type": "string",
  "participant_ids": ["string"]
}
```

Ответ:
```json
{
  "id": "string",
  "name": "string",
  "type": "string",
  "participants": [
    {
      "id": "string",
      "username": "string",
      "avatar": "string"
    }
  ],
  "created_at": "string",
  "updated_at": "string"
}
```

#### Удаление чата
```http
DELETE /api/chats/{chat_id}
```

Ответ:
```json
{
  "success": true
}
```

### Сообщения

#### Получение сообщений чата
```http
GET /api/chats/{chat_id}/messages
```

Query параметры:
- `page` (int, optional) - номер страницы (по умолчанию 1)
- `limit` (int, optional) - количество сообщений на странице (по умолчанию 50)
- `before` (string, optional) - получить сообщения до указанной даты (ISO 8601)

Ответ:
```json
{
  "messages": [
    {
      "id": "string",
      "content": "string",
      "chat_id": "string",
      "sender_id": "string",
      "created_at": "string",
      "updated_at": "string"
    }
  ],
  "total": 0,
  "page": 1,
  "limit": 50
}
```

#### Отправка сообщения
```http
POST /api/messages
```

Тело запроса:
```json
{
  "chat_id": "string",
  "content": "string"
}
```

Ответ:
```json
{
  "id": "string",
  "content": "string",
  "chat_id": "string",
  "sender_id": "string",
  "created_at": "string",
  "updated_at": "string"
}
```

#### Удаление сообщения
```http
DELETE /api/messages/{message_id}
```

Ответ:
```json
{
  "success": true
}
```

### Пользователи

#### Поиск пользователей
```http
GET /api/users/search
```

Query параметры:
- `query` (string, required) - поисковый запрос
- `page` (int, optional) - номер страницы (по умолчанию 1)
- `limit` (int, optional) - количество результатов на странице (по умолчанию 20)

Ответ:
```json
{
  "users": [
    {
      "id": "string",
      "username": "string",
      "avatar": "string"
    }
  ],
  "total": 0,
  "page": 1,
  "limit": 20
}
```

## WebSocket API

### Подключение
```
ws://localhost:8082/ws
```

Заголовки:
```
Authorization: Bearer <jwt_token>
```

### События

#### Отправка сообщения
```json
{
  "type": "message",
  "data": {
    "chat_id": "string",
    "content": "string"
  }
}
```

#### Индикатор набора текста
```json
{
  "type": "typing",
  "data": {
    "chat_id": "string",
    "is_typing": true
  }
}
```

#### Отметка о прочтении
```json
{
  "type": "read",
  "data": {
    "chat_id": "string",
    "message_id": "string"
  }
}
```

### Получение событий

#### Новое сообщение
```json
{
  "type": "message",
  "data": {
    "id": "string",
    "content": "string",
    "chat_id": "string",
    "sender_id": "string",
    "created_at": "string"
  }
}
```

#### Индикатор набора текста
```json
{
  "type": "typing",
  "data": {
    "chat_id": "string",
    "user_id": "string",
    "is_typing": true
  }
}
```

#### Отметка о прочтении
```json
{
  "type": "read",
  "data": {
    "chat_id": "string",
    "message_id": "string",
    "user_id": "string"
  }
}
```

#### Ошибка
```json
{
  "type": "error",
  "data": {
    "code": "string",
    "message": "string"
  }
}
```

## Коды ошибок

- `400` - Неверный запрос
- `401` - Не авторизован
- `403` - Доступ запрещен
- `404` - Ресурс не найден
- `500` - Внутренняя ошибка сервера

## Ограничения

- Максимальный размер сообщения: 4096 символов
- Максимальное количество участников в групповом чате: 100
- Максимальное количество сообщений в запросе: 100
- Таймаут WebSocket соединения: 60 секунд 