// Package messaging provides unified models for managing conversations across all bot platforms (Discord, Telegram, WhatsApp).
package messaging

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/types"
)

// ConversationStatus represents the status of a conversation
type ConversationStatus string

const (
	ConversationStatusActive   ConversationStatus = "ACTIVE"
	ConversationStatusInactive ConversationStatus = "INACTIVE"
	ConversationStatusArchived ConversationStatus = "ARCHIVED"
	ConversationStatusBlocked  ConversationStatus = "BLOCKED"
	ConversationStatusPending  ConversationStatus = "PENDING"
)

// ConversationType represents the type of conversation
type ConversationType string

const (
	ConversationTypePrivate ConversationType = "PRIVATE"
	ConversationTypeGroup   ConversationType = "GROUP"
	ConversationTypeChannel ConversationType = "CHANNEL"
	ConversationTypeBot     ConversationType = "BOT"
	ConversationTypeSupport ConversationType = "SUPPORT"
)

// Platform represents the bot platform
type Platform string

const (
	PlatformDiscord  Platform = "DISCORD"
	PlatformTelegram Platform = "TELEGRAM"
	PlatformWhatsApp Platform = "WHATSAPP"
	PlatformMeta     Platform = "META"
	PlatformUnified  Platform = "UNIFIED"
)

// IConversationModel defines the interface for conversations
type IConversationModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetPlatform() Platform
	SetPlatform(platform Platform)
	GetPlatformConversationID() string
	SetPlatformConversationID(platformConversationID string)
	GetIntegrationID() string
	SetIntegrationID(integrationID string)
	GetTitle() string
	SetTitle(title string)
	GetDescription() string
	SetDescription(description string)
	GetConversationType() ConversationType
	SetConversationType(conversationType ConversationType)
	GetStatus() ConversationStatus
	SetStatus(status ConversationStatus)
	GetParticipants() t.JSONB
	SetParticipants(participants t.JSONB)
	GetMetadata() t.JSONB
	SetMetadata(metadata t.JSONB)
	GetLastMessageID() string
	SetLastMessageID(lastMessageID string)
	GetLastMessageAt() time.Time
	SetLastMessageAt(lastMessageAt time.Time)
	GetMessageCount() int64
	SetMessageCount(messageCount int64)
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

// ConversationModel represents a unified conversation entity across all platforms
type ConversationModel struct {
	ID                     string             `gorm:"type:uuid;primaryKey" json:"id"`
	Platform               Platform           `gorm:"type:text;not null;index" json:"platform" example:"DISCORD"`
	PlatformConversationID string             `gorm:"type:text;not null;index" json:"platform_conversation_id" example:"123456789012345678"`
	IntegrationID          string             `gorm:"type:uuid;not null;index" json:"integration_id" example:"123e4567-e89b-12d3-a456-426614174001"`
	Title                  string             `gorm:"type:text" json:"title,omitempty" example:"Support Chat #1234"`
	Description            string             `gorm:"type:text" json:"description,omitempty" example:"Customer support conversation"`
	ConversationType       ConversationType   `gorm:"type:text;not null;default:'PRIVATE'" json:"conversation_type" example:"PRIVATE"`
	Status                 ConversationStatus `gorm:"type:text;not null;default:'ACTIVE'" json:"status" example:"ACTIVE"`
	Participants           t.JSONB            `json:"participants,omitempty"`
	Metadata               t.JSONB            `json:"metadata,omitempty"`
	LastMessageID          string             `gorm:"type:uuid;index" json:"last_message_id,omitempty"`
	LastMessageAt          string             `gorm:"type:timestamp;default:now()" json:"last_message_at,omitempty" example:"2024-01-01T00:00:00Z"`
	MessageCount           int64              `gorm:"type:bigint;default:0" json:"message_count" example:"42"`
	UserID                 string             `gorm:"type:uuid;references:users(id)" json:"user_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	TargetTaskID           string             `gorm:"type:uuid;references:mcp_sync_tasks(id)" json:"target_task_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
	CreatedAt              string             `gorm:"type:timestamp;default:now()" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt              string             `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy              string             `gorm:"type:uuid;references:users(id)" json:"created_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	UpdatedBy              string             `gorm:"type:uuid;references:users(id)" json:"updated_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
}

