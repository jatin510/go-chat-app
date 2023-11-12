package services

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

func TestNewRoomService(t *testing.T) {
	type args struct {
		repo *repository.Repository
		l    models.Logger
	}
	tests := []struct {
		name string
		args args
		want RoomServiceInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRoomService(tt.args.repo, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRoomService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomService_Create(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		r       RoomService
		args    args
		want    models.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Create(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoomService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoomService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomService_GetAllRoomsByUserId(t *testing.T) {
	type args struct {
		userId uuid.UUID
	}
	tests := []struct {
		name    string
		r       RoomService
		args    args
		want    []models.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAllRoomsByUserId(tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("RoomService.GetAllRoomsByUserId() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoomService.GetAllRoomsByUserId() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRoomService_GetAll(t *testing.T) {
	tests := []struct {
		name    string
		r       RoomService
		want    []models.Room
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.GetAll()
			if (err != nil) != tt.wantErr {
				t.Errorf("RoomService.GetAll() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RoomService.GetAll() = %v, want %v", got, tt.want)
			}
		})
	}
}
