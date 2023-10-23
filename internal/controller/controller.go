package controller

import (
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/utils/consumers/socket_consumer"
)

type Controllers struct {
	Chat ChatControllerInterface
	Room RoomControllerInterface
	User UserControllerInterface
}

func Init(services *services.Services, l models.Logger) *Controllers {
	return &Controllers{
		Chat: NewChatController(services, l),
		Room: NewRoomController(services, l),
		User: NewUserController(services, l),
	}
}

func (c *Controllers) InitSocketConsumer(socketConsumer socket_consumer.SocketConsumer) {
	c.Room.InitSocketConsumer(socketConsumer)
}
