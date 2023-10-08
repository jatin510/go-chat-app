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

	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/utils"
)

var (
	DB models.DBType
	// DB_Client *mongo.Client
)

func init() {
	// DB_Client, DB = database.InitDatabase()
}

func main() {

	l := utils.NewLogger()

	// init repository
	repo := repository.Init(DB, l)

	// init services
	_ = services.Init(repo, l)

	// init controller
	// _ = controller.Init(services)

	// init router
	// router := route.InitRoute(services, l)

	port := "4000"

	s := http.Server{
		Addr:     ":" + port,
		ErrorLog: slog.NewLogLogger(slog.NewJSONHandler(os.Stderr, nil), slog.LevelError),
		// Handler:      router,
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
		// extra handling here
		cancel()
	}()

	if err := s.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown Failed:%+v", err)
	}
	log.Print("Server Exited Properly")
}
