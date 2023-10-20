package services

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

type UserServiceInterface interface {
	Create(name string, email string, password string) (models.User, error)
	Login(email string, password string) (models.User, error)
}

type UserService struct {
	l    models.Logger
	repo *repository.Repository
}

func NewUserService(repo *repository.Repository, l models.Logger) UserServiceInterface {
	return &UserService{
		repo: repo,
		l:    l,
	}
}

func (u UserService) Create(name string, email string, password string) (models.User, error) {
	u.l.Info("creating a new user")

	currentTime := time.Now()
	user := models.User{
		ID:        uuid.New(),
		Name:      name,
		Email:     email,
		Password:  password,
		CreatedAt: currentTime,
		UpdatedAt: currentTime,
	}

	user, err := u.repo.User.Create(user)
	if err != nil {
		u.l.Error("error in create user ", err.Error())
		return models.User{}, err
	}

	return user, nil
}

func (u UserService) Login(email string, password string) (models.User, error) {
	u.l.Info("login a user")
	user, err := u.repo.User.FindOneByEmail(email)
	if err != nil {
		u.l.Error("error in login user ", err.Error())
		return models.User{}, err
	}

	if user.Password != password {
		return models.User{}, errors.New("invalid email or password")
	}

	user.Password = ""

	return user, nil
}
