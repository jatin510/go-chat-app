package utils

import (
	"log/slog"
	"os"

	"github.com/jatin510/go-chat-app/internal/models"
)

type logger struct {
	l *slog.Logger
}

func NewLogger() models.Logger {
	opts := slog.HandlerOptions{}
	l := slog.New(slog.NewTextHandler(os.Stdout, &opts))
	return logger{l: l}
}

func (log logger) Info(m string, args ...any) {
	log.l.Info(m, args...)
}

func (log logger) Debug(m string, args ...any) {
	log.l.Debug(m, args...)
}

func (log logger) Error(m string, args ...any) {
	log.l.Error(m, args...)
}
