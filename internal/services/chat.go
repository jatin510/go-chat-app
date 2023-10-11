package services

import (
	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

type ChatServiceInterface interface {
	Send(msg string, roomId uuid.UUID, userId uuid.UUID) error
}

type ChatService struct {
	l    models.Logger
	repo *repository.Repository
}

func NewChatService(repo *repository.Repository, l models.Logger) ChatServiceInterface {
	return &ChatService{
		repo: repo,
		l:    l,
	}
}

func (c ChatService) Send(m string, rId uuid.UUID, userId uuid.UUID) error {
	c.l.Info("Sending a message")
	return nil
}
