package models

type CID int64

type DBType int

type Logger interface {
	Info(...string)
	Debug(...string)
	Error(...string)
}
