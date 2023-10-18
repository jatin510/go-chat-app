package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
)

type RoomControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Join(http.ResponseWriter, *http.Request)
}

type RoomController struct {
	l        models.Logger
	services *services.Services
}

func NewRoomController(services *services.Services, l models.Logger) RoomControllerInterface {
	return &RoomController{
		l:        l,
		services: services,
	}
}

// type PostRoomPayload struct {
// 	Message string    `json:"message"`
// 	RoomId  uuid.UUID `json:"roomId"`
// 	UserId  uuid.UUID `json:"userId"`
// }

func (rc RoomController) Create(rw http.ResponseWriter, r *http.Request) {
	var room *models.Room

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		// handle error
		rc.l.Error("error in unmarshalling create room payload", err)
		// TODO write error
		// return err
	}

	// c.services.Chat.Send(p.Message, p.RoomId, p.UserId)
	rc.services.Room.Create(room.Name)
	rw.WriteHeader(http.StatusCreated)

	// return nil
}

func (rc RoomController) GetAll(rw http.ResponseWriter, r *http.Request) {
	// var p *PostChatPayload

	// err := json.NewDecoder(r.Body).Decode(&p)
	// if err != nil {
	// 	// handle error
	// 	rc.l.Error("error in unmarshalling postchat payload", err)
	// 	return err
	// }

	// c.services.Chat.Send(p.Message, p.RoomId, p.UserId)

	// return nil
}

func (rc RoomController) Join(rw http.ResponseWriter, r *http.Request) {
	// var p *PostChatPayload

	// err := json.NewDecoder(r.Body).Decode(&p)
	// if err != nil {
	// 	// handle error
	// 	rc.l.Error("error in unmarshalling postchat payload", err)
	// 	return err
	// }

	// c.services.Chat.Send(p.Message, p.RoomId, p.UserId)

	// return nil
}
