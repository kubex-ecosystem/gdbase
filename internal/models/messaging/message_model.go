package messaging

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

// MessageStatus represents the status of a message
type MessageStatus string

const (
	MessageStatusSent      MessageStatus = "SENT"
	MessageStatusDelivered MessageStatus = "DELIVERED"
	MessageStatusRead      MessageStatus = "READ"
	MessageStatusFailed    MessageStatus = "FAILED"
	MessageStatusPending   MessageStatus = "PENDING"
	MessageStatusDeleted   MessageStatus = "DELETED"
)

// MessageType represents the type of message content
type MessageType string

const (
	MessageTypeText     MessageType = "TEXT"
	MessageTypeImage    MessageType = "IMAGE"
	MessageTypeVideo    MessageType = "VIDEO"
	MessageTypeAudio    MessageType = "AUDIO"
	MessageTypeDocument MessageType = "DOCUMENT"
	MessageTypeLocation MessageType = "LOCATION"
	MessageTypeContact  MessageType = "CONTACT"
	MessageTypeSticker  MessageType = "STICKER"
	MessageTypeEmoji    MessageType = "EMOJI"
	MessageTypeFile     MessageType = "FILE"
	MessageTypeButton   MessageType = "BUTTON"
	MessageTypeList     MessageType = "LIST"
	MessageTypeTemplate MessageType = "TEMPLATE"
	MessageTypeSystem   MessageType = "SYSTEM"
)

// MessageDirection represents the direction of the message
type MessageDirection string

const (
	MessageDirectionInbound  MessageDirection = "INBOUND"
	MessageDirectionOutbound MessageDirection = "OUTBOUND"
)

// IMessageModel defines the interface for messages
type IMessageModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetConversationID() string
	SetConversationID(conversationID string)
	GetPlatform() Platform
	SetPlatform(platform Platform)
	GetPlatformMessageID() string
	SetPlatformMessageID(platformMessageID string)
	GetMessageType() MessageType
	SetMessageType(messageType MessageType)
	GetDirection() MessageDirection
	SetDirection(direction MessageDirection)
	GetStatus() MessageStatus
	SetStatus(status MessageStatus)
	GetSenderID() string
	SetSenderID(senderID string)
	GetSenderName() string
	SetSenderName(senderName string)
	GetRecipientID() string
	SetRecipientID(recipientID string)
	GetRecipientName() string
	SetRecipientName(recipientName string)
	GetContent() string
	SetContent(content string)
	GetAttachments() t.JSONBImpl
	SetAttachments(attachments t.JSONBImpl)
	GetMetadata() t.JSONBImpl
	SetMetadata(metadata t.JSONBImpl)
	GetReplyToMessageID() string
	SetReplyToMessageID(replyToMessageID string)
	GetThreadID() string
	SetThreadID(threadID string)
	GetTimestamp() time.Time
	SetTimestamp(timestamp time.Time)
	GetDeliveredAt() time.Time
	SetDeliveredAt(deliveredAt time.Time)
	GetReadAt() time.Time
	SetReadAt(readAt time.Time)
	GetUserID() string
	SetUserID(userID string)
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

// MessageModel represents a unified message entity across all platforms
type MessageModel struct {
	ID                string           `gorm:"type:uuid;primaryKey" json:"id"`
	ConversationID    string           `gorm:"type:uuid;not null;index;references:mcp_conversations(id)" json:"conversation_id"`
	Platform          Platform         `gorm:"type:text;not null;index" json:"platform" example:"DISCORD"`
	PlatformMessageID string           `gorm:"type:text;not null;index" json:"platform_message_id" example:"123456789012345678"`
	MessageType       MessageType      `gorm:"type:text;not null;default:'TEXT'" json:"message_type" example:"TEXT"`
	Direction         MessageDirection `gorm:"type:text;not null" json:"direction" example:"INBOUND"`
	Status            MessageStatus    `gorm:"type:text;not null;default:'SENT'" json:"status" example:"SENT"`
	SenderID          string           `gorm:"type:text;not null;index" json:"sender_id" example:"123456789"`
	SenderName        string           `gorm:"type:text" json:"sender_name,omitempty" example:"John Doe"`
	RecipientID       string           `gorm:"type:text;index" json:"recipient_id,omitempty" example:"987654321"`
	RecipientName     string           `gorm:"type:text" json:"recipient_name,omitempty" example:"Bot Support"`
	Content           string           `gorm:"type:text" json:"content,omitempty" example:"Hello, I need help with my order"`
	Attachments       t.JSONBImpl      `json:"attachments,omitempty"`
	Metadata          t.JSONBImpl      `json:"metadata,omitempty"`
	ReplyToMessageID  string           `gorm:"type:uuid;index" json:"reply_to_message_id,omitempty"`
	ThreadID          string           `gorm:"type:text;index" json:"thread_id,omitempty"`
	Timestamp         string           `gorm:"type:timestamp;not null;default:now()" json:"timestamp" example:"2024-01-01T00:00:00Z"`
	DeliveredAt       string           `gorm:"type:timestamp" json:"delivered_at,omitempty" example:"2024-01-01T00:00:01Z"`
	ReadAt            string           `gorm:"type:timestamp" json:"read_at,omitempty" example:"2024-01-01T00:00:02Z"`
	UserID            string           `gorm:"type:uuid;references:users(id)" json:"user_id,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	CreatedAt         string           `gorm:"type:timestamp;default:now()" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt         string           `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy         string           `gorm:"type:uuid;references:users(id)" json:"created_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	UpdatedBy         string           `gorm:"type:uuid;references:users(id)" json:"updated_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
}