func NewConversationModel() *ConversationModel {
	return &ConversationModel{
		ID:                     uuid.New().String(),
		Platform:               PlatformUnified,
		PlatformConversationID: "",
		IntegrationID:          "",
		Title:                  "",
		Description:            "",
		ConversationType:       ConversationTypePrivate,
		Status:                 ConversationStatusActive,
		Participants:           t.JSONB{},
		Metadata:               t.JSONB{},
		LastMessageID:          "",
		LastMessageAt:          time.Now().Format(time.RFC3339),
		MessageCount:           0,
		UserID:                 "",
		TargetTaskID:           "",
		CreatedAt:              time.Now().Format(time.RFC3339),
		UpdatedAt:              time.Now().Format(time.RFC3339),
		CreatedBy:              "",
		UpdatedBy:              "",
	}
}

func (c *ConversationModel) TableName() string { return "mcp_conversations" }

// Basic getters and setters
func (c *ConversationModel) GetID() string                                                  { return c.ID }
func (c *ConversationModel) SetID(id string)                                               { c.ID = id }
func (c *ConversationModel) GetPlatform() Platform                                         { return c.Platform }
func (c *ConversationModel) SetPlatform(platform Platform)                                 { c.Platform = platform }
func (c *ConversationModel) GetPlatformConversationID() string                             { return c.PlatformConversationID }
func (c *ConversationModel) SetPlatformConversationID(platformConversationID string)      { c.PlatformConversationID = platformConversationID }
func (c *ConversationModel) GetIntegrationID() string                                      { return c.IntegrationID }
func (c *ConversationModel) SetIntegrationID(integrationID string)                         { c.IntegrationID = integrationID }
func (c *ConversationModel) GetTitle() string                                              { return c.Title }
func (c *ConversationModel) SetTitle(title string)                                         { c.Title = title }
func (c *ConversationModel) GetDescription() string                                        { return c.Description }
func (c *ConversationModel) SetDescription(description string)                             { c.Description = description }
func (c *ConversationModel) GetConversationType() ConversationType                         { return c.ConversationType }
func (c *ConversationModel) SetConversationType(conversationType ConversationType)         { c.ConversationType = conversationType }
func (c *ConversationModel) GetStatus() ConversationStatus                                 { return c.Status }
func (c *ConversationModel) SetStatus(status ConversationStatus)                           { c.Status = status }
func (c *ConversationModel) GetParticipants() t.JSONB                                      { return c.Participants }
func (c *ConversationModel) SetParticipants(participants t.JSONB)                          { c.Participants = participants }
func (c *ConversationModel) GetMetadata() t.JSONB                                          { return c.Metadata }
func (c *ConversationModel) SetMetadata(metadata t.JSONB)                                  { c.Metadata = metadata }
func (c *ConversationModel) GetLastMessageID() string                                      { return c.LastMessageID }
func (c *ConversationModel) SetLastMessageID(lastMessageID string)                         { c.LastMessageID = lastMessageID }

func (c *ConversationModel) GetLastMessageAt() time.Time {
	lastMessageAt, _ := time.Parse(time.RFC3339, c.LastMessageAt)
	return lastMessageAt
}
func (c *ConversationModel) SetLastMessageAt(lastMessageAt time.Time) {
	c.LastMessageAt = lastMessageAt.Format(time.RFC3339)
}

