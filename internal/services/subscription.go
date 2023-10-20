package services

import (
	"time"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

type SubscriptionInterface interface {
	Create(roomId uuid.UUID, userId uuid.UUID) (models.Subscription, error)
}

type SubscriptionService struct {
	l    models.Logger
	repo *repository.Repository
}

func NewSubscriptionService(repo *repository.Repository, l models.Logger) SubscriptionInterface {
	return &SubscriptionService{
		repo: repo,
		l:    l,
	}
}

func (r SubscriptionService) Create(roomId uuid.UUID, userId uuid.UUID) (models.Subscription, error) {
	r.l.Info("Creating subscription for room:" + roomId.String() + " and user: " + userId.String())

	// TODO: check if room already exists with given name

	t := time.Now()
	sub := models.Subscription{
		ID:        uuid.New(),
		RoomId:    roomId,
		UserId:    userId,
		CreatedAt: t,
		UpdatedAt: t,
	}

	sub, err := r.repo.Subscription.Create(sub)
	if err != nil {
		r.l.Error("error in create sub ", err.Error())
		return models.Subscription{}, err
	}

	return sub, nil
}
