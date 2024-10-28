package repository

import (
	"github.com/Ateto1204/swep-msg-serv/entity"
	"gorm.io/gorm"
)

type MsgRepository interface {
	Save(msg *entity.Message) error
	GetByID(id string) (entity.Message, error)
	UpdByID(id string) error
}

type msgRepository struct {
	db *gorm.DB
}

func NewMsgRepository(db *gorm.DB) MsgRepository {
	return &msgRepository{db}
}

func (r *msgRepository) Save(msg *entity.Message) error {
	return r.db.Create(msg).Error
}

func (r *msgRepository) GetByID(id string) (entity.Message, error) {
	var msg entity.Message
	err := r.db.Where("id = ?", id).Order("id").First(&msg).Error
	return msg, err
}

func (r *msgRepository) UpdByID(id string) error {
	return nil
}
