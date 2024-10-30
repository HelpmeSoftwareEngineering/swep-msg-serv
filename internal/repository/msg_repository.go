package repository

import (
	"time"

	"github.com/Ateto1204/swep-msg-serv/entity"
	"gorm.io/gorm"
)

type MsgRepository interface {
	Save(id, sender, content string, t time.Time) (*entity.Message, error)
	GetByID(id string) (entity.Message, error)
	UpdByID(id string) error
}

type msgRepository struct {
	db *gorm.DB
}

func NewMsgRepository(db *gorm.DB) MsgRepository {
	return &msgRepository{db}
}

func (r *msgRepository) Save(id, sender, content string, t time.Time) (*entity.Message, error) {
	msg := &entity.Message{
		ID:       id,
		Content:  content,
		Sender:   sender,
		CreateAt: t,
		Read:     false,
	}
	err := r.db.Create(msg).Error
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (r *msgRepository) GetByID(id string) (entity.Message, error) {
	var msg entity.Message
	err := r.db.Where("id = ?", id).Order("id").First(&msg).Error
	return msg, err
}

func (r *msgRepository) UpdByID(id string) error {
	return nil
}
