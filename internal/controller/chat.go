package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/utils"
)

type ChatControllerInterface interface {
	Send(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
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

func (c ChatController) Send(rw http.ResponseWriter, r *http.Request) {
	var chatPayload models.PostChatPayload

	err := json.NewDecoder(r.Body).Decode(&chatPayload)
	if err != nil {
		c.l.Error("error in unmarshalling postchat payload", err)
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	chat, err := c.services.Chat.Send(chatPayload.Message, chatPayload.RoomId, chatPayload.UserId)
	if err != nil {
		c.l.Error("unable to send chat message", err)
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendHttpResponse(rw, http.StatusOK, chat)

}

func (c ChatController) GetAll(rw http.ResponseWriter, r *http.Request) {

	chat, err := c.services.Chat.GetAll()
	if err != nil {
		c.l.Error("unable to send chat message", err)
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendHttpResponse(rw, http.StatusOK, chat)

}
