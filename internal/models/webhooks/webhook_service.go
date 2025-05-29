package webhooks

import (
	"fmt"

	"github.com/google/uuid"
	t "github.com/rafa-mori/gdbase/internal/types"
)

type IWebhookService interface {
	RegisterWebhook(webhook IWebhook) (IWebhook, error)
	ListWebhooks() ([]IWebhook, error)
	RemoveWebhook(id uuid.UUID) error
}

type WebhookService struct {
	Mutexes   *t.Mutexes
	Reference *t.Reference
	repo      IWebhookRepo
}

func NewWebhookService(repo IWebhookRepo) IWebhookService {
	return &WebhookService{
		Mutexes:   t.NewMutexesType(),
		Reference: t.NewReference("webhook_service").GetReference(),
		repo:      repo,
	}
}

func (s *WebhookService) RegisterWebhook(webhook IWebhook) (IWebhook, error) {
	// Aqui você pode incluir validações
	webhook.SetStatus("ativo")
	created, err := s.repo.Create(webhook.(*Webhook))
	if err != nil {
		return nil, err
	}
	fmt.Printf("Webhook registrado: %+v\n", created)
	return created, nil
}

func (s *WebhookService) ListWebhooks() ([]IWebhook, error) {
	return s.repo.List()
}

func (s *WebhookService) RemoveWebhook(id uuid.UUID) error {
	// Aqui você pode incluir lógica para encerrar child servers associados
	return s.repo.Delete(id)
}

// Poderia incluir métodos adicionais, como atualizar ou monitorar tasks associadas.
