package usecase

import (
	"time"

	"github.com/Ateto1204/swep-msg-serv/entity"
	"github.com/Ateto1204/swep-msg-serv/internal/repository"
)

type MsgUseCase interface {
	SaveMsg(id, name string) error
	GetMsg(id string) (*entity.Message, error)
	GenerateID() string
}

type msgUseCase struct {
	repository repository.MsgRepository
}

func NewMsgUseCase(repo repository.MsgRepository) MsgUseCase {
	return &msgUseCase{
		repository: repo,
	}
}

func (uc *msgUseCase) SaveMsg(id, name string) error {
	t := time.Now()
	msg := &entity.Message{
		ID:       id,
		CreateAt: t,
	}
	err := uc.repository.Save(msg)
	if err != nil {
		return err
	}
	return nil
}

func (uc *msgUseCase) GetMsg(id string) (*entity.Message, error) {
	msg, err := uc.repository.GetByID(id)
	if err != nil {
		return nil, err
	}
	return &msg, nil
}

func (uc *msgUseCase) GenerateID() string {
	return ""
}
