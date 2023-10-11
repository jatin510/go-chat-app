package models

import "github.com/google/uuid"

type CID uuid.UUID

type Logger interface {
	Info(string, ...any)
	Debug(string, ...any)
	Error(string, ...any)
}
