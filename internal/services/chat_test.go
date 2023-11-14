package services

import (
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
	"github.com/jatin510/go-chat-app/internal/repository/mocks"
	"github.com/jatin510/go-chat-app/internal/utils"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestNewChatService(t *testing.T) {

	repo := &repository.Repository{}
	l := utils.NewLogger()

	type args struct {
		repo *repository.Repository
		l    models.Logger
	}
	tests := []struct {
		name string
		args args
		want ChatServiceInterface
	}{
		{
			name: "chat service initialization",
			args: args{
				repo: repo,
				l:    l,
			},
			want: &ChatService{
				repo: repo,
				l:    l,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewChatService(tt.args.repo, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewChatService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestChatService_Send(t *testing.T) {

	repo, messageRepo := mocks.Init()

	l := utils.NewLogger()

	commonUUID := uuid.New()

	type SendReturnType struct {
		message models.Message
		err     error
	}

	type fields struct {
		l    models.Logger
		repo *repository.Repository
	}
	type args struct {
		m      string
		rId    uuid.UUID
		userId uuid.UUID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    SendReturnType
		wantErr bool
	}{
		// passing case
		{
			name: "message sent",
			fields: fields{
				l:    l,
				repo: repo,
			},
			args: args{
				m:      "hello",
				rId:    commonUUID,
				userId: commonUUID,
			},
			want: SendReturnType{
				message: models.Message{
					Msg:    "hello",
					RoomId: commonUUID,
					UserId: commonUUID,
				},
			},
			wantErr: false,
		},
		// error case
		{
			name: "message sent failed",
			fields: fields{
				l:    l,
				repo: repo,
			},
			args: args{
				m:      "hello",
				rId:    commonUUID,
				userId: commonUUID,
			},
			want: SendReturnType{
				message: models.Message{
					Msg:    "hello",
					RoomId: commonUUID,
					UserId: commonUUID,
				},
				err: errors.New("error in create message"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatService{
				l:    tt.fields.l,
				repo: tt.fields.repo,
			}

			// we will use mock repo
			messageRepoMock := messageRepo.On("Create", mock.Anything).Return(tt.want.message, tt.want.err)

			defer messageRepoMock.Unset()

			got, err := c.Send(tt.args.m, tt.args.rId, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatService.Send() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				assert.Equal(t, tt.want.err, err)
				return
			}

			assert.Equal(t, tt.want.message.Msg, got.Msg)
			assert.Equal(t, tt.want.message.RoomId, got.RoomId)
			assert.Equal(t, tt.want.message.UserId, got.UserId)
		})
	}
}

func TestChatService_GetAll(t *testing.T) {
	repo, messageRepo := mocks.Init()

	l := utils.NewLogger()

	commonUUID := uuid.New()

	type GetAllReturnType struct {
		messages []models.Message
		err      error
	}
	type fields struct {
		l    models.Logger
		repo *repository.Repository
	}
	tests := []struct {
		name    string
		fields  fields
		want    GetAllReturnType
		wantErr bool
	}{
		{
			name: "Get all rooms",
			fields: fields{
				repo: repo,
				l:    l,
			},
			want: GetAllReturnType{
				messages: []models.Message{{
					ID:        commonUUID,
					Msg:       "Hello",
					RoomId:    commonUUID,
					UserId:    commonUUID,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
				}},
			},
			wantErr: false,
		},
		{
			name: "Error in Get all rooms",
			fields: fields{
				repo: repo,
				l:    l,
			},
			want: GetAllReturnType{
				messages: []models.Message{},
				err:      errors.New("error in fetching all rooms"),
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := ChatService{
				l:    tt.fields.l,
				repo: tt.fields.repo,
			}

			// we will use mock repo
			messageRepoMock := messageRepo.On("FindAll", mock.Anything).Return(tt.want.messages, tt.want.err)

			defer messageRepoMock.Unset()

			got, err := c.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("ChatService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if tt.wantErr {
				assert.Equal(t, tt.want.err, err)
				return
			}

			if !reflect.DeepEqual(got, tt.want.messages) {
				t.Errorf("ChatService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
