package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/Ateto1204/swep-msg-serv/internal/domain"
	"github.com/Ateto1204/swep-msg-serv/internal/repository"
)

type MsgUseCase interface {
	SaveMsg(userId, content string) (*domain.Message, error)
	GetMsg(id string) (*domain.Message, error)
	ReadMsg(msgID, userID string) error
	DeleteMsg(msgID string) error
}

type msgUseCase struct {
	repository repository.MsgRepository
}

func NewMsgUseCase(repo repository.MsgRepository) MsgUseCase {
	return &msgUseCase{repo}
}

func (uc *msgUseCase) SaveMsg(userID, content string) (*domain.Message, error) {
	t := time.Now()
	msgID := GenerateID()
	msg, err := uc.repository.Save(msgID, userID, content, t)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (uc *msgUseCase) GetMsg(msgID string) (*domain.Message, error) {
    var msg domain.Message
    err := uc.repository.db.Select("id", "content", "sender", "create_at").Where("id = ?", msgID).First(&msg).Error
    if err != nil {
        return nil, err
    }
    return &msg, nil
}


func (uc *msgUseCase) ReadMsg(msgID, userID string) error {
	msg, err := uc.repository.GetByID(msgID)
	if err != nil {
		return err
	}
	for _, member := range msg.Read {
		if member == userID {
			return errors.New("the member was already read this msg")
		}
	}
	msg.Read = append(msg.Read, userID)
	if err := uc.repository.UpdByID(msg); err != nil {
		return err
	}
	return nil
}

func (uc *msgUseCase) DeleteMsg(msgID string) error {
	return uc.repository.DeleteByID(msgID)
}

func GenerateID() string {
	timestamp := time.Now().UnixNano()

	input := fmt.Sprintf("%d", timestamp)

	hash := sha256.New()
	hash.Write([]byte(input))
	hashID := hex.EncodeToString(hash.Sum(nil))

	return hashID
}
