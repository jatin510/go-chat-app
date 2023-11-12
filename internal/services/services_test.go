package services

import (
	"reflect"
	"testing"

	"github.com/jatin510/go-chat-app/internal/models"
	"github.com/jatin510/go-chat-app/internal/repository"
)

func TestInit(t *testing.T) {
	type args struct {
		repo *repository.Repository
		l    models.Logger
	}
	tests := []struct {
		name string
		args args
		want *Services
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Init(tt.args.repo, tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Init() = %v, want %v", got, tt.want)
			}
		})
	}
}
