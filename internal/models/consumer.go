package models

type Payload struct {
	Event string
	Data  interface{}
	Write chan interface{}
}
