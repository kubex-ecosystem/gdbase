package models

import (
	"context"

	t "github.com/kubex-ecosystem/gdbase/internal/models/telegram"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
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

func NewTelegramRepo(ctx context.Context, dbService *svc.DBServiceImpl) TelegramRepo {
	return t.NewTelegramRepository(ctx, dbService)
}

func NewTelegramModel() TelegramModelInterface {
	return t.NewTelegramModel()
}
