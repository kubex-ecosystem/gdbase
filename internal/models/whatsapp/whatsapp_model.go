// Package whatsapp provides the model for managing WhatsApp/Meta integrations in the MCP system.
package whatsapp

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

// WhatsAppUserType represents the type of WhatsApp user/entity
type WhatsAppUserType string

const (
	WhatsAppUserTypeBusiness WhatsAppUserType = "BUSINESS"
	WhatsAppUserTypeUser     WhatsAppUserType = "USER"
	WhatsAppUserTypeBot      WhatsAppUserType = "BOT"
	WhatsAppUserTypeSystem   WhatsAppUserType = "SYSTEM"
)

// WhatsAppStatus represents the status of WhatsApp integration
type WhatsAppStatus string

const (
	WhatsAppStatusActive       WhatsAppStatus = "ACTIVE"
	WhatsAppStatusInactive     WhatsAppStatus = "INACTIVE"
	WhatsAppStatusDisconnected WhatsAppStatus = "DISCONNECTED"
	WhatsAppStatusError        WhatsAppStatus = "ERROR"
	WhatsAppStatusBlocked      WhatsAppStatus = "BLOCKED"
	WhatsAppStatusPending      WhatsAppStatus = "PENDING"
)

// WhatsAppIntegrationType represents the type of WhatsApp integration
type WhatsAppIntegrationType string

const (
	WhatsAppIntegrationTypeBusinessAPI WhatsAppIntegrationType = "BUSINESS_API"
	WhatsAppIntegrationTypeCloudAPI    WhatsAppIntegrationType = "CLOUD_API"
	WhatsAppIntegrationTypeWebhook     WhatsAppIntegrationType = "WEBHOOK"
	WhatsAppIntegrationTypeGraph       WhatsAppIntegrationType = "GRAPH_API"
)

// WhatsAppMessageType represents the type of messages supported
type WhatsAppMessageType string

const (
	WhatsAppMessageTypeText     WhatsAppMessageType = "TEXT"
	WhatsAppMessageTypeImage    WhatsAppMessageType = "IMAGE"
	WhatsAppMessageTypeVideo    WhatsAppMessageType = "VIDEO"
	WhatsAppMessageTypeAudio    WhatsAppMessageType = "AUDIO"
	WhatsAppMessageTypeDocument WhatsAppMessageType = "DOCUMENT"
	WhatsAppMessageTypeLocation WhatsAppMessageType = "LOCATION"
	WhatsAppMessageTypeTemplate WhatsAppMessageType = "TEMPLATE"
	WhatsAppMessageTypeButton   WhatsAppMessageType = "BUTTON"
	WhatsAppMessageTypeList     WhatsAppMessageType = "LIST"
)

// IWhatsAppModel defines the interface for WhatsApp integrations
type IWhatsAppModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetWhatsAppBusinessID() string
	SetWhatsAppBusinessID(businessID string)
	GetWhatsAppUserID() string
	SetWhatsAppUserID(userID string)
	GetPhoneNumber() string
	SetPhoneNumber(phoneNumber string)
	GetPhoneNumberID() string
	SetPhoneNumberID(phoneNumberID string)
	GetDisplayName() string
	SetDisplayName(displayName string)
	GetBusinessName() string
	SetBusinessName(businessName string)
	GetUserType() WhatsAppUserType
	SetUserType(userType WhatsAppUserType)
	GetStatus() WhatsAppStatus
	SetStatus(status WhatsAppStatus)
	GetIntegrationType() WhatsAppIntegrationType
	SetIntegrationType(integrationType WhatsAppIntegrationType)
	GetAccessToken() string
	SetAccessToken(accessToken string)
	GetRefreshToken() string
	SetRefreshToken(refreshToken string)
	GetTokenExpiresAt() time.Time
	SetTokenExpiresAt(expiresAt time.Time)
	GetWebhookURL() string
	SetWebhookURL(webhookURL string)
	GetWebhookVerifyToken() string
	SetWebhookVerifyToken(verifyToken string)
	GetAppID() string
	SetAppID(appID string)
	GetAppSecret() string
	SetAppSecret(appSecret string)
	GetBusinessConfig() t.JSONB
	SetBusinessConfig(config t.JSONB)
	GetSupportedMessageTypes() []string
	SetSupportedMessageTypes(types []string)
	GetConfig() t.JSONB
	SetConfig(config t.JSONB)
	GetLastActivity() time.Time
	SetLastActivity(lastActivity time.Time)
	GetUserID() string
	SetUserID(userID string)
	GetTargetTaskID() string
	SetTargetTaskID(targetTaskID string)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetCreatedBy() string
	SetCreatedBy(createdBy string)
	GetUpdatedBy() string
	SetUpdatedBy(updatedBy string)
	Validate() error
	Sanitize()
}

