package controllers

import (
	"chat-service/external"
	"context"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
)

type User struct {
	ID        int64  `json:"id"`
	Username  string `json:"username"`
	AvatarURL string `json:"avatar_url"`
}

func SearchUsers(c *gin.Context) {
	query := c.Query("query")
	if query == "*" {
		query = ":all"
	} else if query == ":all" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "not allowed"})
		return
	}

	if query != ":all" && len(strings.TrimSpace(query)) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "query too short"})
		return
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	var rows pgx.Rows
	var err error

	if query == ":all" {
		rows, err = external.AuthDB.Query(ctx,
			`SELECT id, username, avatar_url FROM users LIMIT 100`)
	} else {
		rows, err = external.AuthDB.Query(ctx,
			`SELECT id, username, avatar_url FROM users WHERE username ILIKE '%' || $1 || '%' LIMIT 10`,
			query)
	}

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "db error"})
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var u User
		if err := rows.Scan(&u.ID, &u.Username, &u.AvatarURL); err == nil {
			users = append(users, u)
		}
	}

	c.JSON(http.StatusOK, users)
}
