package models

import (
	"time"

	"github.com/google/uuid"
)

type Message struct {
	ID        uuid.UUID `json:"id"`
	Msg       string    `json:"msg"`
	RoomId    uuid.UUID `json:"roomId"`
	UserId    uuid.UUID `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
