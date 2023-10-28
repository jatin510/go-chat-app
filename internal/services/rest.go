package services

import (
	"io"
	"net/http"

	"github.com/jatin510/go-chat-app/internal/models"
)

type RestServiceInterface interface {
	Get(url string) error
}

type RestService struct {
	l models.Logger
}

func NewRestService(l models.Logger) RestServiceInterface {
	return &RestService{
		l: l,
	}
}

func (r RestService) Get(url string) error {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		r.l.Error("client: could not create request: ", err)
		return err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		r.l.Error("client: error making http request: ", err)
		return err
	}

	r.l.Info("client: got response!\n")
	r.l.Info("client: status code: %d\n", res.StatusCode)

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		r.l.Error("client: could not read response body: ", err)
		return err
	}
	r.l.Info("client: response body: ", resBody)

	return nil
}
