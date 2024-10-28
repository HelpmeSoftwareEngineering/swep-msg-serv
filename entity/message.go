package entity

import "time"

type Message struct {
	ID       string    `gorm:"primaryKey" json:"id"`
	Content  string    `json:"content"`
	Sender   string    `json:"sender"` // User.ID
	CreateAt time.Time `json:"time"`
	Read     bool      `json:"read"`
}
