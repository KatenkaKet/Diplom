# Установка и запуск Auth Service

## Требования
- Go 1.19 или выше
- PostgreSQL 13 или выше

## Установка

1. Перейдите в директорию auth-service:
```bash
cd auth-service
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте файл `.env` в корне auth-service директории:
```env
# Настройки приложения
APP_PORT=8080
APP_ENV=development

# Настройки базы данных
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=auth_service

# JWT настройки
JWT_SECRET=your_jwt_secret
JWT_EXPIRATION=24h
```

4. Создайте базу данных:
```sql
CREATE DATABASE auth_service;
```

5. Примените миграции:
```bash
go run cmd/migrate/main.go
```

## Запуск

### Режим разработки
```bash
go run cmd/main.go
```
Сервис будет доступен по адресу: http://localhost:8080

### Сборка
```bash
go build -o auth-service cmd/main.go
```

### Запуск собранного бинарника
```bash
./auth-service
```

## Структура проекта

```
auth-service/
├── cmd/              # Точки входа
├── config/          # Конфигурация
├── controllers/     # HTTP контроллеры
├── database/        # Работа с БД
├── middleware/      # Middleware функции
├── models/          # Модели данных
├── routes/          # Маршруты
└── utils/           # Вспомогательные функции
```

## API Endpoints

### Аутентификация
- `POST /api/auth/register` - Регистрация пользователя
- `POST /api/auth/login` - Вход в систему
- `POST /api/auth/logout` - Выход из системы
- `GET /api/auth/me` - Получение информации о текущем пользователе

### Пользователи
- `GET /api/users` - Получение списка пользователей
- `GET /api/users/:id` - Получение информации о пользователе
- `PUT /api/users/:id` - Обновление информации о пользователе
- `DELETE /api/users/:id` - Удаление пользователя

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
GOOS=linux GOARCH=amd64 go build -o auth-service cmd/main.go
```

### Деплой на сервер
1. Скопируйте бинарник и .env файл на сервер
2. Настройте systemd сервис:
```ini
[Unit]
Description=Auth Service
After=network.target

[Service]
Type=simple
User=app
WorkingDirectory=/opt/auth-service
ExecStart=/opt/auth-service/auth-service
Restart=always
Environment=APP_ENV=production

[Install]
WantedBy=multi-user.target
```

3. Запустите сервис:
```bash
sudo systemctl enable auth-service
sudo systemctl start auth-service
``` 