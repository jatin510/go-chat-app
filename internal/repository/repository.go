package repository

import "github.com/jatin510/go-chat-app/internal/models"

type Repository struct {
	Message MessageRepoInterface
	User    UserRepoInterface
	Room    RoomRepoInterface
}

func Init(db models.DBType, l models.Logger) *Repository {
	return &Repository{
		User:    NewUserRepo(db, l),
		Room:    NewRoomRepo(db, l),
		Message: NewMessageRepo(db, l),
	}
}