func NewMessageModel() *MessageModel {
	return &MessageModel{
		ID:                uuid.New().String(),
		ConversationID:    "",
		Platform:          PlatformUnified,
		PlatformMessageID: "",
		MessageType:       MessageTypeText,
		Direction:         MessageDirectionInbound,
		Status:            MessageStatusSent,
		SenderID:          "",
		SenderName:        "",
		RecipientID:       "",
		RecipientName:     "",
		Content:           "",
		Attachments:       t.JSONBImpl{},
		Metadata:          t.JSONBImpl{},
		ReplyToMessageID:  "",
		ThreadID:          "",
		Timestamp:         time.Now().Format(time.RFC3339),
		DeliveredAt:       "",
		ReadAt:            "",
		UserID:            "",
		CreatedAt:         time.Now().Format(time.RFC3339),
		UpdatedAt:         time.Now().Format(time.RFC3339),
		CreatedBy:         "",
		UpdatedBy:         "",
	}
}

func (m *MessageModel) TableName() string { return "mcp_messages" }

// Basic getters and setters

func (m *MessageModel) GetID() string                           { return m.ID }
func (m *MessageModel) SetID(id string)                         { m.ID = id }
func (m *MessageModel) GetConversationID() string               { return m.ConversationID }
func (m *MessageModel) SetConversationID(conversationID string) { m.ConversationID = conversationID }
func (m *MessageModel) GetPlatform() Platform                   { return m.Platform }
func (m *MessageModel) SetPlatform(platform Platform)           { m.Platform = platform }
func (m *MessageModel) GetPlatformMessageID() string            { return m.PlatformMessageID }
func (m *MessageModel) SetPlatformMessageID(platformMessageID string) {
	m.PlatformMessageID = platformMessageID
}
func (m *MessageModel) GetMessageType() MessageType             { return m.MessageType }
func (m *MessageModel) SetMessageType(messageType MessageType)  { m.MessageType = messageType }
func (m *MessageModel) GetDirection() MessageDirection          { return m.Direction }
func (m *MessageModel) SetDirection(direction MessageDirection) { m.Direction = direction }
func (m *MessageModel) GetStatus() MessageStatus                { return m.Status }
func (m *MessageModel) SetStatus(status MessageStatus)          { m.Status = status }
func (m *MessageModel) GetSenderID() string                     { return m.SenderID }
func (m *MessageModel) SetSenderID(senderID string)             { m.SenderID = senderID }
func (m *MessageModel) GetSenderName() string                   { return m.SenderName }
func (m *MessageModel) SetSenderName(senderName string)         { m.SenderName = senderName }
func (m *MessageModel) GetRecipientID() string                  { return m.RecipientID }
func (m *MessageModel) SetRecipientID(recipientID string)       { m.RecipientID = recipientID }
func (m *MessageModel) GetRecipientName() string                { return m.RecipientName }
func (m *MessageModel) SetRecipientName(recipientName string)   { m.RecipientName = recipientName }
func (m *MessageModel) GetContent() string                      { return m.Content }
func (m *MessageModel) SetContent(content string)               { m.Content = content }
func (m *MessageModel) GetAttachments() t.JSONBImpl             { return m.Attachments }
func (m *MessageModel) SetAttachments(attachments t.JSONBImpl)  { m.Attachments = attachments }
func (m *MessageModel) GetMetadata() t.JSONBImpl                { return m.Metadata }
func (m *MessageModel) SetMetadata(metadata t.JSONBImpl)        { m.Metadata = metadata }
func (m *MessageModel) GetReplyToMessageID() string             { return m.ReplyToMessageID }
func (m *MessageModel) SetReplyToMessageID(replyToMessageID string) {
	m.ReplyToMessageID = replyToMessageID
}
func (m *MessageModel) GetThreadID() string         { return m.ThreadID }
func (m *MessageModel) SetThreadID(threadID string) { m.ThreadID = threadID }

func (m *MessageModel) GetTimestamp() time.Time {
	timestamp, _ := time.Parse(time.RFC3339, m.Timestamp)
	return timestamp
}
func (m *MessageModel) SetTimestamp(timestamp time.Time) {
	m.Timestamp = timestamp.Format(time.RFC3339)
}

