package models

import (
	t "github.com/kubex-ecosystem/gdbase/internal/models/telegram"
	"gorm.io/gorm"
)

type TelegramModel = t.TelegramModel
type TelegramModelInterface = t.ITelegramModel
type TelegramService = t.ITelegramService
type TelegramRepo = t.ITelegramRepo

// Type aliases for enums

type TelegramUserType = t.TelegramUserType
type TelegramStatus = t.TelegramStatus
type TelegramIntegrationType = t.TelegramIntegrationType

// Enum constants
const (
	TelegramUserTypeBot     = t.TelegramUserTypeBot
	TelegramUserTypeUser    = t.TelegramUserTypeUser
	TelegramUserTypeChannel = t.TelegramUserTypeChannel
	TelegramUserTypeGroup   = t.TelegramUserTypeGroup
	TelegramUserTypeSystem  = t.TelegramUserTypeSystem

	TelegramStatusActive       = t.TelegramStatusActive
	TelegramStatusInactive     = t.TelegramStatusInactive
	TelegramStatusDisconnected = t.TelegramStatusDisconnected
	TelegramStatusError        = t.TelegramStatusError
	TelegramStatusBlocked      = t.TelegramStatusBlocked

	TelegramIntegrationTypeBot     = t.TelegramIntegrationTypeBot
	TelegramIntegrationTypeWebhook = t.TelegramIntegrationTypeWebhook
	TelegramIntegrationTypeAPI     = t.TelegramIntegrationTypeAPI
)

func NewTelegramService(telegramRepo TelegramRepo) TelegramService {
	return t.NewTelegramService(telegramRepo)
}

func NewTelegramRepo(db *gorm.DB) TelegramRepo {
	return t.NewTelegramRepository(db)
}

func NewTelegramModel() TelegramModelInterface {
	return t.NewTelegramModel()
}
