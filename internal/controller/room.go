package controller

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/utils"
	"github.com/jatin510/go-chat-app/internal/utils/consumers/socket_consumer"
)

type RoomControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	GetAll(http.ResponseWriter, *http.Request)
	Join(http.ResponseWriter, *http.Request)

	InitSocketConsumer(socket_consumer.SocketConsumer)
}

type RoomController struct {
	l              models.Logger
	services       *services.Services
	socketConsumer socket_consumer.SocketConsumer
}

func NewRoomController(services *services.Services, l models.Logger) RoomControllerInterface {
	return &RoomController{
		l:        l,
		services: services,
	}
}

func (rc *RoomController) InitSocketConsumer(socketConsumer socket_consumer.SocketConsumer) {
	rc.socketConsumer = socketConsumer
}

func (rc RoomController) Create(rw http.ResponseWriter, r *http.Request) {
	var room models.Room

	err := json.NewDecoder(r.Body).Decode(&room)
	if err != nil {
		// handle error
		rc.l.Error("error in unmarshalling create room payload", err)
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	room, err = rc.services.Room.Create(room.Name)
	if err != nil {
		rc.l.Error("error in create room ", err.Error())
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendHttpResponse(rw, http.StatusCreated, room)
}

func (rc RoomController) GetAll(rw http.ResponseWriter, r *http.Request) {

	rooms, err := rc.services.Room.GetAll()
	if err != nil {
		// handle error
		rc.l.Error("error in getting all rooms", err)
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	utils.SendHttpResponse(rw, http.StatusOK, rooms)
}

func (rc RoomController) Join(rw http.ResponseWriter, r *http.Request) {
	roomId := mux.Vars(r)["roomId"]
	userId := r.URL.Query().Get("userId")

	if !utils.IsUUID(roomId) {
		rc.l.Error("invalid room id")
		utils.SendHttpResponse(rw, http.StatusInternalServerError, "invalid room id")
		return
	}

	if !utils.IsUUID(userId) {
		rc.l.Error("invalid user id")
		utils.SendHttpResponse(rw, http.StatusInternalServerError, "invalid user id")
		return
	}

	roomUUID, err := uuid.Parse(roomId)
	if err != nil {
		rc.l.Error("unable to convert provided room id to UUID", err)
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	userUUID, err := uuid.Parse(userId)
	if err != nil {
		rc.l.Error("unable to convert provided user id to UUID", err)
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	sub, err := rc.services.Subscription.Create(roomUUID, userUUID)
	if err != nil {
		rc.l.Error("unable to create subscription", err)
		utils.SendHttpResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	write := make(chan interface{})
	p := models.Payload{
		Event: utils.JOIN_ROOM,
		Data: socket_consumer.JoinRoomPayload{
			UserId: sub.UserId,
			RoomId: sub.RoomId,
		},
		Write: write,
	}
	go func() {
		rc.socketConsumer.Job <- p
	}()

	e := <-write
	e = e.(error)
	fmt.Println(e)

	utils.SendHttpResponse(rw, http.StatusOK, sub)
}
