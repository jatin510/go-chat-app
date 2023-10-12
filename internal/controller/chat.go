package controller

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
)

type ChatControllerInterface interface {
	Send(http.ResponseWriter, *http.Request) error
}

type ChatController struct {
	l        models.Logger
	services *services.Services
}

func NewChatController(services *services.Services, l models.Logger) ChatControllerInterface {
	return &ChatController{
		l:        l,
		services: services,
	}
}

type PostChatPayload struct {
	Message string    `json:"message"`
	RoomId  uuid.UUID `json:"roomId"`
	UserId  uuid.UUID `json:"userId"`
}

func (c ChatController) Send(rw http.ResponseWriter, r *http.Request) error {
	var p *PostChatPayload

	err := json.NewDecoder(r.Body).Decode(&p)
	if err != nil {
		// handle error
		c.l.Error("error in unmarshalling postchat payload", err)
		return err
	}

	c.services.Chat.Send(p.Message, p.RoomId, p.UserId)

	return nil
}
