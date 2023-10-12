package services

import (
	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

type RoomServiceInterface interface {
	GetAllRoomsByUserId(userId uuid.UUID) ([]models.Room, error)
}

type RoomService struct {
	l    models.Logger
	repo *repository.Repository
}

func NewRoomService(repo *repository.Repository, l models.Logger) RoomServiceInterface {
	return &RoomService{
		repo: repo,
		l:    l,
	}
}

func (c RoomService) GetAllRoomsByUserId(userId uuid.UUID) ([]models.Room, error) {
	c.l.Info("Fetching rooms of the user: " + userId.String())
	var rooms []models.Room

	rooms, err := c.repo.Room.FindAllRoomsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
