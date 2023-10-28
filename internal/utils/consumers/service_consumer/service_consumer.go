package service_consumer

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/services"
	"github.com/jatin510/go-chat-app/internal/utils"
)

type ServiceConsumer struct {
	Job      chan models.Payload
	services *services.Services
}

func NewServiceConsumer(services *services.Services) *ServiceConsumer {
	return &ServiceConsumer{
		Job:      make(chan models.Payload),
		services: services,
	}
}

func (c ServiceConsumer) Init() {
	fmt.Println("InitServiceConsumer")

	for {
		select {
		case job := <-c.Job:
			fmt.Println("job: ", job)
			go c.handle(job.Event, job.Data, job.Write)
		}
	}
}

func (c *ServiceConsumer) handle(eventName string, payload interface{}, write chan interface{}) {
	switch eventName {
	case utils.GET_ALL_ROOMS_BY_USERID:
		userId := payload.(uuid.UUID)
		rooms, err := c.services.Room.GetAllRoomsByUserId(userId)
		if err != nil {
			fmt.Println("err: ", err)
		}
		write <- rooms
	}
}
