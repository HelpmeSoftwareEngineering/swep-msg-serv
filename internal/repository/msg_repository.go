package repository

import (
	"time"

	"github.com/Ateto1204/swep-msg-serv/entity"
	"github.com/Ateto1204/swep-msg-serv/internal/domain"
	"gorm.io/gorm"
)

type MsgRepository interface {
	Save(msgID, sender, content string, t time.Time) (*domain.Message, error)
	GetByID(msgID string) (*domain.Message, error)
	UpdByID(msgID string) error
}

type msgRepository struct {
	db *gorm.DB
}

func NewMsgRepository(db *gorm.DB) MsgRepository {
	return &msgRepository{db}
}

func (r *msgRepository) Save(msgID, sender, content string, t time.Time) (*domain.Message, error) {
	msgModel := domain.NewMessage(msgID, sender, content, t)
	msgEntity, err := parseToMsgEntity(msgModel)
	if err != nil {
		return nil, err
	}
	if err := r.db.Create(msgEntity).Error; err != nil {
		return nil, err
	}
	return msgModel, nil
}

func (r *msgRepository) GetByID(msgID string) (*domain.Message, error) {
	var msgEntity *entity.Message
	if err := r.db.Where("id = ?", msgID).Order("id").First(&msgEntity).Error; err != nil {
		return nil, err
	}
	msgModel, err := parseToMsgModel(msgEntity)
	return msgModel, err
}

func (r *msgRepository) UpdByID(id string) error {
	err := r.db.Model(&entity.Message{}).Where("id = ?", id).Update("read", true).Error
	if err != nil {
		return err
	}
	return nil
}

func parseToMsgEntity(msg *domain.Message) (*entity.Message, error) {
	msgEntity := &entity.Message{
		ID:       msg.ID,
		Content:  msg.Content,
		Sender:   msg.Sender,
		CreateAt: msg.CreateAt,
		Read:     msg.Read,
	}
	return msgEntity, nil
}

func parseToMsgModel(msg *entity.Message) (*domain.Message, error) {
	msgModel := &domain.Message{
		ID:       msg.ID,
		Content:  msg.Content,
		Sender:   msg.Sender,
		CreateAt: msg.CreateAt,
		Read:     msg.Read,
	}
	return msgModel, nil
}