// WhatsAppModel represents the WhatsApp integration entity in the database
type WhatsAppModel struct {
	ID                    string                  `gorm:"type:uuid;primaryKey" json:"id"`
	WhatsAppBusinessID    string                  `gorm:"type:text;index" json:"whatsapp_business_id,omitempty" example:"102290129340398"`
	WhatsAppUserID        string                  `gorm:"type:text;index" json:"whatsapp_user_id,omitempty" example:"102290129340398"`
	PhoneNumber           string                  `gorm:"type:text;not null;uniqueIndex" json:"phone_number" example:"+5511987654321"`
	PhoneNumberID         string                  `gorm:"type:text;index" json:"phone_number_id,omitempty" example:"109529185312345"`
	DisplayName           string                  `gorm:"type:text" json:"display_name,omitempty" example:"My Business"`
	BusinessName          string                  `gorm:"type:text" json:"business_name,omitempty" example:"My Business Corp"`
	UserType              WhatsAppUserType        `gorm:"type:text;not null;default:'USER'" json:"user_type" example:"BUSINESS"`
	Status                WhatsAppStatus          `gorm:"type:text;not null;default:'ACTIVE'" json:"status" example:"ACTIVE"`
	IntegrationType       WhatsAppIntegrationType `gorm:"type:text;not null;default:'BUSINESS_API'" json:"integration_type" example:"BUSINESS_API"`
	AccessToken           string                  `gorm:"type:text" json:"access_token,omitempty"`
	RefreshToken          string                  `gorm:"type:text" json:"refresh_token,omitempty"`
	TokenExpiresAt        string                  `gorm:"type:timestamp" json:"token_expires_at,omitempty" example:"2024-01-01T00:00:00Z"`
	WebhookURL            string                  `gorm:"type:text" json:"webhook_url,omitempty"`
	WebhookVerifyToken    string                  `gorm:"type:text" json:"webhook_verify_token,omitempty"`
	AppID                 string                  `gorm:"type:text" json:"app_id,omitempty" example:"1234567890123456"`
	AppSecret             string                  `gorm:"type:text" json:"app_secret,omitempty"`
	BusinessConfig        t.JSONB                 `json:"business_config,omitempty"`
	SupportedMessageTypes []string                `gorm:"type:text[]" json:"supported_message_types,omitempty" example:"[\"text\", \"image\", \"document\"]"`
	Config                t.JSONB                 `json:"config" binding:"omitempty"`
	LastActivity          string                  `gorm:"type:timestamp;default:now()" json:"last_activity,omitempty" example:"2024-01-01T00:00:00Z"`
	UserID                string                  `gorm:"type:uuid;references:users(id)" json:"user_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	TargetTaskID          string                  `gorm:"type:uuid;references:mcp_sync_tasks(id)" json:"target_task_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
	CreatedAt             string                  `gorm:"type:timestamp;default:now()" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt             string                  `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy             string                  `gorm:"type:uuid;references:users(id)" json:"created_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	UpdatedBy             string                  `gorm:"type:uuid;references:users(id)" json:"updated_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
}

func NewWhatsAppModel() *WhatsAppModel {
	return &WhatsAppModel{
		ID:                    uuid.New().String(),
		WhatsAppBusinessID:    "",
		WhatsAppUserID:        "",
		PhoneNumber:           "",
		PhoneNumberID:         "",
		DisplayName:           "",
		BusinessName:          "",
		UserType:              WhatsAppUserTypeUser,
		Status:                WhatsAppStatusActive,
		IntegrationType:       WhatsAppIntegrationTypeBusinessAPI,
		AccessToken:           "",
		RefreshToken:          "",
		TokenExpiresAt:        time.Now().Add(24 * time.Hour).Format(time.RFC3339),
		WebhookURL:            "",
		WebhookVerifyToken:    "",
		AppID:                 "",
		AppSecret:             "",
		BusinessConfig:        t.JSONB{},
		SupportedMessageTypes: []string{"text", "image", "document"},
		Config:                t.JSONB{},
		LastActivity:          time.Now().Format(time.RFC3339),
		UserID:                "",
		TargetTaskID:          "",
		CreatedAt:             time.Now().Format(time.RFC3339),
		UpdatedAt:             time.Now().Format(time.RFC3339),
		CreatedBy:             "",
		UpdatedBy:             "",
	}
}

