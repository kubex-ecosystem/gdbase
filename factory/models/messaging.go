package models

import (
	"context"

	m "github.com/kubex-ecosystem/gdbase/internal/models/messaging"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

// Conversation types

type ConversationModel = m.ConversationModel
type ConversationModelInterface = m.IConversationModel
type ConversationRepo = m.IConversationRepo

// Message types

type MessageModel = m.MessageModel
type MessageModelInterface = m.IMessageModel

// Enum type aliases

type ConversationStatus = m.ConversationStatus
type ConversationType = m.ConversationType
type Platform = m.Platform
type MessageStatus = m.MessageStatus
type MessageType = m.MessageType
type MessageDirection = m.MessageDirection

// Platform constants

const (
	PlatformDiscord  = m.PlatformDiscord
	PlatformTelegram = m.PlatformTelegram
	PlatformWhatsApp = m.PlatformWhatsApp
	PlatformMeta     = m.PlatformMeta
	PlatformUnified  = m.PlatformUnified
)

// Conversation status constants

const (
	ConversationStatusActive   = m.ConversationStatusActive
	ConversationStatusInactive = m.ConversationStatusInactive
	ConversationStatusArchived = m.ConversationStatusArchived
	ConversationStatusBlocked  = m.ConversationStatusBlocked
	ConversationStatusPending  = m.ConversationStatusPending
)

// Conversation type constants

const (
	ConversationTypePrivate = m.ConversationTypePrivate
	ConversationTypeGroup   = m.ConversationTypeGroup
	ConversationTypeChannel = m.ConversationTypeChannel
	ConversationTypeBot     = m.ConversationTypeBot
	ConversationTypeSupport = m.ConversationTypeSupport
)

// Message status constants

const (
	MessageStatusSent      = m.MessageStatusSent
	MessageStatusDelivered = m.MessageStatusDelivered
	MessageStatusRead      = m.MessageStatusRead
	MessageStatusFailed    = m.MessageStatusFailed
	MessageStatusPending   = m.MessageStatusPending
	MessageStatusDeleted   = m.MessageStatusDeleted
)

// Message type constants

const (
	MessageTypeText     = m.MessageTypeText
	MessageTypeImage    = m.MessageTypeImage
	MessageTypeVideo    = m.MessageTypeVideo
	MessageTypeAudio    = m.MessageTypeAudio
	MessageTypeDocument = m.MessageTypeDocument
	MessageTypeLocation = m.MessageTypeLocation
	MessageTypeContact  = m.MessageTypeContact
	MessageTypeSticker  = m.MessageTypeSticker
	MessageTypeEmoji    = m.MessageTypeEmoji
	MessageTypeFile     = m.MessageTypeFile
	MessageTypeButton   = m.MessageTypeButton
	MessageTypeList     = m.MessageTypeList
	MessageTypeTemplate = m.MessageTypeTemplate
	MessageTypeSystem   = m.MessageTypeSystem
)

// Message direction constants

const (
	MessageDirectionInbound  = m.MessageDirectionInbound
	MessageDirectionOutbound = m.MessageDirectionOutbound
)

// Conversation functions

func NewConversationRepo(ctx context.Context, dbService *svc.DBServiceImpl) ConversationRepo {
	return m.NewConversationRepository(ctx, dbService)
}

func NewConversationModel() ConversationModelInterface {
	return m.NewConversationModel()
}

// Message functions

func NewMessageModel() MessageModelInterface {
	return m.NewMessageModel()
}
