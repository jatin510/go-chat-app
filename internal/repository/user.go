package repository

import (
	"github.com/google/uuid"
	"github.com/jackc/pgx/v4"
	"github.com/jatin510/go-chat-app/internal/models"
)

type UserRepoInterface interface {
	Create(models.User) (models.User, error)
	Update(models.User) (models.User, error)
	Delete(uuid.UUID) error
	FindOne(any) (models.User, error)
	FindAll(any) ([]models.User, error)
	Count(any) (int, error)
}

type user struct {
	db *pgx.Conn
	l  models.Logger
}

func NewUserRepo(db *pgx.Conn, l models.Logger) UserRepoInterface {
	return &user{
		db: db,
		l:  l,
	}
}

func (u user) Create(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (u user) Update(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (u user) Delete(id uuid.UUID) error {
	return nil
}

func (u user) FindOne(filter any) (models.User, error) {
	return models.User{}, nil
}

func (u user) FindAll(filter any) ([]models.User, error) {
	return []models.User{}, nil
}

func (u user) Count(filter any) (int, error) {
	return 0, nil
}
