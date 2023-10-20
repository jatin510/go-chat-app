package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

type RoomServiceInterface interface {
	Create(name string) (models.Room, error)
	GetAllRoomsByUserId(userId uuid.UUID) ([]models.Room, error)
	GetAll() ([]models.Room, error)
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

func (r RoomService) Create(name string) (models.Room, error) {
	r.l.Info("Creating room " + name)

	// TODO: check if room already exists with given name

	t := time.Now()
	room := models.Room{
		ID:        uuid.New(),
		Name:      name,
		CreatedAt: t,
		UpdatedAt: t,
	}

	room, err := r.repo.Room.Create(room)
	if err != nil {
		r.l.Error("error in create room ", err.Error())
		return models.Room{}, err
	}

	return room, nil
}

func (r RoomService) GetAllRoomsByUserId(userId uuid.UUID) ([]models.Room, error) {
	r.l.Info("Fetching rooms of the user: " + userId.String())
	var rooms []models.Room

	rooms, err := r.repo.Room.FindAllRoomsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (r RoomService) GetAll() ([]models.Room, error) {
	var rooms []models.Room
	var filter = make(map[string]any)

	rooms, err := r.repo.Room.FindAll(filter)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