func (m *MessageModel) GetDeliveredAt() time.Time {
	if m.DeliveredAt == "" {
		return time.Time{}
	}
	deliveredAt, _ := time.Parse(time.RFC3339, m.DeliveredAt)
	return deliveredAt
}
func (m *MessageModel) SetDeliveredAt(deliveredAt time.Time) {
	m.DeliveredAt = deliveredAt.Format(time.RFC3339)
}

func (m *MessageModel) GetReadAt() time.Time {
	if m.ReadAt == "" {
		return time.Time{}
	}
	readAt, _ := time.Parse(time.RFC3339, m.ReadAt)
	return readAt
}
func (m *MessageModel) SetReadAt(readAt time.Time) {
	m.ReadAt = readAt.Format(time.RFC3339)
}

func (m *MessageModel) GetUserID() string       { return m.UserID }
func (m *MessageModel) SetUserID(userID string) { m.UserID = userID }

// Timestamp getters and setters

func (m *MessageModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, m.CreatedAt)
	return createdAt
}
func (m *MessageModel) SetCreatedAt(createdAt time.Time) {
	m.CreatedAt = createdAt.Format(time.RFC3339)
}
func (m *MessageModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, m.UpdatedAt)
	return updatedAt
}
func (m *MessageModel) SetUpdatedAt(updatedAt time.Time) {
	m.UpdatedAt = updatedAt.Format(time.RFC3339)
}
func (m *MessageModel) GetCreatedBy() string          { return m.CreatedBy }
func (m *MessageModel) SetCreatedBy(createdBy string) { m.CreatedBy = createdBy }
func (m *MessageModel) GetUpdatedBy() string          { return m.UpdatedBy }
func (m *MessageModel) SetUpdatedBy(updatedBy string) { m.UpdatedBy = updatedBy }

// Validation method

func (m *MessageModel) Validate() error {
	if m.ConversationID == "" {
		return fmt.Errorf("conversation_id is required")
	}
	if m.Platform == "" {
		return fmt.Errorf("platform is required")
	}
	if m.PlatformMessageID == "" {
		return fmt.Errorf("platform_message_id is required")
	}
	if m.SenderID == "" {
		return fmt.Errorf("sender_id is required")
	}

	// Validate platform
	validPlatforms := map[Platform]bool{
		PlatformDiscord:  true,
		PlatformTelegram: true,
		PlatformWhatsApp: true,
		PlatformMeta:     true,
		PlatformUnified:  true,
	}
	if !validPlatforms[m.Platform] {
		return fmt.Errorf("invalid platform: %s", m.Platform)
	}

	// Validate message type
	validMessageTypes := map[MessageType]bool{
		MessageTypeText:     true,
		MessageTypeImage:    true,
		MessageTypeVideo:    true,
		MessageTypeAudio:    true,
		MessageTypeDocument: true,
		MessageTypeLocation: true,
		MessageTypeContact:  true,
		MessageTypeSticker:  true,
		MessageTypeEmoji:    true,
		MessageTypeFile:     true,
		MessageTypeButton:   true,
		MessageTypeList:     true,
		MessageTypeTemplate: true,
		MessageTypeSystem:   true,
	}
	if !validMessageTypes[m.MessageType] {
		return fmt.Errorf("invalid message type: %s", m.MessageType)
	}

	// Validate direction
	validDirections := map[MessageDirection]bool{
		MessageDirectionInbound:  true,
		MessageDirectionOutbound: true,
	}
	if !validDirections[m.Direction] {
		return fmt.Errorf("invalid direction: %s", m.Direction)
	}

	// Validate status
	validStatuses := map[MessageStatus]bool{
		MessageStatusSent:      true,
		MessageStatusDelivered: true,
		MessageStatusRead:      true,
		MessageStatusFailed:    true,
		MessageStatusPending:   true,
		MessageStatusDeleted:   true,
	}
	if !validStatuses[m.Status] {
		return fmt.Errorf("invalid status: %s", m.Status)
	}

	// Validate content requirements based on message type
	if m.MessageType == MessageTypeText && m.Content == "" {
		return fmt.Errorf("content is required for text messages")
	}

	return nil
}

// Sanitize method
func (m *MessageModel) Sanitize() {
	m.UpdatedAt = time.Now().Format(time.RFC3339)
	if m.Timestamp == "" {
		m.Timestamp = time.Now().Format(time.RFC3339)
	}

	// Auto-set sender name if empty
	if m.SenderName == "" && m.SenderID != "" {
		m.SenderName = fmt.Sprintf("User-%s", m.SenderID[len(m.SenderID)-6:])
	}

	// Auto-set recipient name if empty
	if m.RecipientName == "" && m.RecipientID != "" {
		m.RecipientName = fmt.Sprintf("User-%s", m.RecipientID[len(m.RecipientID)-6:])
	}
}
