// Package providers provides the implementation of the providers service.
package providers

import (
	"fmt"
	"time"

	t "github.com/kubex-ecosystem/gdbase/types"
)

type IProvidersModel interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetProvider() string
	SetProvider(provider string)
	GetOrgOrGroup() string
	SetOrgOrGroup(orgOrGroup string)
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

type ProvidersModel struct {
	ID         string  `gorm:"type:uuid;primaryKey" json:"id"`
	Provider   string  `gorm:"type:text;not null" json:"provider" example:"github"`
	OrgOrGroup string  `gorm:"type:text;not null" json:"org_or_group" example:"my-org"`
	Config     t.JSONB `json:"config" binding:"omitempty"`
	CreatedAt  string  `gorm:"type:timestamp;default:now()" json:"created_at,omitempty" example:"2024-01-01T00:00:00Z"`
	UpdatedAt  string  `gorm:"type:timestamp;default:now()" json:"updated_at,omitempty" example:"2024-01-01T00:00:00Z"`
	CreatedBy  string  `gorm:"type:uuid;references:users(id)" json:"created_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174001"`
	UpdatedBy  string  `gorm:"type:uuid;references:users(id)" json:"updated_by,omitempty" example:"123e4567-e89b-12d3-a456-426614174002"`
}

func NewProvidersModel() *ProvidersModel {
	return &ProvidersModel{
		ID:         "",
		Provider:   "",
		OrgOrGroup: "",
		Config:     t.JSONB{},
		CreatedAt:  time.Now().Format(time.RFC3339),
		UpdatedAt:  time.Now().Format(time.RFC3339),
	}
}

func (p *ProvidersModel) TableName() string               { return "mcp_provider_configs" }
func (p *ProvidersModel) GetID() string                   { return p.ID }
func (p *ProvidersModel) SetID(id string)                 { p.ID = id }
func (p *ProvidersModel) GetProvider() string             { return p.Provider }
func (p *ProvidersModel) SetProvider(provider string)     { p.Provider = provider }
func (p *ProvidersModel) GetOrgOrGroup() string           { return p.OrgOrGroup }
func (p *ProvidersModel) SetOrgOrGroup(orgOrGroup string) { p.OrgOrGroup = orgOrGroup }
func (p *ProvidersModel) GetConfig() t.JSONB              { return p.Config }
func (p *ProvidersModel) SetConfig(config t.JSONB)        { p.Config = config }
func (p *ProvidersModel) GetCreatedAt() time.Time {
	createdAt, _ := time.Parse(time.RFC3339, p.CreatedAt)
	return createdAt
}
func (p *ProvidersModel) SetCreatedAt(createdAt time.Time) {
	p.CreatedAt = createdAt.Format(time.RFC3339)
}
func (p *ProvidersModel) GetUpdatedAt() time.Time {
	updatedAt, _ := time.Parse(time.RFC3339, p.UpdatedAt)
	return updatedAt
}
func (p *ProvidersModel) SetUpdatedAt(updatedAt time.Time) {
	p.UpdatedAt = updatedAt.Format(time.RFC3339)
}
func (p *ProvidersModel) GetCreatedBy() string          { return p.CreatedBy }
func (p *ProvidersModel) SetCreatedBy(createdBy string) { p.CreatedBy = createdBy }
func (p *ProvidersModel) GetUpdatedBy() string          { return p.UpdatedBy }
func (p *ProvidersModel) SetUpdatedBy(updatedBy string) { p.UpdatedBy = updatedBy }

func (p *ProvidersModel) Validate() error {
	if p.Provider == "" {
		return fmt.Errorf("provider cannot be empty")
	}
	if p.OrgOrGroup == "" {
		return fmt.Errorf("org_or_group cannot be empty")
	}
	if p.Config == nil {
		return fmt.Errorf("config cannot be nil")
	}
	return nil
}

func (p *ProvidersModel) Sanitize() {
	p.UpdatedAt = time.Now().Format(time.RFC3339)
}
