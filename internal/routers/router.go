package router

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jatin510/go-chat-app/internal/controller"
	"github.com/jatin510/go-chat-app/internal/models"
)

func Init(controller *controller.Controllers, l models.Logger) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte("ok"))
	}).Methods("GET")

	// TODO
	// restRouter.Use(authMiddlware)

	_ = router.PathPrefix("/chat").Subrouter()
	userRouter := router.PathPrefix("/user").Subrouter()
	roomRouter := router.PathPrefix("/room").Subrouter()

	userRouter.HandleFunc("/signup", controller.User.Create).Methods(http.MethodPost)
	userRouter.HandleFunc("/login", controller.User.Login).Methods(http.MethodPost)

	roomRouter.HandleFunc("", controller.Room.Create).Methods(http.MethodPost)
	roomRouter.HandleFunc("", controller.Room.GetAll).Methods(http.MethodGet)
	roomRouter.HandleFunc("/{roomId}/join", controller.Room.Join).Methods(http.MethodPost)

	return router
}
