package socket

import (
	"net/http"

	"github.com/gorilla/mux"
)

func Init(router *mux.Router) {

	hub := newHub()
	go hub.run()

	router.HandleFunc("/ws", func(w http.ResponseWriter, r *http.Request) {
		serveWs(hub, w, r)
	})
}
