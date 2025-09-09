package models

import (
	m "github.com/kubex-ecosystem/gdbase/internal/models/discord"
	"gorm.io/gorm"
)

type DiscordModel = m.DiscordModel
type DiscordService = m.IDiscordService
type DiscordRepo = m.IDiscordRepo

func NewDiscordService(discordRepo DiscordRepo) DiscordService {
	return m.NewDiscordService(discordRepo)
}

func NewDiscordRepo(db *gorm.DB) DiscordRepo {
	return m.NewDiscordRepo(db)
}
