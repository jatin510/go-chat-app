package socket_consumer

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/socket"
	"github.com/jatin510/go-chat-app/internal/utils"
)

type SocketConsumer struct {
	Job    chan models.Payload
	socket *socket.Socket
}

func NewSocketConsumer(socket *socket.Socket) *SocketConsumer {
	return &SocketConsumer{
		Job:    make(chan models.Payload),
		socket: socket,
	}
}

func (c SocketConsumer) Init() {
	fmt.Println("InitSocketConsumer")

	for {
		select {
		case job := <-c.Job:
			fmt.Println("job: ", job)
			go c.handle(job.Event, job.Data, job.Write)
		}
	}
}

type JoinRoomPayload struct {
	UserId uuid.UUID
	RoomId uuid.UUID
}

func (c *SocketConsumer) handle(eventName string, payload interface{}, write chan interface{}) {
	switch eventName {
	case utils.JOIN_ROOM:
		p := payload.(JoinRoomPayload)
		err := c.socket.JoinRoom(p.UserId, p.RoomId)
		if err != nil {
			fmt.Println("err: ", err)
		}
		write <- err
	}
}
