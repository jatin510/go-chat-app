package repository

import "github.com/jatin510/go-chat-app/internal/models"

type RoomRepoInterface interface {
	create(models.Room) (models.Room, error)
	update(models.Room) (models.Room, error)
	delete(models.CID) error
	findOne(any) (models.Room, error)
	findAll(any) ([]models.Room, error)
	count(any) (int, error)
}

type room struct {
	db models.DBType
	l  models.Logger
}

func NewRoomRepo(db models.DBType, l models.Logger) RoomRepoInterface {
	return &room{
		db: db,
		l:  l,
	}
}

func (r room) create(room models.Room) (models.Room, error) {
	return models.Room{}, nil
}

func (r room) update(room models.Room) (models.Room, error) {
	return models.Room{}, nil
}

func (r room) delete(id models.CID) error {
	return nil
}

func (r room) findOne(filter any) (models.Room, error) {
	return models.Room{}, nil
}

func (r room) findAll(filter any) ([]models.Room, error) {
	return []models.Room{}, nil
}

func (r room) count(filter any) (int, error) {
	return 0, nil
}
