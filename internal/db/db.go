package db

import (
	"context"
	"os"

	"github.com/jackc/pgx/v4"
	"github.com/jatin510/go-chat-app/internal/models"
)

func Init(l models.Logger) *pgx.Conn {
	dsn := os.Getenv("DATABASE_URL")
	ctx := context.Background()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		l.Error("failed to connect database", err)
		panic(err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		l.Error("failed to ping database", err)
		panic(err)
	}

	l.Info("DB connected successfully")

	return conn
}
