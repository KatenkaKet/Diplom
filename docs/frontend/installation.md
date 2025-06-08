# Установка и запуск Frontend

## Требования
- Node.js 16.x или выше
- npm 7.x или выше

## Установка

1. Перейдите в директорию frontend:
```bash
cd frontend
```

2. Установите зависимости:
```bash
npm install
```

3. Создайте файл `.env` в корне frontend директории:
```env
# API URLs
AUTH_SERVICE_URL=http://localhost:8080/api
COURSE_SERVICE_URL=http://localhost:8081/api
CHAT_SERVICE_URL=http://localhost:8082/api
```

## Запуск

### Режим разработки
```bash
npm run dev
```
Приложение будет доступно по адресу: http://localhost:3000

### Сборка для продакшена
```bash
npm run build
```

### Запуск продакшен версии
```bash
npm run start
```

## Структура проекта

```
frontend/
├── assets/          # Статические ресурсы
├── components/      # Vue компоненты
├── composables/     # Композабл функции
├── layouts/         # Шаблоны страниц
├── middleware/      # Middleware функции
├── pages/          # Страницы приложения
├── plugins/        # Nuxt плагины
├── public/         # Публичные файлы
├── stores/         # Pinia stores
└── types/          # TypeScript типы
```

## Основные компоненты

### Stores
- `auth.ts` - управление аутентификацией
- `chat.ts` - управление чатами
- `course.ts` - управление курсами

### Компоненты
- `CourseCard.vue` - карточка курса
- `ChatList.vue` - список чатов
- `DialogBox.vue` - окно диалога

## Разработка

### Стиль кода
- Используйте TypeScript для всех компонентов
- Следуйте Vue.js Style Guide
- Используйте Composition API

### Тестирование
```bash
# Запуск unit тестов
npm run test

# Запуск e2e тестов
npm run test:e2e
```

### Линтинг
```bash
# Проверка кода
npm run lint

# Автоматическое исправление
npm run lint:fix
```

## Деплой

### Подготовка к деплою
1. Обновите переменные окружения для продакшена
2. Соберите приложение:
```bash
npm run build
```

### Деплой на сервер
1. Скопируйте содержимое директории `.output` на сервер
2. Настройте nginx или другой веб-сервер
3. Запустите приложение:
```bash
node .output/server/index.mjs
``` 