package models

import (
	"context"

	d "github.com/kubex-ecosystem/gdbase/internal/models/discord"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

type DiscordModel = d.DiscordModel
type DiscordModelInterface = d.IDiscordModel
type DiscordService = d.IDiscordService
type DiscordRepo = d.IDiscordRepo

// Type aliases for enums

type DiscordUserType = d.DiscordUserType
type DiscordStatus = d.DiscordStatus
type DiscordIntegrationType = d.DiscordIntegrationType

// Enum constants

const (
	DiscordUserTypeBot    = d.DiscordUserTypeBot
	DiscordUserTypeUser   = d.DiscordUserTypeUser
	DiscordUserTypeSystem = d.DiscordUserTypeSystem

	DiscordStatusActive       = d.DiscordStatusActive
	DiscordStatusInactive     = d.DiscordStatusInactive
	DiscordStatusDisconnected = d.DiscordStatusDisconnected
	DiscordStatusError        = d.DiscordStatusError

	DiscordIntegrationTypeBot     = d.DiscordIntegrationTypeBot
	DiscordIntegrationTypeWebhook = d.DiscordIntegrationTypeWebhook
	DiscordIntegrationTypeOAuth2  = d.DiscordIntegrationTypeOAuth2
)

func NewDiscordService(discordRepo DiscordRepo) DiscordService {
	return d.NewDiscordService(discordRepo)
}

func NewDiscordRepo(ctx context.Context, dbService *svc.DBServiceImpl) DiscordRepo {
	return d.NewDiscordRepo(ctx, dbService)
}

func NewDiscordModel() DiscordModelInterface {
	return d.NewDiscordModel()
}
