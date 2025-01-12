package repository

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/Ateto1204/swep-msg-serv/entity"
	"github.com/Ateto1204/swep-msg-serv/internal/domain"
	"gorm.io/gorm"
)

type MsgRepository interface {
	Save(msgID, sender, content string, t time.Time) (*domain.Message, error)
	GetByID(msgID string) (*domain.Message, error)
	UpdByID(msg *domain.Message) error
	DeleteByID(msgID string) error
	GetMsgByID(msgID string) (*domain.Message, error)
}

type msgRepository struct {
	db *gorm.DB
}

func NewMsgRepository(db *gorm.DB) MsgRepository {
	return &msgRepository{db}
}

func (r *msgRepository) Save(msgID, sender, content string, t time.Time) (*domain.Message, error) {
	msgEntity := &entity.Message{
		ID:       msgID,
		Content:  content,
		Sender:   sender,
		CreateAt: t,
		Read:     "false", // 預設為 false
	}

	if err := r.db.Create(msgEntity).Error; err != nil {
		log.Println("DB Create message fail: msgID %s", msgID)
		return nil, err
	}

	msgModel := &domain.Message{
		ID:       msgEntity.ID,
		Content:  msgEntity.Content,
		Sender:   msgEntity.Sender,
		CreateAt: msgEntity.CreateAt,
		Read:     false,
	}

	log.Println("Save msg model success")
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

func (r *msgRepository) GetMsgByID(msgID string) (*domain.Message, error) {
	var msgEntity *entity.Message
	if err := r.db.Select("id", "content", "sender", "create_at", "read").Where("id = ?", msgID).First(&msgEntity).Error; err != nil {
		log.Println("DB Select message fail: msgID %s", msgID)
		return nil, err
	}

	msgModel := &domain.Message{
		ID:       msgEntity.ID,
		Content:  msgEntity.Content,
		Sender:   msgEntity.Sender,
		CreateAt: msgEntity.CreateAt,
		Read:     false, // 布林值直接映射
	}
	log.Println("Get msg model success %t", msgEntity.Read)
	return msgModel, nil
}

func (r *msgRepository) UpdByID(msg *domain.Message) error {
	msgEntity, err := parseToMsgEntity(msg)
	if err != nil {
		return err
	}
	if err = r.db.Model(&entity.Message{}).Where("id = ?", msgEntity.ID).Update("read", msgEntity.Read).Error; err != nil {
		return err
	}
	return nil
}

func (r *msgRepository) DeleteByID(msgID string) error {
	result := r.db.Where("id = ?", msgID).Delete(&entity.Message{})
	if result.Error != nil {
		return fmt.Errorf("error occur when deleting the msg: %w", result.Error)
	}

	if result.RowsAffected == 0 {
		return fmt.Errorf("msg %s was not found", msgID)
	}
	return nil
}

func parseToMsgEntity(msg *domain.Message) (*entity.Message, error) {
	readStr := "false"
	msgEntity := &entity.Message{
		ID:       msg.ID,
		Content:  msg.Content,
		Sender:   msg.Sender,
		CreateAt: msg.CreateAt,
		Read:     readStr, // 直接將 bool 賦值給 Read
	}
	return msgEntity, nil
}

func parseToMsgModel(msg *entity.Message) (*domain.Message, error) {
	msgModel := &domain.Message{
		ID:       msg.ID,
		Content:  msg.Content,
		Sender:   msg.Sender,
		CreateAt: msg.CreateAt,
		Read:     false,
	}
	return msgModel, nil
}

func strSerialize(sa []string) (string, error) {
	s, err := json.Marshal(sa)
	if err != nil {
		return "", err
	}
	return string(s), nil
}

func strUnserialize(s string) ([]string, error) {
	var ca []string
	err := json.Unmarshal([]byte(s), &ca)
	return ca, err
}
