package config

import (
    "fmt"
    "os"
)

type Config struct {
    AppPort  string
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
    DSN      string
}

func LoadConfig() Config {
    cfg := Config{
        AppPort:  os.Getenv("PORT"),
        Host:     os.Getenv("DB_HOST"),
        Port:     os.Getenv("DB_PORT"),
        User:     os.Getenv("DB_USER"),
        Password: os.Getenv("DB_PASSWORD"),
        DBName:   os.Getenv("DB_NAME"),
    }

    cfg.DSN = fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
        cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName)

    return cfg
}
