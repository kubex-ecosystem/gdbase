// Package mcp provides the LLM model and its methods for text generation and manipulation
package mcp

import (
	"time"

	"github.com/google/uuid"
	m "github.com/kubex-ecosystem/gdbase/internal/models/mcp/llm"
	"gorm.io/gorm"
)

type LLMModelType = m.LLMModel
type LLMModel = m.ILLMModel
type LLMService = m.ILLMService
type LLMRepo = m.ILLMRepo

func NewLLMService(llmRepo LLMRepo) LLMService {
	return m.NewLLMService(llmRepo)
}

func NewLLMRepo(db *gorm.DB) LLMRepo {
	return m.NewLLMRepo(db)
}

func NewLLMModel(
	enabled bool,
	provider string,
	model string,
	temperature float64,
	maxTokens int,
	topP float64,
	frequencyPenalty float64,
	presencePenalty float64,
	stopSequences []string,
) LLMModel {
	return &m.LLMModel{
		ID:               uuid.New().String(),
		Enabled:          enabled,
		Provider:         provider,
		Model:            model,
		Temperature:      temperature,
		MaxTokens:        maxTokens,
		TopP:             topP,
		FrequencyPenalty: frequencyPenalty,
		PresencePenalty:  presencePenalty,
		StopSequences:    stopSequences,
		CreatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:        time.Now().Format("2006-01-02 15:04:05"),
		CreatedBy:        "admin", // Temporary, should be set by the service layer
	}
}
