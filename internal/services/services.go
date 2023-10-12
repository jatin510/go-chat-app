package services

import (
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

type Services struct {
	Chat ChatServiceInterface
	Room RoomServiceInterface
}

func Init(repo *repository.Repository, l models.Logger) *Services {
	return &Services{
		Chat: NewChatService(repo, l),
		Room: NewRoomService(repo, l),
	}
}
