package models

import "time"

type Message struct {
	ID        CID       `json:"id"`
	Text      string    `json:"text"`
	RoomId    int64     `json:"roomId"`
	User      CID       `json:"user"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Room struct {
	ID        CID       `json:"id"`
	Name      string    `json:"name"`
	Users     []CID     `json:"users"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}
