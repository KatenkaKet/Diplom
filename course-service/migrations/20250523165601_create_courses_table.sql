-- +goose Up
-- +goose StatementBegin
CREATE TABLE courses (
    id SERIAL PRIMARY KEY,
    title TEXT NOT NULL,
    short_description TEXT,
    description TEXT,
    outcomes TEXT,
    audience TEXT,
    about_author TEXT,
    category TEXT,
    price NUMERIC(10, 2) DEFAULT 0.00,
    image_url TEXT,
    author_id INTEGER,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS courses;
-- +goose StatementEnd
