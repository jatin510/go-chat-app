package services

import "github.com/jatin510/go-chat-app/internal/models"

type Chat interface {
	Send(string, models.CID, models.CID) error
}

type ChatService struct {
	l  models.Logger
	db *models.DBType
}

func NewChatService(l models.Logger, db *models.DBType) Chat {
	return &ChatService{
		l:  l,
		db: db,
	}
}

func (c ChatService) Send(m string, rId models.CID, userId models.CID) error {
	c.l.Info("Sending a message")
	return nil
}