func (w *WhatsAppModel) TableName() string { return "mcp_whatsapp_integrations" }

// Basic getters and setters

func (w *WhatsAppModel) GetID() string                           { return w.ID }
func (w *WhatsAppModel) SetID(id string)                         { w.ID = id }
func (w *WhatsAppModel) GetWhatsAppBusinessID() string           { return w.WhatsAppBusinessID }
func (w *WhatsAppModel) SetWhatsAppBusinessID(businessID string) { w.WhatsAppBusinessID = businessID }
func (w *WhatsAppModel) GetWhatsAppUserID() string               { return w.WhatsAppUserID }
func (w *WhatsAppModel) SetWhatsAppUserID(userID string)         { w.WhatsAppUserID = userID }
func (w *WhatsAppModel) GetPhoneNumber() string                  { return w.PhoneNumber }
func (w *WhatsAppModel) SetPhoneNumber(phoneNumber string)       { w.PhoneNumber = phoneNumber }
func (w *WhatsAppModel) GetPhoneNumberID() string                { return w.PhoneNumberID }
func (w *WhatsAppModel) SetPhoneNumberID(phoneNumberID string)   { w.PhoneNumberID = phoneNumberID }
func (w *WhatsAppModel) GetDisplayName() string                  { return w.DisplayName }
func (w *WhatsAppModel) SetDisplayName(displayName string)       { w.DisplayName = displayName }
func (w *WhatsAppModel) GetBusinessName() string                 { return w.BusinessName }
func (w *WhatsAppModel) SetBusinessName(businessName string)     { w.BusinessName = businessName }

// Type-specific getters and setters

func (w *WhatsAppModel) GetUserType() WhatsAppUserType               { return w.UserType }
func (w *WhatsAppModel) SetUserType(userType WhatsAppUserType)       { w.UserType = userType }
func (w *WhatsAppModel) GetStatus() WhatsAppStatus                   { return w.Status }
func (w *WhatsAppModel) SetStatus(status WhatsAppStatus)             { w.Status = status }
func (w *WhatsAppModel) GetIntegrationType() WhatsAppIntegrationType { return w.IntegrationType }
func (w *WhatsAppModel) SetIntegrationType(integrationType WhatsAppIntegrationType) {
	w.IntegrationType = integrationType
}

// Token and authentication getters and setters

func (w *WhatsAppModel) GetAccessToken() string              { return w.AccessToken }
func (w *WhatsAppModel) SetAccessToken(accessToken string)   { w.AccessToken = accessToken }
func (w *WhatsAppModel) GetRefreshToken() string             { return w.RefreshToken }
func (w *WhatsAppModel) SetRefreshToken(refreshToken string) { w.RefreshToken = refreshToken }

func (w *WhatsAppModel) GetTokenExpiresAt() time.Time {
	expiresAt, _ := time.Parse(time.RFC3339, w.TokenExpiresAt)
	return expiresAt
}
func (w *WhatsAppModel) SetTokenExpiresAt(expiresAt time.Time) {
	w.TokenExpiresAt = expiresAt.Format(time.RFC3339)
}

func (w *WhatsAppModel) GetWebhookURL() string                    { return w.WebhookURL }
func (w *WhatsAppModel) SetWebhookURL(webhookURL string)          { w.WebhookURL = webhookURL }
func (w *WhatsAppModel) GetWebhookVerifyToken() string            { return w.WebhookVerifyToken }
func (w *WhatsAppModel) SetWebhookVerifyToken(verifyToken string) { w.WebhookVerifyToken = verifyToken }
func (w *WhatsAppModel) GetAppID() string                         { return w.AppID }
func (w *WhatsAppModel) SetAppID(appID string)                    { w.AppID = appID }
func (w *WhatsAppModel) GetAppSecret() string                     { return w.AppSecret }
func (w *WhatsAppModel) SetAppSecret(appSecret string)            { w.AppSecret = appSecret }
func (w *WhatsAppModel) GetBusinessConfig() t.JSONB               { return w.BusinessConfig }
func (w *WhatsAppModel) SetBusinessConfig(config t.JSONB)         { w.BusinessConfig = config }
func (w *WhatsAppModel) GetSupportedMessageTypes() []string       { return w.SupportedMessageTypes }
func (w *WhatsAppModel) SetSupportedMessageTypes(types []string)  { w.SupportedMessageTypes = types }
func (w *WhatsAppModel) GetConfig() t.JSONB                       { return w.Config }
func (w *WhatsAppModel) SetConfig(config t.JSONB)                 { w.Config = config }

