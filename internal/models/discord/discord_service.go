// Package discord provides the service for managing Discord integrations in the MCP system.
package discord

import (
	"errors"
	"fmt"
	"time"

	is "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type IDiscordService interface {
	CreateDiscordIntegration(discord IDiscordModel) (IDiscordModel, error)
	GetDiscordIntegrationByID(id string) (IDiscordModel, error)
	UpdateDiscordIntegration(discord IDiscordModel) (IDiscordModel, error)
	DeleteDiscordIntegration(id string) error
	ListDiscordIntegrations() ([]IDiscordModel, error)
	GetDiscordIntegrationByDiscordUserID(discordUserID string) (IDiscordModel, error)
	GetDiscordIntegrationsByUserID(userID string) ([]IDiscordModel, error)
	GetDiscordIntegrationsByGuildID(guildID string) ([]IDiscordModel, error)
	GetDiscordIntegrationsByChannelID(channelID string) ([]IDiscordModel, error)
	GetDiscordIntegrationsByStatus(status DiscordStatus) ([]IDiscordModel, error)
	GetDiscordIntegrationsByType(integrationType DiscordIntegrationType) ([]IDiscordModel, error)
	GetDiscordIntegrationsByTargetTaskID(targetTaskID string) ([]IDiscordModel, error)
	GetActiveDiscordIntegrations() ([]IDiscordModel, error)
	GetDiscordBotIntegrations() ([]IDiscordModel, error)
	GetDiscordWebhookIntegrations() ([]IDiscordModel, error)
	GetDiscordOAuth2Integrations() ([]IDiscordModel, error)

	// Activity and status management
	UpdateLastActivity(id string) error
	SetDiscordIntegrationStatus(id string, status DiscordStatus) error
	SetDiscordIntegrationActive(id string) error
	SetDiscordIntegrationInactive(id string) error
	SetDiscordIntegrationDisconnected(id string) error
	SetDiscordIntegrationError(id string) error

	// Token management
	UpdateAccessToken(id string, accessToken string, refreshToken string, expiresAt time.Time) error
	RefreshToken(id string) error
	IsTokenExpired(id string) (bool, error)

	// Configuration management
	UpdateDiscordIntegrationConfig(id string, config t.JSONB) error
	GetDiscordIntegrationConfig(id string) (t.JSONB, error)

	// Upsert operations
	UpsertDiscordIntegrationByDiscordUserID(discordUserID string, discord IDiscordModel) (IDiscordModel, error)

	GetContextDBService() is.IDBService
}

type DiscordService struct {
	repo IDiscordRepo
}

func NewDiscordService(repo IDiscordRepo) IDiscordService {
	return &DiscordService{repo: repo}
}

func (ds *DiscordService) CreateDiscordIntegration(discord IDiscordModel) (IDiscordModel, error) {
	if discord.GetDiscordUserID() == "" {
		return nil, errors.New("missing required field: discord_user_id is required")
	}

	if discord.GetUsername() == "" {
		return nil, errors.New("missing required field: username is required")
	}

	// Validate the model
	if err := discord.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	createdDiscord, err := ds.repo.Create(discord)
	if err != nil {
		return nil, fmt.Errorf("error creating Discord integration: %w", err)
	}
	return createdDiscord, nil
}

func (ds *DiscordService) GetDiscordIntegrationByID(id string) (IDiscordModel, error) {
	discord, err := ds.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error fetching Discord integration: %w", err)
	}
	return discord, nil
}

func (ds *DiscordService) UpdateDiscordIntegration(discord IDiscordModel) (IDiscordModel, error) {
	// Validate the model before updating
	if err := discord.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	updatedDiscord, err := ds.repo.Update(discord)
	if err != nil {
		return nil, fmt.Errorf("error updating Discord integration: %w", err)
	}
	return updatedDiscord, nil
}

func (ds *DiscordService) DeleteDiscordIntegration(id string) error {
	err := ds.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting Discord integration: %w", err)
	}
	return nil
}

