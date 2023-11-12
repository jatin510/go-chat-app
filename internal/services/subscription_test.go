package services

import (
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

func TestNewSubscriptionService(t *testing.T) {
	type args struct {
		repo *repository.Repository
		l    models.Logger
	}
	tests := []struct {
		name string
		args args
		want SubscriptionInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewSubscriptionService(tt.args.repo, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewSubscriptionService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestSubscriptionService_Create(t *testing.T) {
	type args struct {
		roomId uuid.UUID
		userId uuid.UUID
	}
	tests := []struct {
		name    string
		r       SubscriptionService
		args    args
		want    models.Subscription
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.r.Create(tt.args.roomId, tt.args.userId)
			if (err != nil) != tt.wantErr {
				t.Errorf("SubscriptionService.Create() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("SubscriptionService.Create() = %v, want %v", got, tt.want)
			}
		})
	}
}
