// Package telegram provides the model for managing Telegram integrations in the MCP system.
package telegram

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/types"
)

// TelegramUserType represents the type of Telegram user/entity
type TelegramUserType string

const (
	TelegramUserTypeBot     TelegramUserType = "BOT"
	TelegramUserTypeUser    TelegramUserType = "USER"
	TelegramUserTypeChannel TelegramUserType = "CHANNEL"
	TelegramUserTypeGroup   TelegramUserType = "GROUP"
	TelegramUserTypeSystem  TelegramUserType = "SYSTEM"
)

// TelegramStatus represents the status of Telegram integration
type TelegramStatus string

const (
	TelegramStatusActive       TelegramStatus = "ACTIVE"
	TelegramStatusInactive     TelegramStatus = "INACTIVE"
	TelegramStatusDisconnected TelegramStatus = "DISCONNECTED"
	TelegramStatusError        TelegramStatus = "ERROR"
	TelegramStatusBlocked      TelegramStatus = "BLOCKED"
)

// TelegramIntegrationType represents the type of Telegram integration
type TelegramIntegrationType string

const (
	TelegramIntegrationTypeBot     TelegramIntegrationType = "BOT"
	TelegramIntegrationTypeWebhook TelegramIntegrationType = "WEBHOOK"
	TelegramIntegrationTypeAPI     TelegramIntegrationType = "API"
)

// ITelegramModel defines the interface for Telegram integrations
type ITelegramModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetTelegramUserID() string
	SetTelegramUserID(telegramUserID string)
	GetUsername() string
	SetUsername(username string)
	GetFirstName() string
	SetFirstName(firstName string)
	GetLastName() string
	SetLastName(lastName string)
	GetDisplayName() string
	SetDisplayName(displayName string)
	GetPhoneNumber() string
	SetPhoneNumber(phoneNumber string)
	GetLanguageCode() string
	SetLanguageCode(languageCode string)
	GetUserType() TelegramUserType
	SetUserType(userType TelegramUserType)
	GetStatus() TelegramStatus
	SetStatus(status TelegramStatus)
	GetIntegrationType() TelegramIntegrationType
	SetIntegrationType(integrationType TelegramIntegrationType)
	GetChatID() string
	SetChatID(chatID string)
	GetChannelID() string
	SetChannelID(channelID string)
	GetGroupID() string
	SetGroupID(groupID string)
	GetBotToken() string
	SetBotToken(botToken string)
	GetWebhookURL() string
	SetWebhookURL(webhookURL string)
	GetApiKey() string
	SetApiKey(apiKey string)
	GetBotPermissions() t.JSONB
	SetBotPermissions(permissions t.JSONB)
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

// TelegramModel represents the Telegram integration entity in the database
type TelegramModel struct {
	ID              string                      `gorm:"type:uuid;primaryKey" json:"id"`
	TelegramUserID  string                      `gorm:"type:text;not null;uniqueIndex" json:"telegram_user_id" example:"123456789"`
	Username        string                      `gorm:"type:text" json:"username,omitempty" example:"telegram_user"`
	FirstName       string                      `gorm:"type:text;not null" json:"first_name" example:"John"`
	LastName        string                      `gorm:"type:text" json:"last_name,omitempty" example:"Doe"`
	DisplayName     string                      `gorm:"type:text" json:"display_name,omitempty" example:"John Doe"`
	PhoneNumber     string                      `gorm:"type:text" json:"phone_number,omitempty" example:"+5511987654321"`
	LanguageCode    string                      `gorm:"type:text" json:"language_code,omitempty" example:"pt-BR"`
	UserType        TelegramUserType            `gorm:"type:text;not null;default:'USER'" json:"user_type" example:"USER"`
	Status          TelegramStatus              `gorm:"type:text;not null;default:'ACTIVE'" json:"status" example:"ACTIVE"`
	IntegrationType TelegramIntegrationType     `gorm:"type:text;not null;default:'BOT'" json:"integration_type" example:"BOT"`
	ChatID          string                      `gorm:"type:text" json:"chat_id,omitempty" example:"123456789"`
	ChannelID       string                      `gorm:"type:text" json:"channel_id,omitempty" example:"@mychannel"`
	GroupID         string                      `gorm:"type:text" json:"group_id,omitempty" example:"-123456789"`
	BotToken        string                      `gorm:"type:text" json:"bot_token,omitempty"`
	WebhookURL      string                      `gorm:"type:text" json:"webhook_url,omitempty"`
	ApiKey          string                      `gorm:"type:text" json:"api_key,omitempty"`
	BotPermissions  t.JSONB                     `json:"bot_permissions,omitempty"`
	Config          t.JSONB                     `json:"config" binding:"omitempty"`
	LastActivity    string                      `gorm:"type:timestamp;default:now()" json:"last_activity,omitempty" example:"2024-01-01T00:00:00Z"`
	UserID          string                      `gorm:"type:uuid;references:users(id)" json:"user_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	TargetTaskID    string                      `gorm:"type:uuid;references:mcp_sync_tasks(id)" json:"target_task_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
	CreatedAt       string                      `gorm:"type:timestamp;default:now()" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt       string                      `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy       string                      `gorm:"type:uuid;references:users(id)" json:"created_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	UpdatedBy       string                      `gorm:"type:uuid;references:users(id)" json:"updated_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
}

