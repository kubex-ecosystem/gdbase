package preferences

import (
	"errors"
	"fmt"

	is "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type IPreferencesService interface {
	CreatePreferences(preferences IPreferencesModel) (IPreferencesModel, error)
	GetPreferencesByID(id string) (IPreferencesModel, error)
	UpdatePreferences(preferences IPreferencesModel) (IPreferencesModel, error)
	DeletePreferences(id string) error
	ListPreferences() ([]IPreferencesModel, error)
	GetPreferencesByScope(scope string) (IPreferencesModel, error)
	GetPreferencesByUserID(userID string) ([]IPreferencesModel, error)
	UpsertPreferencesByScope(scope string, config t.JSONB, userID string) (IPreferencesModel, error)
	GetContextDBService() *is.DBServiceImpl
}

type PreferencesService struct {
	repo IPreferencesRepo
}

func NewPreferencesService(repo IPreferencesRepo) IPreferencesService {
	return &PreferencesService{repo: repo}
}

func (ps *PreferencesService) CreatePreferences(preferences IPreferencesModel) (IPreferencesModel, error) {
	if preferences.GetScope() == "" {
		return nil, errors.New("missing required field: scope is required")
	}

	if preferences.GetConfig() == nil {
		return nil, errors.New("missing required field: config cannot be nil")
	}

	// Validate the model
	if err := preferences.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	createdPreferences, err := ps.repo.Create(preferences)
	if err != nil {
		return nil, fmt.Errorf("error creating preferences: %w", err)
	}
	return createdPreferences, nil
}

func (ps *PreferencesService) GetPreferencesByID(id string) (IPreferencesModel, error) {
	preferences, err := ps.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error fetching preferences: %w", err)
	}
	return preferences, nil
}

func (ps *PreferencesService) UpdatePreferences(preferences IPreferencesModel) (IPreferencesModel, error) {
	// Validate the model before updating
	if err := preferences.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	updatedPreferences, err := ps.repo.Update(preferences)
	if err != nil {
		return nil, fmt.Errorf("error updating preferences: %w", err)
	}
	return updatedPreferences, nil
}

func (ps *PreferencesService) DeletePreferences(id string) error {
	err := ps.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting preferences: %w", err)
	}
	return nil
}

func (ps *PreferencesService) ListPreferences() ([]IPreferencesModel, error) {
	preferences, err := ps.repo.FindAll("")
	if err != nil {
		return nil, fmt.Errorf("error listing preferences: %w", err)
	}
	return preferences, nil
}

func (ps *PreferencesService) GetPreferencesByScope(scope string) (IPreferencesModel, error) {
	preferences, err := ps.repo.FindOne("scope = ?", scope)
	if err != nil {
		return nil, fmt.Errorf("error fetching preferences by scope: %w", err)
	}
	return preferences, nil
}

func (ps *PreferencesService) GetPreferencesByUserID(userID string) ([]IPreferencesModel, error) {
	preferences, err := ps.repo.FindAll("created_by = ? OR updated_by = ?", userID, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching preferences by user ID: %w", err)
	}
	return preferences, nil
}

func (ps *PreferencesService) UpsertPreferencesByScope(scope string, config t.JSONB, userID string) (IPreferencesModel, error) {
	// Try to find existing preferences by scope
	existing, err := ps.repo.FindOne("scope = ?", scope)
	if err != nil {
		// If not found, create new
		newPrefs := NewPreferencesModel()
		newPrefs.SetScope(scope)
		newPrefs.SetConfig(config)
		if userID != "" {
			newPrefs.SetCreatedBy(userID)
			newPrefs.SetUpdatedBy(userID)
		}
		return ps.CreatePreferences(newPrefs)
	}

	// If found, update existing
	existing.SetConfig(config)
	if userID != "" {
		existing.SetUpdatedBy(userID)
	}
	return ps.UpdatePreferences(existing)
}

func (ps *PreferencesService) GetContextDBService() *is.DBServiceImpl {
	return ps.repo.GetContextDBService()
}
