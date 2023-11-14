package services

import (
	"reflect"
	"testing"

	"github.com/jatin510/go-chat-app/internal/models"
)

func TestNewRestService(t *testing.T) {
	type args struct {
		l models.Logger
	}
	tests := []struct {
		name string
		args args
		want RestServiceInterface
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewRestService(tt.args.l); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewRestService() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRestService_Get(t *testing.T) {
	type args struct {
		url string
	}
	tests := []struct {
		name    string
		r       RestService
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.r.Get(tt.args.url); (err != nil) != tt.wantErr {
				t.Errorf("RestService.Get() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
