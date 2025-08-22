// Package discord provides the model for managing Discord integrations in the MCP system.
package discord

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	t "github.com/rafa-mori/gdbase/types"
)

// DiscordUserType represents the type of Discord user/entity
type DiscordUserType string

const (
	DiscordUserTypeBot    DiscordUserType = "BOT"
	DiscordUserTypeUser   DiscordUserType = "USER"
	DiscordUserTypeSystem DiscordUserType = "SYSTEM"
)

// DiscordStatus represents the status of Discord integration
type DiscordStatus string

const (
	DiscordStatusActive       DiscordStatus = "ACTIVE"
	DiscordStatusInactive     DiscordStatus = "INACTIVE"
	DiscordStatusDisconnected DiscordStatus = "DISCONNECTED"
	DiscordStatusError        DiscordStatus = "ERROR"
)

// DiscordIntegrationType represents the type of Discord integration
type DiscordIntegrationType string

const (
	DiscordIntegrationTypeBot     DiscordIntegrationType = "BOT"
	DiscordIntegrationTypeWebhook DiscordIntegrationType = "WEBHOOK"
	DiscordIntegrationTypeOAuth2  DiscordIntegrationType = "OAUTH2"
)

// IDiscordModel defines the interface for Discord integrations
type IDiscordModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetDiscordUserID() string
	SetDiscordUserID(discordUserID string)
	GetUsername() string
	SetUsername(username string)
	GetDisplayName() string
	SetDisplayName(displayName string)
	GetDiscriminator() string
	SetDiscriminator(discriminator string)
	GetAvatar() string
	SetAvatar(avatar string)
	GetEmail() string
	SetEmail(email string)
	GetLocale() string
	SetLocale(locale string)
	GetUserType() DiscordUserType
	SetUserType(userType DiscordUserType)
	GetStatus() DiscordStatus
	SetStatus(status DiscordStatus)
	GetIntegrationType() DiscordIntegrationType
	SetIntegrationType(integrationType DiscordIntegrationType)
	GetGuildID() string
	SetGuildID(guildID string)
	GetChannelID() string
	SetChannelID(channelID string)
	GetAccessToken() string
	SetAccessToken(accessToken string)
	GetRefreshToken() string
	SetRefreshToken(refreshToken string)
	GetTokenExpiresAt() time.Time
	SetTokenExpiresAt(expiresAt time.Time)
	GetWebhookURL() string
	SetWebhookURL(webhookURL string)
	GetScopes() []string
	SetScopes(scopes []string)
	GetBotPermissions() int64
	SetBotPermissions(permissions int64)
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

// DiscordModel represents the Discord integration entity in the database
type DiscordModel struct {
	ID              string                 `gorm:"type:uuid;primaryKey" json:"id"`
	DiscordUserID   string                 `gorm:"type:text;not null;uniqueIndex" json:"discord_user_id" example:"123456789012345678"`
	Username        string                 `gorm:"type:text;not null" json:"username" example:"discord_user"`
	DisplayName     string                 `gorm:"type:text" json:"display_name,omitempty" example:"Discord User"`
	Discriminator   string                 `gorm:"type:text" json:"discriminator,omitempty" example:"0001"`
	Avatar          string                 `gorm:"type:text" json:"avatar,omitempty" example:"a_1234567890abcdef"`
	Email           string                 `gorm:"type:text" json:"email,omitempty" example:"user@example.com"`
	Locale          string                 `gorm:"type:text" json:"locale,omitempty" example:"en-US"`
	UserType        DiscordUserType        `gorm:"type:text;not null;default:'USER'" json:"user_type" example:"USER"`
	Status          DiscordStatus          `gorm:"type:text;not null;default:'ACTIVE'" json:"status" example:"ACTIVE"`
	IntegrationType DiscordIntegrationType `gorm:"type:text;not null;default:'BOT'" json:"integration_type" example:"BOT"`
	GuildID         string                 `gorm:"type:text" json:"guild_id,omitempty" example:"123456789012345678"`
	ChannelID       string                 `gorm:"type:text" json:"channel_id,omitempty" example:"123456789012345678"`
	AccessToken     string                 `gorm:"type:text" json:"access_token,omitempty"`
	RefreshToken    string                 `gorm:"type:text" json:"refresh_token,omitempty"`
	TokenExpiresAt  string                 `gorm:"type:timestamp" json:"token_expires_at,omitempty" example:"2024-01-01T00:00:00Z"`
	WebhookURL      string                 `gorm:"type:text" json:"webhook_url,omitempty"`
	Scopes          []string               `gorm:"type:text[]" json:"scopes,omitempty" example:"[\"bot\", \"identify\"]"`
	BotPermissions  int64                  `gorm:"type:bigint;default:0" json:"bot_permissions" example:"274877908992"`
	Config          t.JSONB                `json:"config" binding:"omitempty"`
	LastActivity    string                 `gorm:"type:timestamp;default:now()" json:"last_activity,omitempty" example:"2024-01-01T00:00:00Z"`
	UserID          string                 `gorm:"type:uuid;references:users(id)" json:"user_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	TargetTaskID    string                 `gorm:"type:uuid;references:mcp_sync_tasks(id)" json:"target_task_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
	CreatedAt       string                 `gorm:"type:timestamp;default:now()" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt       string                 `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy       string                 `gorm:"type:uuid;references:users(id)" json:"created_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	UpdatedBy       string                 `gorm:"type:uuid;references:users(id)" json:"updated_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
}