func (ds *DiscordService) ListDiscordIntegrations() ([]IDiscordModel, error) {
	discords, err := ds.repo.FindAll("")
	if err != nil {
		return nil, fmt.Errorf("error listing Discord integrations: %w", err)
	}
	return discords, nil
}

func (ds *DiscordService) GetDiscordIntegrationByDiscordUserID(discordUserID string) (IDiscordModel, error) {
	discord, err := ds.repo.FindOne("discord_user_id = ?", discordUserID)
	if err != nil {
		return nil, fmt.Errorf("error fetching Discord integration by Discord User ID: %w", err)
	}
	return discord, nil
}

func (ds *DiscordService) GetDiscordIntegrationsByUserID(userID string) ([]IDiscordModel, error) {
	discords, err := ds.repo.FindAll("user_id = ?", userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching Discord integrations by User ID: %w", err)
	}
	return discords, nil
}

func (ds *DiscordService) GetDiscordIntegrationsByGuildID(guildID string) ([]IDiscordModel, error) {
	discords, err := ds.repo.FindAll("guild_id = ?", guildID)
	if err != nil {
		return nil, fmt.Errorf("error fetching Discord integrations by Guild ID: %w", err)
	}
	return discords, nil
}

func (ds *DiscordService) GetDiscordIntegrationsByChannelID(channelID string) ([]IDiscordModel, error) {
	discords, err := ds.repo.FindAll("channel_id = ?", channelID)
	if err != nil {
		return nil, fmt.Errorf("error fetching Discord integrations by Channel ID: %w", err)
	}
	return discords, nil
}

func (ds *DiscordService) GetDiscordIntegrationsByStatus(status DiscordStatus) ([]IDiscordModel, error) {
	discords, err := ds.repo.FindAll("status = ?", status)
	if err != nil {
		return nil, fmt.Errorf("error fetching Discord integrations by status: %w", err)
	}
	return discords, nil
}

func (ds *DiscordService) GetDiscordIntegrationsByType(integrationType DiscordIntegrationType) ([]IDiscordModel, error) {
	discords, err := ds.repo.FindAll("integration_type = ?", integrationType)
	if err != nil {
		return nil, fmt.Errorf("error fetching Discord integrations by type: %w", err)
	}
	return discords, nil
}

func (ds *DiscordService) GetDiscordIntegrationsByTargetTaskID(targetTaskID string) ([]IDiscordModel, error) {
	discords, err := ds.repo.FindAll("target_task_id = ?", targetTaskID)
	if err != nil {
		return nil, fmt.Errorf("error fetching Discord integrations by target task ID: %w", err)
	}
	return discords, nil
}

func (ds *DiscordService) GetActiveDiscordIntegrations() ([]IDiscordModel, error) {
	return ds.GetDiscordIntegrationsByStatus(DiscordStatusActive)
}

func (ds *DiscordService) GetDiscordBotIntegrations() ([]IDiscordModel, error) {
	return ds.GetDiscordIntegrationsByType(DiscordIntegrationTypeBot)
}

func (ds *DiscordService) GetDiscordWebhookIntegrations() ([]IDiscordModel, error) {
	return ds.GetDiscordIntegrationsByType(DiscordIntegrationTypeWebhook)
}

func (ds *DiscordService) GetDiscordOAuth2Integrations() ([]IDiscordModel, error) {
	return ds.GetDiscordIntegrationsByType(DiscordIntegrationTypeOAuth2)
}

// Activity and status management

func (ds *DiscordService) UpdateLastActivity(id string) error {
	discord, err := ds.GetDiscordIntegrationByID(id)
	if err != nil {
		return err
	}

	discord.SetLastActivity(time.Now())
	_, err = ds.UpdateDiscordIntegration(discord)
	return err
}

func (ds *DiscordService) SetDiscordIntegrationStatus(id string, status DiscordStatus) error {
	discord, err := ds.GetDiscordIntegrationByID(id)
	if err != nil {
		return err
	}

	discord.SetStatus(status)
	discord.SetLastActivity(time.Now())
	_, err = ds.UpdateDiscordIntegration(discord)
	return err
}

func (ds *DiscordService) SetDiscordIntegrationActive(id string) error {
	return ds.SetDiscordIntegrationStatus(id, DiscordStatusActive)
}

