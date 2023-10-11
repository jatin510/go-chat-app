package socket

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/websocket"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/utils"
)

const (
	// Time allowed to write a message to the peer.
	writeWait = 10 * time.Second

	// Time allowed to read the next pong message from the peer.
	pongWait = 60 * time.Second

	// Send pings to peer with this period. Must be less than pongWait.
	pingPeriod = (pongWait * 9) / 10

	// Maximum message size allowed from peer.
	maxMessageSize = 512
)

var (
	newline = []byte{'\n'}
	space   = []byte{' '}
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

// Client is a middleman between the websocket connection and the hub.
type Client struct {
	hub *Hub

	// The websocket connection.
	conn *websocket.Conn

	// Buffered channel of outbound messages.
	send chan []byte

	userId uuid.UUID
}

type Socket struct {
	hub      *Hub
	services *services.Services
	l        models.Logger
}

func NewSocket(hub *Hub, services *services.Services, l models.Logger) *Socket {
	return &Socket{
		hub:      hub,
		services: services,
		l:        l,
	}
}

// readPump pumps messages from the websocket connection to the hub.
//
// The application runs readPump in a per-connection goroutine. The application
// ensures that there is at most one reader on a connection by executing all
// reads from this goroutine.
func (c *Client) readPump() {
	defer func() {
		c.hub.unregister <- c
		c.conn.Close()
	}()
	c.conn.SetReadLimit(maxMessageSize)
	c.conn.SetReadDeadline(time.Now().Add(pongWait))
	c.conn.SetPongHandler(func(string) error { c.conn.SetReadDeadline(time.Now().Add(pongWait)); return nil })
	for {
		_, message, err := c.conn.ReadMessage()
		if err != nil {
			if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
				fmt.Printf("error: %v\n", err)
			}
			break
		}
		message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))
		c.hub.broadcast <- message
	}
}

// writePump pumps messages from the hub to the websocket connection.
//
// A goroutine running writePump is started for each connection. The
// application ensures that there is at most one writer to a connection by
// executing all writes from this goroutine.
func (c *Client) writePump() {
	ticker := time.NewTicker(pingPeriod)
	defer func() {
		ticker.Stop()
		c.conn.Close()
	}()
	for {
		select {
		case message, ok := <-c.send:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if !ok {
				// The hub closed the channel.
				c.conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			w, err := c.conn.NextWriter(websocket.TextMessage)
			if err != nil {
				return
			}
			w.Write(message)

			// Add queued chat messages to the current websocket message.
			n := len(c.send)
			for i := 0; i < n; i++ {
				w.Write(newline)
				w.Write(<-c.send)
			}

			if err := w.Close(); err != nil {
				return
			}
		case <-ticker.C:
			c.conn.SetWriteDeadline(time.Now().Add(writeWait))
			if err := c.conn.WriteMessage(websocket.PingMessage, nil); err != nil {
				return
			}
		}
	}
}

// serveWs handles websocket requests from the peer.
func (sck *Socket) ServeWs(w http.ResponseWriter, r *http.Request) {
	userId, err := sck.getUserIdFromRequest(r)
	if err != nil {
		// TODO: update to slog
		sck.l.Error(err.Error())
		w.WriteHeader(http.StatusUnauthorized)
		w.Write([]byte("Unauthorized"))
		return
	}

	// TODO: handle authentication
	// token := r.URL.Query().Get("token")

	rooms, err := sck.getUserRooms(userId)
	if err != nil {
		// TODO: update to slog
		sck.l.Error(err.Error())
	}

	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		sck.l.Error(err.Error())
		return
	}

	// user info from database
	// userId := getUserData(conn)

	client := &Client{hub: sck.hub, conn: conn, send: make(chan []byte, 256), userId: userId}
	client.hub.register <- client

	err = sck.insertClientInRooms(rooms, client)
	if err != nil {
		sck.l.Error(err.Error())
		return
	}

	// Allow collection of memory referenced by the caller by doing all work in
	// new goroutines.
	go client.writePump()
	go client.readPump()
}

func (sck *Socket) getUserIdFromRequest(r *http.Request) (uuid.UUID, error) {
	input := r.URL.Query().Get("userId")
	if !utils.IsUUID(input) {
		return uuid.UUID{}, errors.New("invalid UUID")
	}

	userId := uuid.UUID([]byte(input))
	if len(userId) == 0 {
		return uuid.UUID{}, errors.New("userId is required")
	}

	return userId, nil
}

func (sck *Socket) readMessage(conn *websocket.Conn) models.SocketMessage {
	_, message, err := conn.ReadMessage()
	if err != nil {
		if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
			sck.l.Info("error: %v", err)
		}
		// break
	}
	message = bytes.TrimSpace(bytes.Replace(message, newline, space, -1))

	var m models.SocketMessage
	err = json.Unmarshal(message, &m)
	if err != nil {
		sck.l.Error("error: %v", err)
	}

	fmt.Println("Messagge", m)
	return m
}

func (sck *Socket) getUserRooms(userId uuid.UUID) ([]models.Room, error) {
	rooms, err := sck.services.Room.GetAllRoomsByUserId(userId)
	if err != nil {
		return nil, err
	}

	return rooms, nil
}

func (sck *Socket) insertClientInRooms(rooms []models.Room, client *Client) error {
	for _, room := range rooms {
		err := sck.hub.insertClientInRoom(client, room.ID)
		if err != nil {
			sck.l.Error(err.Error())
		}
	}

	return nil
}
