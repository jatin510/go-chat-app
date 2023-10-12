package socket

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
)

func Init(router *mux.Router, services *services.Services, l models.Logger) {

	hub := newHub()
	go hub.run()

	socket := NewSocket(hub, services, l)

	// router.HandleFunc("/ws/register", func(w http.ResponseWriter, r *http.Request) {
	// 	// userId
	// 	serveWs(hub, w, r)
	// })

	// router.HandleFunc("/ws/chat", func(w http.ResponseWriter, r *http.Request) {
	// 	// userId, roomId, message
	// 	serveWs(hub, w, r)
	// })

	// router.HandleFunc("/ws/room", func(w http.ResponseWriter, r *http.Request) {
	// 	// userId, roomId, roomStatus
	// 	serveWs(hub, w, r)
	// })

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		socket.ServeWs(w, r)
	})
}
