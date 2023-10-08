package models

type CID int64

type DBType int

type Logger interface {
	Info(string, ...any)
	Debug(string, ...any)
	Error(string, ...any)
}