func (ds *DiscordService) SetDiscordIntegrationInactive(id string) error {
	return ds.SetDiscordIntegrationStatus(id, DiscordStatusInactive)
}

func (ds *DiscordService) SetDiscordIntegrationDisconnected(id string) error {
	return ds.SetDiscordIntegrationStatus(id, DiscordStatusDisconnected)
}

func (ds *DiscordService) SetDiscordIntegrationError(id string) error {
	return ds.SetDiscordIntegrationStatus(id, DiscordStatusError)
}

// Token management

func (ds *DiscordService) UpdateAccessToken(id string, accessToken string, refreshToken string, expiresAt time.Time) error {
	discord, err := ds.GetDiscordIntegrationByID(id)
	if err != nil {
		return err
	}

	discord.SetAccessToken(accessToken)
	discord.SetRefreshToken(refreshToken)
	discord.SetTokenExpiresAt(expiresAt)
	discord.SetLastActivity(time.Now())
	_, err = ds.UpdateDiscordIntegration(discord)
	return err
}

func (ds *DiscordService) RefreshToken(id string) error {
	// This would need to be implemented with Discord OAuth2 token refresh logic
	// For now, just update the last activity
	return ds.UpdateLastActivity(id)
}

func (ds *DiscordService) IsTokenExpired(id string) (bool, error) {
	discord, err := ds.GetDiscordIntegrationByID(id)
	if err != nil {
		return false, err
	}

	expiresAt := discord.GetTokenExpiresAt()
	return time.Now().After(expiresAt), nil
}

// Configuration management

func (ds *DiscordService) UpdateDiscordIntegrationConfig(id string, config t.JSONB) error {
	discord, err := ds.GetDiscordIntegrationByID(id)
	if err != nil {
		return err
	}

	discord.SetConfig(config)
	discord.SetLastActivity(time.Now())
	_, err = ds.UpdateDiscordIntegration(discord)
	return err
}

func (ds *DiscordService) GetDiscordIntegrationConfig(id string) (t.JSONB, error) {
	discord, err := ds.GetDiscordIntegrationByID(id)
	if err != nil {
		return nil, err
	}
	return discord.GetConfig(), nil
}

// Upsert operations

func (ds *DiscordService) UpsertDiscordIntegrationByDiscordUserID(discordUserID string, discord IDiscordModel) (IDiscordModel, error) {
	// Try to find existing integration by Discord User ID
	existing, err := ds.repo.FindOne("discord_user_id = ?", discordUserID)
	if err != nil {
		// If not found, create new
		discord.SetDiscordUserID(discordUserID)
		return ds.CreateDiscordIntegration(discord)
	}

	// If found, update existing
	existing.SetUsername(discord.GetUsername())
	existing.SetDisplayName(discord.GetDisplayName())
	existing.SetDiscriminator(discord.GetDiscriminator())
	existing.SetAvatar(discord.GetAvatar())
	existing.SetEmail(discord.GetEmail())
	existing.SetLocale(discord.GetLocale())
	existing.SetUserType(discord.GetUserType())
	existing.SetStatus(discord.GetStatus())
	existing.SetIntegrationType(discord.GetIntegrationType())
	existing.SetGuildID(discord.GetGuildID())
	existing.SetChannelID(discord.GetChannelID())
	existing.SetScopes(discord.GetScopes())
	existing.SetBotPermissions(discord.GetBotPermissions())
	existing.SetConfig(discord.GetConfig())
	existing.SetLastActivity(time.Now())

	if discord.GetUserID() != "" {
		existing.SetUserID(discord.GetUserID())
	}
	if discord.GetTargetTaskID() != "" {
		existing.SetTargetTaskID(discord.GetTargetTaskID())
	}
	if discord.GetUpdatedBy() != "" {
		existing.SetUpdatedBy(discord.GetUpdatedBy())
	}

	return ds.UpdateDiscordIntegration(existing)
}

func (ds *DiscordService) GetContextDBService() is.IDBService {
	return ds.repo.GetContextDBService()
}
