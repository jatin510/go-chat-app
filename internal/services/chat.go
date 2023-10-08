package services

import (
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

type ChatServiceInterface interface {
	Send(msg string, roomId models.CID, userId models.CID) error
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

func (c ChatService) Send(m string, rId models.CID, userId models.CID) error {
	c.l.Info("Sending a message")
	return nil
}
