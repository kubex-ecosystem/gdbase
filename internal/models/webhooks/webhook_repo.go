package webhooks

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"

	svc "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type IWebhookRepo interface {
	Create(webhook IWebhook) (IWebhook, error)
	List() ([]IWebhook, error)
	Update(webhook IWebhook) error
	Delete(id uuid.UUID) error
}

type WebhookRepo struct {
	*t.Mutexes
	webhooks  []IWebhook
	nextID    int
	dbService *svc.DBServiceImpl
}

func NewWebhookRepo(ctx context.Context, dbService svc.DBService) IWebhookRepo {
	var dbServiceImpl *svc.DBServiceImpl
	if dbService != nil {
		var ok bool
		dbServiceImpl, ok = dbService.(*svc.DBServiceImpl)
		if !ok {
			dbServiceImpl = nil
		}
	}

	return &WebhookRepo{
		webhooks:  make([]IWebhook, 0),
		nextID:    1,
		Mutexes:   &t.Mutexes{},
		dbService: dbServiceImpl,
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
