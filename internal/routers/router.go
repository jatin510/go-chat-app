package router

import (
	"github.com/gorilla/mux"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
)

func Init(services services.Services, l models.Logger) *mux.Router {
	router := mux.NewRouter()

	// TODO
	// restRouter.Use(authMiddlware)

	_ = router.PathPrefix("/chat").Subrouter()
	_ = router.PathPrefix("/user").Subrouter()
	_ = router.PathPrefix("/room").Subrouter()

	return router
}
