-- +goose Up
-- Курсы (с фиксированными ID)
INSERT INTO courses (id, title, short_description, description, outcomes, audience, about_author, category, price, image_url, author_id)
VALUES 
(101, 'Основы маркетинга', 'Базовый курс по маркетингу', 'Описание курса 1', 'Навыки маркетинга', 'Новички', 'Redford School', 'Маркетинг', 1000.00, 'https://example.com/m1.jpg', 1),
(102, 'Продвинутый маркетинг', 'Углублённый курс по маркетингу', 'Описание курса 2', 'Глубокий анализ', 'Продвинутые', 'Redford School', 'Маркетинг', 1500.00, 'https://example.com/m2.jpg', 1),
(103, 'SMM-стратегии', 'Социальные сети и рост', 'Описание курса 3', 'Навыки продвижения в соцсетях', 'Для всех', 'Redford School', 'Маркетинг', 1200.00, 'https://example.com/m3.jpg', 1);

-- Главы (с фиксированными ID)
INSERT INTO chapters (id, course_id, title, position) VALUES
(201, 101, 'Введение', 1),
(202, 101, 'Анализ аудитории', 2),
(203, 101, 'Стратегия', 3),
(204, 102, 'Метрики эффективности', 1),
(205, 102, 'Работа с данными', 2),
(206, 103, 'Контент-планирование', 1),
(207, 103, 'Работа с платформами', 2);

-- Темы (теперь ссылаются на реальные chapter_id)
INSERT INTO topics (chapter_id, title, content, position) VALUES
(201, 'Что такое маркетинг?', 'Описание темы', 1),
(202, 'Сегментация', 'Описание темы', 1),
(203, 'Создание стратегии', 'Описание темы', 1),
(204, 'KPI и ROI', 'Описание темы', 1),
(205, 'Анализ данных', 'Описание темы', 1),
(206, 'Типы контента', 'Описание темы', 1),
(207, 'Instagram и TikTok', 'Описание темы', 1);

-- +goose Down
DELETE FROM topics WHERE chapter_id BETWEEN 201 AND 207;
DELETE FROM chapters WHERE id BETWEEN 201 AND 207;
DELETE FROM courses WHERE id BETWEEN 101 AND 103;