package models

import (
	"time"

	"github.com/google/uuid"
)

type Subscription struct {
	ID        uuid.UUID `json:"id"`
	RoomId    uuid.UUID `json:"roomId"`
	UserId    uuid.UUID `json:"userId"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
