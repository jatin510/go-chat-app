package main

import (
	"context"
	"log"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jatin510/go-chat-app/internal/controller"
	"github.com/jatin510/go-chat-app/internal/db"
	"github.com/jatin510/go-chat-app/internal/repository"
	router "github.com/jatin510/go-chat-app/internal/routers"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/socket"
	"github.com/jatin510/go-chat-app/internal/utils"
	"github.com/joho/godotenv"
	"github.com/rs/cors"
)

var (
	DB *pgx.Conn
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func main() {

	l := utils.NewLogger()

	DB = db.Init(l)

	// init repository
	repo := repository.Init(DB, l)

	// init services
	services := services.Init(repo, l)

	// init controller
	controllers := controller.Init(services, l)

	// init router
	router := router.Init(controllers, l)

	// init socket connection
	socket.Init(router, services, l)

	port := "4000"

	handler := cors.Default().Handler(router)

	s := http.Server{
		Addr:         ":" + port,
		ErrorLog:     slog.NewLogLogger(slog.NewJSONHandler(os.Stderr, nil), slog.LevelError),
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	done := make(chan os.Signal, 1)
	signal.Notify(done, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		if err := s.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	log.Print("Server Started on port ", port)
	<-done
	log.Print("Server Stopped")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer func() {
		// closing database connection
		DB.Close(context.Background())

		// closing context
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
