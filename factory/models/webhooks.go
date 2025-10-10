package models

import (
	"context"

	m "github.com/kubex-ecosystem/gdbase/internal/models/webhooks"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

type Webhook = m.IWebhook
type WebhookService = m.IWebhookService
type WebhookRepo = m.IWebhookRepo

// Define RegisterWebhookRequest here if it does not exist in the imported package

type RegisterWebhookRequest struct {
	// Add appropriate fields here, for example:
	FullURL string `json:"fullUrl"`
	Event   string `json:"event"`
	Status  string `json:"status"`
}

// Define WebhookResponse here if it does not exist in the imported package

type WebhookResponse struct {
	ID      uint   `json:"id"`
	FullURL string `json:"fullUrl"`
	Event   string `json:"event"`
	Status  string `json:"status"`
}

func NewWebhookService(webhookRepo WebhookRepo) WebhookService {
	return m.NewWebhookService(webhookRepo)
}

func NewWebhookRepo(ctx context.Context, dbService svc.DBService) WebhookRepo {
	return m.NewWebhookRepo(ctx, dbService)
}

func NewWebhookModel(fullURL, event, status string) Webhook {
	return m.NewWebhook(fullURL, event, status)
}
