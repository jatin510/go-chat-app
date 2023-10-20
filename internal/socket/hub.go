package socket

import (
	"errors"
	"fmt"

	"github.com/google/uuid"
)

// Hub maintains the set of active clients and broadcasts messages to the
// clients.
type Hub struct {
	// Registered clients.
	clients map[*Client]bool

	// Inbound messages from the clients.
	broadcast chan []byte

	// Register requests from the clients.
	register chan *Client

	// Unregister requests from clients.
	unregister chan *Client

	rooms map[uuid.UUID][]*Client
}

func newHub() *Hub {
	return &Hub{
		broadcast:  make(chan []byte),
		register:   make(chan *Client),
		unregister: make(chan *Client),
		clients:    make(map[*Client]bool),
		rooms:      make(map[uuid.UUID][]*Client),
	}
}

func (h *Hub) run() {
	for {
		select {
		case client := <-h.register:
			h.clients[client] = true
		case client := <-h.unregister:
			if _, ok := h.clients[client]; ok {
				delete(h.clients, client)
				close(client.send)
			}
		case message := <-h.broadcast:
			for client := range h.clients {
				select {
				case client.send <- message:
				default:
					close(client.send)
					delete(h.clients, client)
				}
			}
		}
	}
}

func (h *Hub) AddUserToRoom(userId uuid.UUID, roomId uuid.UUID) error {
	client, err := h.getClientByUserId(userId)
	if err != nil {
		return err
	}
	return h.insertClientInRoom(client, roomId)
}

func (h *Hub) insertClientInRoom(client *Client, roomId uuid.UUID) error {
	fmt.Println("insertClientInRoom called")
	room, ok := h.rooms[roomId]
	if !ok {
		h.rooms[roomId] = []*Client{}
		room = h.rooms[roomId]
	}
	fmt.Println("rooms before", h.rooms)

	for _, c := range room {
		if c == client {
			return errors.New("client not inserted, already in room")
		}
	}
	room = append(room, client)
	h.rooms[roomId] = room

	fmt.Println("hub clients: ", h.clients)
	fmt.Println("hub rooms: ", h.rooms)
	return nil
}

func (h *Hub) getClientByUserId(userId uuid.UUID) (*Client, error) {
	for c := range h.clients {
		if c.userId == userId {
			return c, nil
		}
	}
	return nil, errors.New("client not found for user ID: " + userId.String())
}
