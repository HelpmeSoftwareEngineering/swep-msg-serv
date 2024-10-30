package usecase

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"

	"github.com/Ateto1204/swep-msg-serv/entity"
	"github.com/Ateto1204/swep-msg-serv/internal/repository"
)

type MsgUseCase interface {
	SaveMsg(userId, content string) (*entity.Message, error)
	GetMsg(id string) (*entity.Message, error)
}

type msgUseCase struct {
	repository repository.MsgRepository
}

func NewMsgUseCase(repo repository.MsgRepository) MsgUseCase {
	return &msgUseCase{
		repository: repo,
	}
}

func (uc *msgUseCase) SaveMsg(userID, content string) (*entity.Message, error) {
	t := time.Now()
	id := GenerateID()
	msg, err := uc.repository.Save(id, userID, content, t)
	if err != nil {
		return nil, err
	}
	return msg, nil
}

func (uc *msgUseCase) GetMsg(id string) (*entity.Message, error) {
	msg, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func GenerateID() string {
	timestamp := time.Now().UnixNano()

	input := fmt.Sprintf("%d", timestamp)

	hash := sha256.New()
	hash.Write([]byte(input))
	hashID := hex.EncodeToString(hash.Sum(nil))

	return hashID
}
