// Package preferences provides the model for user preferences in the application.
package preferences

import (
	"fmt"
	"time"

	t "github.com/kubex-ecosystem/gdbase/types"
)

type IPreferencesModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetScope() string
	SetScope(scope string)
	GetConfig() t.JSONB
	SetConfig(config t.JSONB)
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

type PreferencesModel struct {
	ID        string  `gorm:"type:uuid;primaryKey" json:"id"`
	Scope     string  `gorm:"type:text;not null;default:'defaults';uniqueIndex" json:"scope" example:"defaults"`
	Config    t.JSONB `json:"config" binding:"omitempty"`
	CreatedAt string  `gorm:"type:timestamp;default:now()" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt string  `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy string  `gorm:"type:uuid;references:users(id)" json:"created_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	UpdatedBy string  `gorm:"type:uuid;references:users(id)" json:"updated_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
}

func NewPreferencesModel() *PreferencesModel {
	return &PreferencesModel{
		ID:        "",
		Scope:     "defaults",
		Config:    make(t.JSONB),
		CreatedAt: time.Now().Format(time.RFC3339),
		UpdatedAt: time.Now().Format(time.RFC3339),
	}
}

func (p *PreferencesModel) TableName() string        { return "mcp_user_preferences" }
func (p *PreferencesModel) GetID() string            { return p.ID }
func (p *PreferencesModel) SetID(id string)          { p.ID = id }
func (p *PreferencesModel) GetScope() string         { return p.Scope }
func (p *PreferencesModel) SetScope(scope string)    { p.Scope = scope }
func (p *PreferencesModel) GetConfig() t.JSONB       { return p.Config }
func (p *PreferencesModel) SetConfig(config t.JSONB) { p.Config = config }
func (p *PreferencesModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, p.CreatedAt)
	return createdAt
}
func (p *PreferencesModel) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt.Format(time.RFC3339)
}
func (p *PreferencesModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, p.UpdatedAt)
	return updatedAt
}
func (p *PreferencesModel) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt.Format(time.RFC3339)
}
func (p *PreferencesModel) GetCreatedBy() string          { return p.CreatedBy }
func (p *PreferencesModel) SetCreatedBy(createdBy string) { p.CreatedBy = createdBy }
func (p *PreferencesModel) GetUpdatedBy() string          { return p.UpdatedBy }
func (p *PreferencesModel) SetUpdatedBy(updatedBy string) { p.UpdatedBy = updatedBy }

func (p *PreferencesModel) Validate() error {
	if p.Scope == "" {
		return fmt.Errorf("scope cannot be empty")
	}
	if p.Config.IsNil() {
		p.Config = make(t.JSONB) // Initialize if nil
	}
	return nil
}

func (p *PreferencesModel) Sanitize() {
	p.UpdatedAt = time.Now().Format(time.RFC3339)
}