func (c *ConversationModel) GetMessageCount() int64                        { return c.MessageCount }
func (c *ConversationModel) SetMessageCount(messageCount int64)            { c.MessageCount = messageCount }
func (c *ConversationModel) GetUserID() string                             { return c.UserID }
func (c *ConversationModel) SetUserID(userID string)                       { c.UserID = userID }
func (c *ConversationModel) GetTargetTaskID() string                       { return c.TargetTaskID }
func (c *ConversationModel) SetTargetTaskID(targetTaskID string)           { c.TargetTaskID = targetTaskID }

// Timestamp getters and setters
func (c *ConversationModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, c.CreatedAt)
	return createdAt
}
func (c *ConversationModel) SetCreatedAt(createdAt time.Time) {
	c.CreatedAt = createdAt.Format(time.RFC3339)
}
func (c *ConversationModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, c.UpdatedAt)
	return updatedAt
}
func (c *ConversationModel) SetUpdatedAt(updatedAt time.Time) {
	c.UpdatedAt = updatedAt.Format(time.RFC3339)
}
func (c *ConversationModel) GetCreatedBy() string                { return c.CreatedBy }
func (c *ConversationModel) SetCreatedBy(createdBy string)       { c.CreatedBy = createdBy }
func (c *ConversationModel) GetUpdatedBy() string                { return c.UpdatedBy }
func (c *ConversationModel) SetUpdatedBy(updatedBy string)       { c.UpdatedBy = updatedBy }

// Validation method
func (c *ConversationModel) Validate() error {
	if c.Platform == "" {
		return fmt.Errorf("platform is required")
	}
	if c.PlatformConversationID == "" {
		return fmt.Errorf("platform_conversation_id is required")
	}
	if c.IntegrationID == "" {
		return fmt.Errorf("integration_id is required")
	}

	// Validate platform
	validPlatforms := map[Platform]bool{
		PlatformDiscord:  true,
		PlatformTelegram: true,
		PlatformWhatsApp: true,
		PlatformMeta:     true,
		PlatformUnified:  true,
	}
	if !validPlatforms[c.Platform] {
		return fmt.Errorf("invalid platform: %s", c.Platform)
	}

	// Validate conversation type
	validConversationTypes := map[ConversationType]bool{
		ConversationTypePrivate: true,
		ConversationTypeGroup:   true,
		ConversationTypeChannel: true,
		ConversationTypeBot:     true,
		ConversationTypeSupport: true,
	}
	if !validConversationTypes[c.ConversationType] {
		return fmt.Errorf("invalid conversation type: %s", c.ConversationType)
	}

	// Validate status
	validStatuses := map[ConversationStatus]bool{
		ConversationStatusActive:   true,
		ConversationStatusInactive: true,
		ConversationStatusArchived: true,
		ConversationStatusBlocked:  true,
		ConversationStatusPending:  true,
	}
	if !validStatuses[c.Status] {
		return fmt.Errorf("invalid status: %s", c.Status)
	}

	return nil
}

// Sanitize method
func (c *ConversationModel) Sanitize() {
	c.UpdatedAt = time.Now().Format(time.RFC3339)
	if c.LastMessageAt == "" {
		c.LastMessageAt = time.Now().Format(time.RFC3339)
	}

	// Generate title if not provided
	if c.Title == "" {
		switch c.ConversationType {
		case ConversationTypePrivate:
			c.Title = fmt.Sprintf("Private Chat - %s", c.Platform)
		case ConversationTypeGroup:
			c.Title = fmt.Sprintf("Group Chat - %s", c.Platform)
		case ConversationTypeChannel:
			c.Title = fmt.Sprintf("Channel - %s", c.Platform)
		case ConversationTypeBot:
			c.Title = fmt.Sprintf("Bot Chat - %s", c.Platform)
		case ConversationTypeSupport:
			c.Title = fmt.Sprintf("Support Chat - %s", c.Platform)
		default:
			c.Title = fmt.Sprintf("Conversation - %s", c.Platform)
		}
	}
}