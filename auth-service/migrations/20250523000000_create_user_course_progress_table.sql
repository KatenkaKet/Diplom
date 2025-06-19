-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS user_course_progress (
    id SERIAL PRIMARY KEY,
    id_user INTEGER NOT NULL REFERENCES users(id) ON DELETE CASCADE,
    id_course INTEGER NOT NULL,
    done_theme_id INTEGER[] DEFAULT '{}',
    all_theme INTEGER NOT NULL DEFAULT 0,
    created_at TIMESTAMP DEFAULT now(),
    updated_at TIMESTAMP DEFAULT now()
);

-- Создаем индекс для быстрого поиска по id_user и id_course
CREATE INDEX  ON user_course_progress(id_user, id_course);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS user_course_progress;
-- +goose StatementEnd 