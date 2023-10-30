package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

type ChatServiceInterface interface {
	Send(msg string, roomId uuid.UUID, userId uuid.UUID) (models.Message, error)
	GetAll() ([]models.Message, error)
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

func (c ChatService) Send(m string, rId uuid.UUID, userId uuid.UUID) (models.Message, error) {
	// TODO
	// check if room exists
	// check if user exists

	currentTime := time.Now()

	message := models.Message{
		ID:        uuid.New(),
		Msg:       m,
		RoomId:    rId,
		UserId:    userId,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}
	msg, err := c.repo.Message.Create(message)
	if err != nil {
		c.l.Error("error in create message ", err.Error())
		return models.Message{}, err
	}

	return msg, nil
}

func (c ChatService) GetAll() ([]models.Message, error) {
	var filter = make(map[string]any)
	msgs, err := c.repo.Message.FindAll(filter)
	if err != nil {
		c.l.Error("error in create message ", err.Error())
		return []models.Message{}, err
	}

	return msgs, nil
}
