package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Text      string    `json:"text"`
	RoomId    int64     `json:"roomId"`
	User      uuid.UUID `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Room struct {
	ID        uuid.UUID   `json:"id"`
	Name      string      `json:"name"`
	Users     []uuid.UUID `json:"users"`
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
}
