// Package providers provides the implementation of the providers service.
package providers

import (
	"errors"
	"fmt"

	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type IProvidersService interface {
	CreateProvider(provider IProvidersModel) (IProvidersModel, error)
	GetProviderByID(id string) (IProvidersModel, error)
	UpdateProvider(provider IProvidersModel) (IProvidersModel, error)
	DeleteProvider(id string) error
	ListProviders() ([]IProvidersModel, error)
	GetProviderByName(providerName string) ([]IProvidersModel, error)
	GetProviderByOrgOrGroup(orgOrGroup string) ([]IProvidersModel, error)
	GetProviderByNameAndOrg(providerName, orgOrGroup string) (IProvidersModel, error)
	GetProvidersByUserID(userID string) ([]IProvidersModel, error)
	UpsertProviderByNameAndOrg(providerName, orgOrGroup string, config t.JSONB, userID string) (IProvidersModel, error)
	GetContextDBService() ci.IDBService
}

type ProvidersService struct {
	repo IProvidersRepo
}

func NewProvidersService(repo IProvidersRepo) IProvidersService {
	return &ProvidersService{repo: repo}
}

func (ps *ProvidersService) CreateProvider(provider IProvidersModel) (IProvidersModel, error) {
	if provider.GetProvider() == "" || provider.GetOrgOrGroup() == "" {
		return nil, errors.New("missing required fields: provider and org_or_group are required")
	}

	if provider.GetConfig() == nil {
		return nil, errors.New("missing required field: config cannot be nil")
	}

	// Validate the model
	if err := provider.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	createdProvider, err := ps.repo.Create(provider)
	if err != nil {
		return nil, fmt.Errorf("error creating provider: %w", err)
	}
	return createdProvider, nil
}

func (ps *ProvidersService) GetProviderByID(id string) (IProvidersModel, error) {
	provider, err := ps.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error fetching provider: %w", err)
	}
	return provider, nil
}

func (ps *ProvidersService) UpdateProvider(provider IProvidersModel) (IProvidersModel, error) {
	// Validate the model before updating
	if err := provider.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	updatedProvider, err := ps.repo.Update(provider)
	if err != nil {
		return nil, fmt.Errorf("error updating provider: %w", err)
	}
	return updatedProvider, nil
}

func (ps *ProvidersService) DeleteProvider(id string) error {
	err := ps.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting provider: %w", err)
	}
	return nil
}

func (ps *ProvidersService) ListProviders() ([]IProvidersModel, error) {
	providers, err := ps.repo.FindAll("")
	if err != nil {
		return nil, fmt.Errorf("error listing providers: %w", err)
	}
	return providers, nil
}

func (ps *ProvidersService) GetProviderByName(providerName string) ([]IProvidersModel, error) {
	providers, err := ps.repo.FindAll("provider = ?", providerName)
	if err != nil {
		return nil, fmt.Errorf("error fetching providers by name: %w", err)
	}
	return providers, nil
}

func (ps *ProvidersService) GetProviderByOrgOrGroup(orgOrGroup string) ([]IProvidersModel, error) {
	providers, err := ps.repo.FindAll("org_or_group = ?", orgOrGroup)
	if err != nil {
		return nil, fmt.Errorf("error fetching providers by org/group: %w", err)
	}
	return providers, nil
}

func (ps *ProvidersService) GetProviderByNameAndOrg(providerName, orgOrGroup string) (IProvidersModel, error) {
	provider, err := ps.repo.FindOne("provider = ? AND org_or_group = ?", providerName, orgOrGroup)
	if err != nil {
		return nil, fmt.Errorf("error fetching provider by name and org: %w", err)
	}
	return provider, nil
}

func (ps *ProvidersService) GetProvidersByUserID(userID string) ([]IProvidersModel, error) {
	providers, err := ps.repo.FindAll("created_by = ? OR updated_by = ?", userID, userID)
	if err != nil {
		return nil, fmt.Errorf("error fetching providers by user ID: %w", err)
	}
	return providers, nil
}

func (ps *ProvidersService) UpsertProviderByNameAndOrg(providerName, orgOrGroup string, config t.JSONB, userID string) (IProvidersModel, error) {
	// Try to find existing provider by name and org
	existing, err := ps.repo.FindOne("provider = ? AND org_or_group = ?", providerName, orgOrGroup)
	if err != nil {
		// If not found, create new
		newProvider := NewProvidersModel()
		newProvider.SetProvider(providerName)
		newProvider.SetOrgOrGroup(orgOrGroup)
		newProvider.SetConfig(config)
		if userID != "" {
			newProvider.SetCreatedBy(userID)
			newProvider.SetUpdatedBy(userID)
		}
		return ps.CreateProvider(newProvider)
	}

	// If found, update existing
	existing.SetConfig(config)
	if userID != "" {
		existing.SetUpdatedBy(userID)
	}
	return ps.UpdateProvider(existing)
}

func (ps *ProvidersService) GetContextDBService() ci.IDBService {
	return ps.repo.GetContextDBService()
}
