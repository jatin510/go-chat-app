package db

import (
	"context"
	"fmt"
	"log"
	"log/slog"
	"os"
	"time"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jatin510/go-chat-app/internal/models"
)

func Init(l models.Logger) *pgx.Conn {
	ctx := context.Background()
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	config.RuntimeParams["application_name"] = "$ docs_simplecrud_gopgx"
	conn, err := pgx.ConnectConfig(context.Background(), config)
	if err != nil {
		log.Fatal(err)
	}
	// defer conn.Close(context.Background())

	if err != nil {
		l.Error("failed to connect database", err)
		panic(err)
	}

	err = conn.Ping(ctx)
	if err != nil {
		l.Error("failed to ping database", err)
		panic(err)
	}

	// Set up table
	err = crdbpgx.ExecuteTx(context.Background(), conn, pgx.TxOptions{}, func(tx pgx.Tx) error {
		return initTable(context.Background(), tx)
	})

	l.Info("DB connected successfully")

	return conn
}

func initTable(ctx context.Context, tx pgx.Tx) error {

	id, _ := uuid.Parse("1089114a-74df-4fd0-ae88-13e7669ea881")

	_, err := tx.Exec(ctx, "DROP TABLE IF EXISTS rooms")
	if err != nil {
		slog.Error("failed to drop rooms table", err)
		panic(err)
	}
	r, err := tx.Exec(ctx, "CREATE TABLE rooms (id UUID PRIMARY KEY, name TEXT NOT NULL, user_id uuid, created_at timestamptz NOT NULL, updated_at timestamp)")
	if err != nil {
		slog.Error("failed to create rooms table", err)
		panic(err)
	}
	fmt.Println(r)
	r, err = tx.Exec(ctx, "INSERT INTO rooms VALUES ($1,$2,$3,$4,$5)", id, "room1", uuid.New(), time.Now(), time.Now())
	if err != nil {
		slog.Error("failed to insert into rooms table", err)
		panic(err)
	}
	fmt.Println(r)

	return nil
}
