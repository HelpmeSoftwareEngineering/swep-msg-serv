package entity

import "time"

type Notification struct {
	ID          string    `gorm:"primaryKey" json:"id"`
	Sender      string    `json:"sender"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}
