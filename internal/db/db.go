package db

import (
	"context"
	"log"
	"os"
	"path/filepath"
	"strings"

	crdbpgx "github.com/cockroachdb/cockroach-go/v2/crdb/crdbpgxv5"
	"github.com/jackc/pgx/v5"
	"github.com/jatin510/go-chat-app/internal/models"
)

func Init(l models.Logger) *pgx.Conn {
	ctx := context.Background()
	config, err := pgx.ParseConfig(os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatal(err)
	}
	config.Database = os.Getenv("DATABASE_NAME")
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

	if err != nil {
		panic(err)
	}

	l.Info("DB connected successfully")

	return conn
}

func initTable(ctx context.Context, tx pgx.Tx) error {

	path, err := filepath.Abs("./internal/db/init.sql")
	if err != nil {
		panic(err)
	}

	c, ioErr := os.ReadFile(path)
	if ioErr != nil {
		panic(ioErr)
	}
	sql := string(c)

	sql = strings.ReplaceAll(sql, "DB_NAME", os.Getenv("DATABASE_NAME"))

	_, err = tx.Exec(ctx, sql)
	if err != nil {
		panic(err)
	}

	return nil
}
