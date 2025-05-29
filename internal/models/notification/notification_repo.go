package notification

import (
	"errors"
	"sync"
)

// MemoryNotificationRepo é uma implementação em memória do repositório.
type MemoryNotificationRepo struct {
	mu            sync.Mutex
	notifications []Notification
	nextID        int
}

// Repository: Interface de persistência de notificações.
type NotificationRepo interface {
	Create(notification Notification) (Notification, error)
	GetAll() ([]Notification, error)
	Delete(id int) error
}

func NewMemoryNotificationRepo() *MemoryNotificationRepo {
	return &MemoryNotificationRepo{
		notifications: make([]Notification, 0),
		nextID:        1,
	}
}

func (repo *MemoryNotificationRepo) Create(notification Notification) (Notification, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	notification.ID = repo.nextID
	repo.nextID++
	repo.notifications = append(repo.notifications, notification)
	return notification, nil
}

func (repo *MemoryNotificationRepo) GetAll() ([]Notification, error) {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	return repo.notifications, nil
}

func (repo *MemoryNotificationRepo) Delete(id int) error {
	repo.mu.Lock()
	defer repo.mu.Unlock()
	for i, notif := range repo.notifications {
		if notif.ID == id {
			repo.notifications = append(repo.notifications[:i], repo.notifications[i+1:]...)
			return nil
		}
	}
	return errors.New("notificação não encontrada")
}
