package mocks

import (
	"time"

	"github.com/jatin510/go-chat-app/internal/repository"
)

const (
	DBQueryTimeout = 10 * time.Second
)

func Init() (*repository.Repository, MessageRepoMock) {
	message := NewMessageRepoMock()

	return &repository.Repository{
		Message: message,
		// User:         NewUserRepoMock(),
		// Room:         NewRoomRepoMock(),
		// Subscription: NewSubscriptionRepoMock(),
	}, message
}
