package controller

import (
	"encoding/json"
	"net/http"

	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/utils"
)

type UserControllerInterface interface {
	Create(http.ResponseWriter, *http.Request)
	Login(http.ResponseWriter, *http.Request)
}

type UserController struct {
	l        models.Logger
	services *services.Services
}

func NewUserController(services *services.Services, l models.Logger) UserControllerInterface {
	return &UserController{
		l:        l,
		services: services,
	}
}

func (c UserController) Create(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		// handle error
		c.l.Error("error in unmarshalling create user payload", err)
		utils.SendHttpResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	u, err = c.services.User.Create(u.Name, u.Email, u.Password)
	if err != nil {
		// handle error
		c.l.Error("error in creating user", err)
		utils.SendHttpResponse(w, http.StatusInternalServerError, err.Error())
		return
	}
	utils.SendHttpResponse(w, http.StatusCreated, u)
}

func (c UserController) Login(w http.ResponseWriter, r *http.Request) {
	var u models.User

	err := json.NewDecoder(r.Body).Decode(&u)
	if err != nil {
		// handle error
		c.l.Error("error in unmarshalling login user payload", err)
		utils.SendHttpResponse(w, http.StatusInternalServerError, err.Error())
		return
	}

	u, err = c.services.User.Login(u.Email, u.Password)
	if err != nil {
		// TODO
		// handle error
		c.l.Error("error in creating user", err)
		if err.Error() == "invalid email or password" {
			utils.SendHttpResponse(w, http.StatusUnauthorized, err.Error())
		} else {
			utils.SendHttpResponse(w, http.StatusInternalServerError, err.Error())
		}
		return
	}
	utils.SendHttpResponse(w, http.StatusOK, u)
}