func NewTelegramModel() *TelegramModel {
	return &TelegramModel{
		ID:              uuid.New().String(),
		TelegramUserID:  "",
		Username:        "",
		FirstName:       "",
		LastName:        "",
		DisplayName:     "",
		PhoneNumber:     "",
		LanguageCode:    "",
		UserType:        TelegramUserTypeUser,
		Status:          TelegramStatusActive,
		IntegrationType: TelegramIntegrationTypeBot,
		ChatID:          "",
		ChannelID:       "",
		GroupID:         "",
		BotToken:        "",
		WebhookURL:      "",
		ApiKey:          "",
		BotPermissions:  t.JSONB{},
		Config:          t.JSONB{},
		LastActivity:    time.Now().Format(time.RFC3339),
		UserID:          "",
		TargetTaskID:    "",
		CreatedAt:       time.Now().Format(time.RFC3339),
		UpdatedAt:       time.Now().Format(time.RFC3339),
		CreatedBy:       "",
		UpdatedBy:       "",
	}
}

func (t *TelegramModel) TableName() string { return "mcp_telegram_integrations" }

// Basic getters and setters
func (t *TelegramModel) GetID() string                            { return t.ID }
func (t *TelegramModel) SetID(id string)                          { t.ID = id }
func (t *TelegramModel) GetTelegramUserID() string                { return t.TelegramUserID }
func (t *TelegramModel) SetTelegramUserID(telegramUserID string)  { t.TelegramUserID = telegramUserID }
func (t *TelegramModel) GetUsername() string                      { return t.Username }
func (t *TelegramModel) SetUsername(username string)              { t.Username = username }
func (t *TelegramModel) GetFirstName() string                     { return t.FirstName }
func (t *TelegramModel) SetFirstName(firstName string)            { t.FirstName = firstName }
func (t *TelegramModel) GetLastName() string                      { return t.LastName }
func (t *TelegramModel) SetLastName(lastName string)              { t.LastName = lastName }
func (t *TelegramModel) GetDisplayName() string                   { return t.DisplayName }
func (t *TelegramModel) SetDisplayName(displayName string)        { t.DisplayName = displayName }
func (t *TelegramModel) GetPhoneNumber() string                   { return t.PhoneNumber }
func (t *TelegramModel) SetPhoneNumber(phoneNumber string)        { t.PhoneNumber = phoneNumber }
func (t *TelegramModel) GetLanguageCode() string                  { return t.LanguageCode }
func (t *TelegramModel) SetLanguageCode(languageCode string)      { t.LanguageCode = languageCode }

// Type-specific getters and setters
func (t *TelegramModel) GetUserType() TelegramUserType                { return t.UserType }
func (t *TelegramModel) SetUserType(userType TelegramUserType)        { t.UserType = userType }
func (t *TelegramModel) GetStatus() TelegramStatus                    { return t.Status }
func (t *TelegramModel) SetStatus(status TelegramStatus)              { t.Status = status }
func (t *TelegramModel) GetIntegrationType() TelegramIntegrationType  { return t.IntegrationType }
func (t *TelegramModel) SetIntegrationType(integrationType TelegramIntegrationType) {
	t.IntegrationType = integrationType
}

