// Package llm provides the model definitions for LLM (Large Language Model) operations.
package llm

import (
	"fmt"
	"time"
)

type ILLMModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetProvider() string
	SetProvider(provider string)
	GetModel() string
	SetModel(model string)
	GetTemperature() float64
	SetTemperature(temperature float64)
	GetMaxTokens() int
	SetMaxTokens(maxTokens int)
	GetTopP() float64
	SetTopP(topP float64)
	GetFrequencyPenalty() float64
	SetFrequencyPenalty(frequencyPenalty float64)
	GetPresencePenalty() float64
	SetPresencePenalty(presencePenalty float64)
	GetStopSequences() []string
	SetStopSequences(stopSequences []string)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	GetCreatedBy() string
	SetCreatedBy(createdBy string)
	GetUpdatedBy() string
	SetUpdatedBy(updatedBy string)
	Validate() error
	Sanitize()
}

type LLMModel struct {
	ID               string   `gorm:"type:uuid;primaryKey" json:"id"`
	Enabled          bool     `gorm:"type:boolean;default:true" json:"enabled" example:"true"`
	Provider         string   `gorm:"type:text" json:"provider" example:"openai"`
	Model            string   `gorm:"type:text" json:"model" example:"gpt-4"`
	Temperature      float64  `gorm:"type:real" json:"temperature" example:"0.7"`
	MaxTokens        int      `gorm:"type:integer" json:"max_tokens" example:"2048"`
	TopP             float64  `gorm:"type:real" json:"top_p" example:"1.0"`
	FrequencyPenalty float64  `gorm:"type:real" json:"frequency_penalty" example:"0.0"`
	PresencePenalty  float64  `gorm:"type:real" json:"presence_penalty" example:"0.0"`
	StopSequences    []string `gorm:"type:text[]" json:"stop_sequences,omitempty" example:"[\"\\n\", \"\\\"\\\"]"`
	CreatedAt        string   `gorm:"type:timestamp" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt        string   `gorm:"type:timestamp" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy        string   `gorm:"type:uuid" json:"created_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	UpdatedBy        string   `gorm:"type:uuid" json:"updated_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
}

func NewLLMModel() *LLMModel {
	return &LLMModel{
		ID:               "",
		Enabled:          true,
		Provider:         "",
		Model:            "",
		Temperature:      0.7,
		MaxTokens:        2048,
		TopP:             1.0,
		FrequencyPenalty: 0.0,
		PresencePenalty:  0.0,
		StopSequences:    []string{},
		CreatedAt:        time.Now().Format(time.RFC3339),
		UpdatedAt:        time.Now().Format(time.RFC3339),
	}
}

func (m *LLMModel) TableName() string                  { return "mcp_llm_models" }
func (m *LLMModel) GetID() string                      { return m.ID }
func (m *LLMModel) SetID(id string)                    { m.ID = id }
func (m *LLMModel) GetProvider() string                { return m.Provider }
func (m *LLMModel) SetProvider(provider string)        { m.Provider = provider }
func (m *LLMModel) GetModel() string                   { return m.Model }
func (m *LLMModel) SetModel(model string)              { m.Model = model }
func (m *LLMModel) GetTemperature() float64            { return m.Temperature }
func (m *LLMModel) SetTemperature(temperature float64) { m.Temperature = temperature }
func (m *LLMModel) GetMaxTokens() int                  { return m.MaxTokens }
func (m *LLMModel) SetMaxTokens(maxTokens int)         { m.MaxTokens = maxTokens }
func (m *LLMModel) GetTopP() float64                   { return m.TopP }
func (m *LLMModel) SetTopP(topP float64)               { m.TopP = topP }
func (m *LLMModel) GetFrequencyPenalty() float64       { return m.FrequencyPenalty }
func (m *LLMModel) SetFrequencyPenalty(frequencyPenalty float64) {
	m.FrequencyPenalty = frequencyPenalty
}
func (m *LLMModel) GetPresencePenalty() float64                { return m.PresencePenalty }
func (m *LLMModel) SetPresencePenalty(presencePenalty float64) { m.PresencePenalty = presencePenalty }
func (m *LLMModel) GetStopSequences() []string                 { return m.StopSequences }
func (m *LLMModel) SetStopSequences(stopSequences []string)    { m.StopSequences = stopSequences }
func (m *LLMModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, m.CreatedAt)
	return createdAt
}
func (m *LLMModel) SetCreatedAt(createdAt time.Time) { m.CreatedAt = createdAt.Format(time.RFC3339) }
func (m *LLMModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, m.UpdatedAt)
	return updatedAt
}
func (m *LLMModel) SetUpdatedAt(updatedAt time.Time) { m.UpdatedAt = updatedAt.Format(time.RFC3339) }
func (m *LLMModel) GetCreatedBy() string             { return m.CreatedBy }
func (m *LLMModel) SetCreatedBy(createdBy string)    { m.CreatedBy = createdBy }
func (m *LLMModel) GetUpdatedBy() string             { return m.UpdatedBy }
func (m *LLMModel) SetUpdatedBy(updatedBy string)    { m.UpdatedBy = updatedBy }
func (m *LLMModel) Validate() error {
	if m.Provider == "" {
		return fmt.Errorf("provider cannot be empty")
	}
	if m.Model == "" {
		return fmt.Errorf("model cannot be empty")
	}
	if m.Temperature < 0 || m.Temperature > 1 {
		return fmt.Errorf("temperature must be between 0 and 1")
	}
	if m.MaxTokens <= 0 {
		return fmt.Errorf("max_tokens must be greater than 0")
	}
	return nil
}
func (m *LLMModel) Sanitize() {
	// m.ID = ""
	// m.CreatedAt = time.Now().Format(time.RFC3339)
	m.UpdatedAt = time.Now().Format(time.RFC3339)
	// m.CreatedBy = ""
	// m.UpdatedBy = ""
}
