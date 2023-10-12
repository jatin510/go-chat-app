package models

type Logger interface {
	Info(string, ...any)
	Debug(string, ...any)
	Error(string, ...any)
}
