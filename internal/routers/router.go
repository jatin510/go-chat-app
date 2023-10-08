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
	_ = router.PathPrefix("/user").Subrouter()
	_ = router.PathPrefix("/room").Subrouter()

	return router
}
