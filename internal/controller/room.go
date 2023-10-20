package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/utils"
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

func (rc RoomController) Create(rw http.ResponseWriter, r *http.Request) {
	var room models.Room

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		// handle error
		rc.l.Error("error in unmarshalling create room payload", err)
		// TODO write error
		// return err
	}

	room, err = rc.services.Room.Create(room.Name)
	if err != nil {
		rc.l.Error("error in create room ", err.Error())
	}

	utils.SendHttpResponse(rw, http.StatusCreated, room)
}

func (rc RoomController) GetAll(rw http.ResponseWriter, r *http.Request) {

	rooms, err := rc.services.Room.GetAll()
	if err != nil {
		// handle error
		rc.l.Error("error in getting all rooms", err)
		// TODO write error
		// return err
	}

	utils.SendHttpResponse(rw, http.StatusOK, rooms)
}

func (rc RoomController) Join(rw http.ResponseWriter, r *http.Request) {

}
