package repository

import (
	"time"

	"github.com/jackc/pgx/v4"
	"github.com/jatin510/go-chat-app/internal/models"
)

const (
	DBQueryTimeout = 10 * time.Second
)

type Repository struct {
	Message MessageRepoInterface
	User    UserRepoInterface
	Room    RoomRepoInterface
}

func Init(db *pgx.Conn, l models.Logger) *Repository {
	return &Repository{
		User:    NewUserRepo(db, l),
		Room:    NewRoomRepo(db, l),
		Message: NewMessageRepo(db, l),
	}
}