func NewDiscordModel() *DiscordModel {
	return &DiscordModel{
		ID:              uuid.New().String(),
		DiscordUserID:   "",
		Username:        "",
		DisplayName:     "",
		Discriminator:   "",
		Avatar:          "",
		Email:           "",
		Locale:          "",
		UserType:        DiscordUserTypeUser,
		Status:          DiscordStatusActive,
		IntegrationType: DiscordIntegrationTypeBot,
		GuildID:         "",
		ChannelID:       "",
		AccessToken:     "",
		RefreshToken:    "",
		TokenExpiresAt:  time.Now().Format(time.RFC3339),
		WebhookURL:      "",
		Scopes:          []string{},
		BotPermissions:  0,
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

func (d *DiscordModel) TableName() string { return "mcp_discord_integrations" }

// Basic getters and setters
func (d *DiscordModel) GetID() string                         { return d.ID }
func (d *DiscordModel) SetID(id string)                       { d.ID = id }
func (d *DiscordModel) GetDiscordUserID() string              { return d.DiscordUserID }
func (d *DiscordModel) SetDiscordUserID(discordUserID string) { d.DiscordUserID = discordUserID }
func (d *DiscordModel) GetUsername() string                   { return d.Username }
func (d *DiscordModel) SetUsername(username string)           { d.Username = username }
func (d *DiscordModel) GetDisplayName() string                { return d.DisplayName }
func (d *DiscordModel) SetDisplayName(displayName string)     { d.DisplayName = displayName }
func (d *DiscordModel) GetDiscriminator() string              { return d.Discriminator }
func (d *DiscordModel) SetDiscriminator(discriminator string) { d.Discriminator = discriminator }
func (d *DiscordModel) GetAvatar() string                     { return d.Avatar }
func (d *DiscordModel) SetAvatar(avatar string)               { d.Avatar = avatar }
func (d *DiscordModel) GetEmail() string                      { return d.Email }
func (d *DiscordModel) SetEmail(email string)                 { d.Email = email }
func (d *DiscordModel) GetLocale() string                     { return d.Locale }
func (d *DiscordModel) SetLocale(locale string)               { d.Locale = locale }

// Type-specific getters and setters
func (d *DiscordModel) GetUserType() DiscordUserType               { return d.UserType }
func (d *DiscordModel) SetUserType(userType DiscordUserType)       { d.UserType = userType }
func (d *DiscordModel) GetStatus() DiscordStatus                   { return d.Status }
func (d *DiscordModel) SetStatus(status DiscordStatus)             { d.Status = status }
func (d *DiscordModel) GetIntegrationType() DiscordIntegrationType { return d.IntegrationType }
func (d *DiscordModel) SetIntegrationType(integrationType DiscordIntegrationType) {
	d.IntegrationType = integrationType
}

// Discord-specific getters and setters
func (d *DiscordModel) GetGuildID() string                  { return d.GuildID }
func (d *DiscordModel) SetGuildID(guildID string)           { d.GuildID = guildID }
func (d *DiscordModel) GetChannelID() string                { return d.ChannelID }
func (d *DiscordModel) SetChannelID(channelID string)       { d.ChannelID = channelID }
func (d *DiscordModel) GetAccessToken() string              { return d.AccessToken }
func (d *DiscordModel) SetAccessToken(accessToken string)   { d.AccessToken = accessToken }
func (d *DiscordModel) GetRefreshToken() string             { return d.RefreshToken }
func (d *DiscordModel) SetRefreshToken(refreshToken string) { d.RefreshToken = refreshToken }

func (d *DiscordModel) GetTokenExpiresAt() time.Time {
	expiresAt, _ := time.Parse(time.RFC3339, d.TokenExpiresAt)
	return expiresAt
}
func (d *DiscordModel) SetTokenExpiresAt(expiresAt time.Time) {
	d.TokenExpiresAt = expiresAt.Format(time.RFC3339)
}

func (d *DiscordModel) GetWebhookURL() string               { return d.WebhookURL }
func (d *DiscordModel) SetWebhookURL(webhookURL string)     { d.WebhookURL = webhookURL }
func (d *DiscordModel) GetScopes() []string                 { return d.Scopes }
func (d *DiscordModel) SetScopes(scopes []string)           { d.Scopes = scopes }
func (d *DiscordModel) GetBotPermissions() int64            { return d.BotPermissions }
func (d *DiscordModel) SetBotPermissions(permissions int64) { d.BotPermissions = permissions }
func (d *DiscordModel) GetConfig() t.JSONB                  { return d.Config }
func (d *DiscordModel) SetConfig(config t.JSONB)            { d.Config = config }

func (d *DiscordModel) GetLastActivity() time.Time {
	lastActivity, _ := time.Parse(time.RFC3339, d.LastActivity)
	return lastActivity
}
func (d *DiscordModel) SetLastActivity(lastActivity time.Time) {
	d.LastActivity = lastActivity.Format(time.RFC3339)
}

// MCP integration getters and setters
func (d *DiscordModel) GetUserID() string                   { return d.UserID }
func (d *DiscordModel) SetUserID(userID string)             { d.UserID = userID }
func (d *DiscordModel) GetTargetTaskID() string             { return d.TargetTaskID }
func (d *DiscordModel) SetTargetTaskID(targetTaskID string) { d.TargetTaskID = targetTaskID }

// Timestamp getters and setters
func (d *DiscordModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, d.CreatedAt)
	return createdAt
}
func (d *DiscordModel) SetCreatedAt(createdAt time.Time) {
	d.CreatedAt = createdAt.Format(time.RFC3339)
}
func (d *DiscordModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, d.UpdatedAt)
	return updatedAt
}
func (d *DiscordModel) SetUpdatedAt(updatedAt time.Time) {
	d.UpdatedAt = updatedAt.Format(time.RFC3339)
}
func (d *DiscordModel) GetCreatedBy() string          { return d.CreatedBy }
func (d *DiscordModel) SetCreatedBy(createdBy string) { d.CreatedBy = createdBy }
func (d *DiscordModel) GetUpdatedBy() string          { return d.UpdatedBy }
func (d *DiscordModel) SetUpdatedBy(updatedBy string) { d.UpdatedBy = updatedBy }

// Validation method
func (d *DiscordModel) Validate() error {
	if d.DiscordUserID == "" {
		return fmt.Errorf("discord_user_id is required")
	}
	if d.Username == "" {
		return fmt.Errorf("username is required")
	}
	if d.UserType == "" {
		d.UserType = DiscordUserTypeUser
	}
	if d.Status == "" {
		d.Status = DiscordStatusActive
	}
	if d.IntegrationType == "" {
		d.IntegrationType = DiscordIntegrationTypeBot
	}
	return nil
}

// Sanitize method
func (d *DiscordModel) Sanitize() {
	d.UpdatedAt = time.Now().Format(time.RFC3339)
	if d.LastActivity == "" {
		d.LastActivity = time.Now().Format(time.RFC3339)
	}
}
