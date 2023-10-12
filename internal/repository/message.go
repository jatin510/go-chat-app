package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jatin510/go-chat-app/internal/models"
)

type MessageRepoInterface interface {
	Create(models.Message) (models.Message, error)
	Update(models.Message) (models.Message, error)
	Delete(uuid.UUID) error
	FindOne(any) (models.Message, error)
	FindAll(any) ([]models.Message, error)
	Count(any) (int, error)
}

type message struct {
	db *pgx.Conn
	l  models.Logger
}

func NewMessageRepo(db *pgx.Conn, l models.Logger) MessageRepoInterface {
	return &message{
		db: db,
		l:  l,
	}
}

func (m message) Create(message models.Message) (models.Message, error) {
	return models.Message{}, nil
}

func (m message) Update(message models.Message) (models.Message, error) {
	return models.Message{}, nil
}

func (m message) Delete(id uuid.UUID) error {
	return nil
}

func (m message) FindOne(filter any) (models.Message, error) {
	return models.Message{}, nil
}

func (m message) FindAll(filter any) ([]models.Message, error) {
	return []models.Message{}, nil
}

func (m message) Count(filter any) (int, error) {
	return 0, nil
}
