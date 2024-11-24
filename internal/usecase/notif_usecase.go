package usecase

import (
	"time"

	"github.com/Ateto1204/swep-msg-serv/internal/domain"
	"github.com/Ateto1204/swep-msg-serv/internal/repository"
)

type NotifUseCase interface {
	SaveNotif(sender, title, content string) (*domain.Notification, error)
	GetNotif(notifID string) (*domain.Notification, error)
	DeleteNotif(notifID string) error
}

type notifUseCase struct {
	repo repository.NotifRepository
}

func NewNotifUseCase(repo repository.NotifRepository) NotifUseCase {
	return &notifUseCase{repo}
}

func (uc *notifUseCase) SaveNotif(sender, title, content string) (*domain.Notification, error) {
	t := time.Now()
	notifID := GenerateID()
	if sender == "" {
		sender = "Admin"
	}
	notif, err := uc.repo.Save(notifID, sender, title, content, t)
	if err != nil {
		return nil, err
	}
	return notif, err
}

func (uc *notifUseCase) GetNotif(notifID string) (*domain.Notification, error) {
	notif, err := uc.repo.GetByID(notifID)
	if err != nil {
		return nil, err
	}
	return notif, nil
}

func (uc *notifUseCase) DeleteNotif(notifID string) error {
	return uc.repo.DeleteByID(notifID)
}
