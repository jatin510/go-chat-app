package repository

import "github.com/jatin510/go-chat-app/internal/models"

type UserRepoInterface interface {
	create(models.User) (models.User, error)
	update(models.User) (models.User, error)
	delete(models.CID) error
	findOne(any) (models.User, error)
	findAll(any) ([]models.User, error)
	count(any) (int, error)
}

type user struct {
	db models.DBType
	l  models.Logger
}

func NewUserRepo(db models.DBType, l models.Logger) UserRepoInterface {
	return &user{
		db: db,
		l:  l,
	}
}

func (u user) create(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (u user) update(user models.User) (models.User, error) {
	return models.User{}, nil
}

func (u user) delete(id models.CID) error {
	return nil
}

func (u user) findOne(filter any) (models.User, error) {
	return models.User{}, nil
}

func (u user) findAll(filter any) ([]models.User, error) {
	return []models.User{}, nil
}

func (u user) count(filter any) (int, error) {
	return 0, nil
}
