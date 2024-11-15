package repository

import (
	"time"

	"github.com/Ateto1204/swep-msg-serv/internal/domain"
	"gorm.io/gorm"
)

type NotifRepository interface {
	Save(notifID, sender, content string, t time.Time) (*domain.Notification, error)
	GetByID(notifID string) (*domain.Notification, error)
	UpdByID(notifID string) error
}

type notifRepository struct {
	db *gorm.DB
}

func NewNotifRepository(db *gorm.DB) NotifRepository {
	return &notifRepository{db}
}

func (r *notifRepository) Save(notifID, sender, content string, t time.Time) (*domain.Notification, error) {
	return nil, nil
}

func (r *notifRepository) GetByID(notifID string) (*domain.Notification, error) {
	return nil, nil
}

func (r *notifRepository) UpdByID(notifID string) error {
	return nil
}
