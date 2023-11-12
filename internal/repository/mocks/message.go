package mocks

import (
	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/stretchr/testify/mock"
)

type MessageRepoMock struct {
	mock.Mock
}

func NewMessageRepoMock() MessageRepoMock {
	return MessageRepoMock{}
}

func (m MessageRepoMock) Create(message models.Message) (models.Message, error) {
	args := m.Called(message)
	return args.Get(0).(models.Message), args.Error(1)

}

func (m MessageRepoMock) Update(message models.Message) (models.Message, error) {
	args := m.Called(message)
	return args.Get(0).(models.Message), args.Error(1)
}

func (m MessageRepoMock) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m MessageRepoMock) FindOne(filter any) (models.Message, error) {
	args := m.Called(filter)
	return args.Get(0).(models.Message), args.Error(1)
}

func (m MessageRepoMock) FindAll(filter any) ([]models.Message, error) {
	args := m.Called(filter)
	return args.Get(0).([]models.Message), args.Error(1)
}

func (m MessageRepoMock) Count(filter any) (int, error) {
	args := m.Called(filter)
	return args.Int(0), args.Error(1)
}
