package repository

import (
	"fmt"
	"time"

	"github.com/Ateto1204/swep-msg-serv/entity"
	"github.com/Ateto1204/swep-msg-serv/internal/domain"
	"gorm.io/gorm"
)

type NotifRepository interface {
	Save(notifID, title, sender, content string, t time.Time) (*domain.Notification, error)
	GetByID(notifID string) (*domain.Notification, error)
	DeleteByID(notifID string) error
}

type notifRepository struct {
	db *gorm.DB
}

func NewNotifRepository(db *gorm.DB) NotifRepository {
	return &notifRepository{db}
}

func (r *notifRepository) Save(notifID, sender, title, content string, t time.Time) (*domain.Notification, error) {
	notifModel := domain.NewNotification(notifID, sender, title, content, t)
	notifEntity := parseToNotifEntity(notifModel)
	if err := r.db.Create(notifEntity).Error; err != nil {
		return nil, err
	}
	return notifModel, nil
}

func (r *notifRepository) GetByID(notifID string) (*domain.Notification, error) {
	var notifEntity *entity.Notification
	if err := r.db.Where("id = ?", notifID).Order("id").First(&notifEntity).Error; err != nil {
		return nil, err
	}
	notifModel := parseToNotifModel(notifEntity)
	return notifModel, nil
}

func (r *notifRepository) DeleteByID(notifID string) error {
	result := r.db.Where("id = ?", notifID).Delete(&entity.Notification{})
	if result.Error != nil {
		return fmt.Errorf("error occur when deleting the cart: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("cart %s was not found", notifID)
	}

	return nil
}

func parseToNotifEntity(notif *domain.Notification) *entity.Notification {
	notifEntity := &entity.Notification{
		ID:          notif.ID,
		Sender:      notif.Sender,
		Title:       notif.Title,
		Description: notif.Description,
		CreateAt:    notif.CreateAt,
	}
	return notifEntity
}

func parseToNotifModel(notif *entity.Notification) *domain.Notification {
	notifModel := &domain.Notification{
		ID:          notif.ID,
		Sender:      notif.Sender,
		Title:       notif.Title,
		Description: notif.Description,
		CreateAt:    notif.CreateAt,
	}
	return notifModel
}
