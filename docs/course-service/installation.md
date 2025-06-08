# Установка и запуск Course Service

## Требования
- Go 1.19 или выше
- PostgreSQL 13 или выше

## Установка

1. Перейдите в директорию course-service:
```bash
cd course-service
```

2. Установите зависимости:
```bash
go mod download
```

3. Создайте файл `.env` в корне course-service директории:
```env
# Настройки приложения
APP_PORT=8081
APP_ENV=development

# Настройки базы данных
DB_HOST=localhost
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=your_password
DB_NAME=course_service

# Настройки файлового хранилища
UPLOAD_DIR=./uploads
MAX_FILE_SIZE=10485760 # 10MB
```

4. Создайте базу данных:
```sql
CREATE DATABASE course_service;
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
Сервис будет доступен по адресу: http://localhost:8081

### Сборка
```bash
go build -o course-service cmd/main.go
```

### Запуск собранного бинарника
```bash
./course-service
```

## Структура проекта

```
course-service/
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

### Курсы
- `GET /api/courses` - Получение списка курсов
- `GET /api/courses/:id` - Получение информации о курсе
- `POST /api/courses` - Создание нового курса
- `PUT /api/courses/:id` - Обновление курса
- `DELETE /api/courses/:id` - Удаление курса

### Уроки
- `GET /api/courses/:id/lessons` - Получение списка уроков курса
- `GET /api/lessons/:id` - Получение информации об уроке
- `POST /api/courses/:id/lessons` - Создание нового урока
- `PUT /api/lessons/:id` - Обновление урока
- `DELETE /api/lessons/:id` - Удаление урока

### Поиск
- `GET /api/courses/search` - Поиск по курсам
- `GET /api/lessons/search` - Поиск по урокам

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
2. Создайте директорию для загрузок:
```bash
mkdir -p /opt/course-service/uploads
```

3. Соберите бинарник:
```bash
GOOS=linux GOARCH=amd64 go build -o course-service cmd/main.go
```

### Деплой на сервер
1. Скопируйте бинарник и .env файл на сервер
2. Настройте systemd сервис:
```ini
[Unit]
Description=Course Service
After=network.target

[Service]
Type=simple
User=app
WorkingDirectory=/opt/course-service
ExecStart=/opt/course-service/course-service
Restart=always
Environment=APP_ENV=production

[Install]
WantedBy=multi-user.target
```

3. Запустите сервис:
```bash
sudo systemctl enable course-service
sudo systemctl start course-service
```

## Хранение файлов

### Структура директорий
```
uploads/
├── courses/         # Изображения курсов
├── lessons/         # Материалы уроков
└── temp/           # Временные файлы
```

### Ограничения
- Максимальный размер файла: 10MB
- Разрешенные форматы изображений: jpg, jpeg, png
- Разрешенные форматы документов: pdf, doc, docx 