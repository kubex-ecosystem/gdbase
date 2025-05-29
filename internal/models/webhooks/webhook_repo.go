package webhooks

import (
	"errors"
	"time"

	"github.com/google/uuid"
	t "github.com/rafa-mori/gdbase/internal/types"
	"gorm.io/gorm"
)

type IWebhookRepo interface {
	Create(webhook IWebhook) (IWebhook, error)
	List() ([]IWebhook, error)
	Update(webhook IWebhook) error
	Delete(id uuid.UUID) error
}

type WebhookRepo struct {
	*t.Mutexes
	webhooks []IWebhook
	nextID   int
}

func NewWebhookRepo(db *gorm.DB) IWebhookRepo {
	return &WebhookRepo{
		webhooks: make([]IWebhook, 0),
		nextID:   1,
	}
}

func (repo *WebhookRepo) Create(webhook IWebhook) (IWebhook, error) {
	repo.Mutexes.MuLock()
	defer repo.Mutexes.MuUnlock()
	webhook.SetCreatedAt(time.Now())
	webhook.SetUpdatedAt(time.Now())
	repo.webhooks = append(repo.webhooks, webhook)
	return webhook, nil
}

func (repo *WebhookRepo) List() ([]IWebhook, error) {
	repo.Mutexes.MuLock()
	defer repo.Mutexes.MuUnlock()
	return repo.webhooks, nil
}

func (repo *WebhookRepo) Update(webhook IWebhook) error {
	repo.Mutexes.MuLock()
	defer repo.Mutexes.MuUnlock()
	for i, w := range repo.webhooks {
		if w.GetID() == webhook.GetID() {
			webhook.SetUpdatedAt(time.Now())
			repo.webhooks[i] = webhook
			return nil
		}
	}
	return errors.New("webhook não encontrado")
}

func (repo *WebhookRepo) Delete(id uuid.UUID) error {
	repo.Mutexes.MuLock()
	defer repo.Mutexes.MuUnlock()
	for i, w := range repo.webhooks {
		if w.GetID() == id {
			repo.webhooks = append(repo.webhooks[:i], repo.webhooks[i+1:]...)
			return nil
		}
	}
	return errors.New("webhook não encontrado")
}