// Telegram-specific getters and setters
func (t *TelegramModel) GetChatID() string               { return t.ChatID }
func (t *TelegramModel) SetChatID(chatID string)         { t.ChatID = chatID }
func (t *TelegramModel) GetChannelID() string            { return t.ChannelID }
func (t *TelegramModel) SetChannelID(channelID string)   { t.ChannelID = channelID }
func (t *TelegramModel) GetGroupID() string              { return t.GroupID }
func (t *TelegramModel) SetGroupID(groupID string)       { t.GroupID = groupID }
func (t *TelegramModel) GetBotToken() string             { return t.BotToken }
func (t *TelegramModel) SetBotToken(botToken string)     { t.BotToken = botToken }
func (t *TelegramModel) GetWebhookURL() string           { return t.WebhookURL }
func (t *TelegramModel) SetWebhookURL(webhookURL string) { t.WebhookURL = webhookURL }
func (t *TelegramModel) GetApiKey() string               { return t.ApiKey }
func (t *TelegramModel) SetApiKey(apiKey string)         { t.ApiKey = apiKey }
func (t *TelegramModel) GetBotPermissions() t.JSONB      { return t.BotPermissions }
func (t *TelegramModel) SetBotPermissions(permissions t.JSONB) { t.BotPermissions = permissions }
func (t *TelegramModel) GetConfig() t.JSONB              { return t.Config }
func (t *TelegramModel) SetConfig(config t.JSONB)        { t.Config = config }

func (t *TelegramModel) GetLastActivity() time.Time {
	lastActivity, _ := time.Parse(time.RFC3339, t.LastActivity)
	return lastActivity
}
func (t *TelegramModel) SetLastActivity(lastActivity time.Time) {
	t.LastActivity = lastActivity.Format(time.RFC3339)
}

// MCP integration getters and setters
func (t *TelegramModel) GetUserID() string                      { return t.UserID }
func (t *TelegramModel) SetUserID(userID string)                { t.UserID = userID }
func (t *TelegramModel) GetTargetTaskID() string                { return t.TargetTaskID }
func (t *TelegramModel) SetTargetTaskID(targetTaskID string)    { t.TargetTaskID = targetTaskID }

// Timestamp getters and setters
func (t *TelegramModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, t.CreatedAt)
	return createdAt
}
func (t *TelegramModel) SetCreatedAt(createdAt time.Time) {
	t.CreatedAt = createdAt.Format(time.RFC3339)
}
func (t *TelegramModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, t.UpdatedAt)
	return updatedAt
}
func (t *TelegramModel) SetUpdatedAt(updatedAt time.Time) {
	t.UpdatedAt = updatedAt.Format(time.RFC3339)
}
func (t *TelegramModel) GetCreatedBy() string             { return t.CreatedBy }
func (t *TelegramModel) SetCreatedBy(createdBy string)    { t.CreatedBy = createdBy }
func (t *TelegramModel) GetUpdatedBy() string             { return t.UpdatedBy }
func (t *TelegramModel) SetUpdatedBy(updatedBy string)    { t.UpdatedBy = updatedBy }

// Validation method
func (t *TelegramModel) Validate() error {
	if t.TelegramUserID == "" {
		return fmt.Errorf("telegram_user_id is required")
	}
	if t.FirstName == "" {
		return fmt.Errorf("first_name is required")
	}
	if t.UserType == "" {
		t.UserType = TelegramUserTypeUser
	}
	if t.Status == "" {
		t.Status = TelegramStatusActive
	}
	if t.IntegrationType == "" {
		t.IntegrationType = TelegramIntegrationTypeBot
	}
	return nil
}

// Sanitize method
func (t *TelegramModel) Sanitize() {
	t.UpdatedAt = time.Now().Format(time.RFC3339)
	if t.LastActivity == "" {
		t.LastActivity = time.Now().Format(time.RFC3339)
	}
	// Build display name if not provided
	if t.DisplayName == "" {
		if t.LastName != "" {
			t.DisplayName = fmt.Sprintf("%s %s", t.FirstName, t.LastName)
		} else {
			t.DisplayName = t.FirstName
		}
	}
}