package domain

import "time"

type Message struct {
	ID       string    `json:"id"`
	Content  string    `json:"content"`
	Sender   string    `json:"sender"` // User.ID
	CreateAt time.Time `json:"time"`
	Read     bool      `json:"read"`
}

type Notification struct {
	ID          string    `json:"id"`
	Sender      string    `json:"sender"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	CreateAt    time.Time `json:"create_at"`
}

func NewMessage(ID, sender, content string, t time.Time) *Message {
	return &Message{
		ID:       ID,
		Content:  content,
		Sender:   sender,
		CreateAt: t,
		Read:     false,
	}
}

func NewNotification(ID, sender, title, descript string, t time.Time) *Notification {
	return &Notification{
		ID:          ID,
		Sender:      sender,
		Title:       title,
		Description: descript,
		CreateAt:    t,
	}
}
