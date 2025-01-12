package entity

import "time"

type Message struct {
	ID       string    `json:"id"`
	Content  string    `json:"content"`
	Sender   string    `json:"sender"`
	CreateAt time.Time `json:"time"`
	Read     string    `json:"read"`
}
