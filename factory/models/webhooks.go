package models

import (
	m "github.com/kubex-ecosystem/gdbase/internal/models/webhooks"
	"gorm.io/gorm"
)

type Webhook = m.IWebhook
type WebhookService = m.IWebhookService
type WebhookRepo = m.IWebhookRepo

// Define RegisterWebhookRequest here if it does not exist in the imported package
type RegisterWebhookRequest struct {
	// Add appropriate fields here, for example:
	FullUrl string `json:"fullUrl"`
	Event   string `json:"event"`
	Status  string `json:"status"`
}

// Define WebhookResponse here if it does not exist in the imported package
type WebhookResponse struct {
	ID      uint   `json:"id"`
	FullUrl string `json:"fullUrl"`
	Event   string `json:"event"`
	Status  string `json:"status"`
}

func NewWebhookService(webhookRepo WebhookRepo) WebhookService {
	return m.NewWebhookService(webhookRepo)
}

func NewWebhookRepo(db *gorm.DB) WebhookRepo {
	return m.NewWebhookRepo(db)
}

func NewWebhookModel(fullUrl, event, status string) Webhook {
	return m.NewWebhook(fullUrl, event, status)
}
