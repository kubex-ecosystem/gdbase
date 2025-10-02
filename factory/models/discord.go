package models

import (
	d "github.com/kubex-ecosystem/gdbase/internal/models/discord"
	"gorm.io/gorm"
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

func NewDiscordRepo(db *gorm.DB) DiscordRepo {
	return d.NewDiscordRepo(db)
}

func NewDiscordModel() DiscordModelInterface {
	return d.NewDiscordModel()
}
