package external

import (
    "context"
    "fmt"
    "os"
    "time"

    "github.com/jackc/pgx/v5"
)

var AuthDB *pgx.Conn

func ConnectAuthDB() error {
    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
    defer cancel()

    dsn := os.Getenv("AUTH_DB_URI") // пример: postgres://user:pass@localhost:5432/auth_db
    conn, err := pgx.Connect(ctx, dsn)
    if err != nil {
        return fmt.Errorf("ошибка подключения к auth БД: %w", err)
    }

    AuthDB = conn
    fmt.Println("✅ Connected to auth-service database")
    return nil
}

type User struct {
    ID        int64
    Username  string
    AvatarURL string
}

func GetUserByID(userID int64) (*User, error) {
    ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
    defer cancel()

    var u User
    row := AuthDB.QueryRow(ctx,
        "SELECT id, username, avatar_url FROM users WHERE id = $1", userID)

    err := row.Scan(&u.ID, &u.Username, &u.AvatarURL)
    if err != nil {
        return nil, fmt.Errorf("пользователь не найден: %w", err)
    }

    return &u, nil
}
