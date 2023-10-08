package repository

import "github.com/jatin510/go-chat-app/internal/models"

type MessageRepoInterface interface {
	create(models.Message) (models.Message, error)
	update(models.Message) (models.Message, error)
	delete(models.CID) error
	findOne(any) (models.Message, error)
	findAll(any) ([]models.Message, error)
	count(any) (int, error)
}

type message struct {
	db models.DBType
	l  models.Logger
}

func NewMessageRepo(db models.DBType, l models.Logger) MessageRepoInterface {
	return &message{
		db: db,
		l:  l,
	}
}

func (m message) create(message models.Message) (models.Message, error) {
	return models.Message{}, nil
}

func (m message) update(message models.Message) (models.Message, error) {
	return models.Message{}, nil
}

func (m message) delete(id models.CID) error {
	return nil
}

func (m message) findOne(filter any) (models.Message, error) {
	return models.Message{}, nil
}

func (m message) findAll(filter any) ([]models.Message, error) {
	return []models.Message{}, nil
}

func (m message) count(filter any) (int, error) {
	return 0, nil
}
