package models

type SocketMessage struct {
	Event string      `json:"event"`
	Data  interface{} `json:"data"`
}
