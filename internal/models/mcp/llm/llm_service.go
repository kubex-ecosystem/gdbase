package llm

import (
	"errors"
	"fmt"

	t "github.com/rafa-mori/gdbase/types"
)

type ILLMService interface {
	CreateLLMModel(model ILLMModel) (ILLMModel, error)
	GetLLMModelByID(id string) (ILLMModel, error)
	UpdateLLMModel(model ILLMModel) (ILLMModel, error)
	DeleteLLMModel(id string) error
	ListLLMModels() ([]ILLMModel, error)
	GetLLMModelByProvider(provider string) ([]ILLMModel, error)
	GetLLMModelByProviderAndModel(provider, model string) (ILLMModel, error)
	GetEnabledLLMModels() ([]ILLMModel, error)
	GetContextDBService() t.IDBService
}

type LLMService struct {
	repo ILLMRepo
}

func NewLLMService(repo ILLMRepo) ILLMService {
	return &LLMService{repo: repo}
}

func (ls *LLMService) CreateLLMModel(model ILLMModel) (ILLMModel, error) {
	if model.GetProvider() == "" || model.GetModel() == "" {
		return nil, errors.New("missing required fields: provider and model are required")
	}

	// Validate the model
	if err := model.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	createdModel, err := ls.repo.Create(model)
	if err != nil {
		return nil, fmt.Errorf("error creating LLM model: %w", err)
	}
	return createdModel, nil
}

func (ls *LLMService) GetLLMModelByID(id string) (ILLMModel, error) {
	model, err := ls.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error fetching LLM model: %w", err)
	}
	return model, nil
}

func (ls *LLMService) UpdateLLMModel(model ILLMModel) (ILLMModel, error) {
	// Validate the model before updating
	if err := model.Validate(); err != nil {
		return nil, fmt.Errorf("validation error: %w", err)
	}

	updatedModel, err := ls.repo.Update(model)
	if err != nil {
		return nil, fmt.Errorf("error updating LLM model: %w", err)
	}
	return updatedModel, nil
}

func (ls *LLMService) DeleteLLMModel(id string) error {
	err := ls.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting LLM model: %w", err)
	}
	return nil
}

func (ls *LLMService) ListLLMModels() ([]ILLMModel, error) {
	models, err := ls.repo.FindAll("")
	if err != nil {
		return nil, fmt.Errorf("error listing LLM models: %w", err)
	}
	return models, nil
}

func (ls *LLMService) GetLLMModelByProvider(provider string) ([]ILLMModel, error) {
	models, err := ls.repo.FindAll("provider = ?", provider)
	if err != nil {
		return nil, fmt.Errorf("error fetching LLM models by provider: %w", err)
	}
	return models, nil
}

func (ls *LLMService) GetLLMModelByProviderAndModel(provider, model string) (ILLMModel, error) {
	llmModel, err := ls.repo.FindOne("provider = ? AND model = ?", provider, model)
	if err != nil {
		return nil, fmt.Errorf("error fetching LLM model by provider and model: %w", err)
	}
	return llmModel, nil
}

func (ls *LLMService) GetEnabledLLMModels() ([]ILLMModel, error) {
	models, err := ls.repo.FindAll("enabled = ?", true)
	if err != nil {
		return nil, fmt.Errorf("error listing enabled LLM models: %w", err)
	}
	return models, nil
}

func (ls *LLMService) GetContextDBService() t.IDBService {
	return ls.repo.GetContextDBService()
}
