package models

import (
	"github.com/google/uuid"
)

type PostChatPayload struct {
	Message string    `json:"message"`
	RoomId  uuid.UUID `json:"roomId"`
	UserId  uuid.UUID `json:"userId"`
}
