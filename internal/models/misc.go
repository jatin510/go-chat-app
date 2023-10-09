package models

type CID int64

type Logger interface {
	Info(string, ...any)
	Debug(string, ...any)
	Error(string, ...any)
}
