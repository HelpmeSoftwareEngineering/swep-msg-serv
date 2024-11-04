package domain

import "time"

type Message struct {
	ID       string
	Content  string
	Sender   string // User.ID
	CreateAt time.Time
	Read     bool
}
