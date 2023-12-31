package repository

import (
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jatin510/go-chat-app/internal/models"
)

const (
	DBQueryTimeout = 10 * time.Second
)

type Repository struct {
	Message      MessageRepoInterface
	User         UserRepoInterface
	Room         RoomRepoInterface
	Subscription SubscriptionRepoInterface
}

func Init(db *pgx.Conn, l models.Logger) *Repository {
	return &Repository{
		User:         NewUserRepo(db, l),
		Room:         NewRoomRepo(db, l),
		Message:      NewMessageRepo(db, l),
		Subscription: NewSubscriptionRepo(db, l),
	}
}