func (w *WhatsAppModel) GetLastActivity() time.Time {
	lastActivity, _ := time.Parse(time.RFC3339, w.LastActivity)
	return lastActivity
}
func (w *WhatsAppModel) SetLastActivity(lastActivity time.Time) {
	w.LastActivity = lastActivity.Format(time.RFC3339)
}

// MCP integration getters and setters

func (w *WhatsAppModel) GetUserID() string                   { return w.UserID }
func (w *WhatsAppModel) SetUserID(userID string)             { w.UserID = userID }
func (w *WhatsAppModel) GetTargetTaskID() string             { return w.TargetTaskID }
func (w *WhatsAppModel) SetTargetTaskID(targetTaskID string) { w.TargetTaskID = targetTaskID }

// Timestamp getters and setters

func (w *WhatsAppModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, w.CreatedAt)
	return createdAt
}
func (w *WhatsAppModel) SetCreatedAt(createdAt time.Time) {
	w.CreatedAt = createdAt.Format(time.RFC3339)
}
func (w *WhatsAppModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, w.UpdatedAt)
	return updatedAt
}
func (w *WhatsAppModel) SetUpdatedAt(updatedAt time.Time) {
	w.UpdatedAt = updatedAt.Format(time.RFC3339)
}
func (w *WhatsAppModel) GetCreatedBy() string          { return w.CreatedBy }
func (w *WhatsAppModel) SetCreatedBy(createdBy string) { w.CreatedBy = createdBy }
func (w *WhatsAppModel) GetUpdatedBy() string          { return w.UpdatedBy }
func (w *WhatsAppModel) SetUpdatedBy(updatedBy string) { w.UpdatedBy = updatedBy }

// Validation method

func (w *WhatsAppModel) Validate() error {
	if w.PhoneNumber == "" {
		return fmt.Errorf("phone_number is required")
	}
	if w.UserType == "" {
		w.UserType = WhatsAppUserTypeUser
	}
	if w.Status == "" {
		w.Status = WhatsAppStatusActive
	}
	if w.IntegrationType == "" {
		w.IntegrationType = WhatsAppIntegrationTypeBusinessAPI
	}

	// Validate integration type requirements
	switch w.IntegrationType {
	case WhatsAppIntegrationTypeBusinessAPI, WhatsAppIntegrationTypeCloudAPI:
		if w.AccessToken == "" {
			return fmt.Errorf("access_token is required for %s integration", w.IntegrationType)
		}
		if w.PhoneNumberID == "" {
			return fmt.Errorf("phone_number_id is required for %s integration", w.IntegrationType)
		}
	case WhatsAppIntegrationTypeWebhook:
		if w.WebhookURL == "" {
			return fmt.Errorf("webhook_url is required for webhook integration")
		}
		if w.WebhookVerifyToken == "" {
			return fmt.Errorf("webhook_verify_token is required for webhook integration")
		}
	case WhatsAppIntegrationTypeGraph:
		if w.AppID == "" {
			return fmt.Errorf("app_id is required for Graph API integration")
		}
		if w.AppSecret == "" {
			return fmt.Errorf("app_secret is required for Graph API integration")
		}
	}

	return nil
}

// Sanitize method
func (w *WhatsAppModel) Sanitize() {
	w.UpdatedAt = time.Now().Format(time.RFC3339)
	if w.LastActivity == "" {
		w.LastActivity = time.Now().Format(time.RFC3339)
	}

	// Build display name if not provided
	if w.DisplayName == "" {
		if w.BusinessName != "" {
			w.DisplayName = w.BusinessName
		} else {
			w.DisplayName = w.PhoneNumber
		}
	}

	// Set default supported message types if empty
	if len(w.SupportedMessageTypes) == 0 {
		w.SupportedMessageTypes = []string{"text", "image", "document"}
	}
}
