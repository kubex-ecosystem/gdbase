package mcp

import (
	"time"

	"github.com/google/uuid"
	m "github.com/rafa-mori/gdbase/internal/models/mcp/providers"
	t "github.com/rafa-mori/gdbase/types"
	"gorm.io/gorm"
)

type ProvidersModelType = m.ProvidersModel
type ProvidersModel = m.IProvidersModel
type ProvidersService = m.IProvidersService
type ProvidersRepo = m.IProvidersRepo

func NewProvidersService(providersRepo ProvidersRepo) ProvidersService {
	return m.NewProvidersService(providersRepo)
}

func NewProvidersRepo(db *gorm.DB) ProvidersRepo {
	return m.NewProvidersRepo(db)
}

func NewProvidersModel(
	provider string,
	orgOrGroup string,
	config t.JSONB,
) ProvidersModel {
	return &m.ProvidersModel{
		ID:         uuid.New().String(),
		Provider:   provider,
		OrgOrGroup: orgOrGroup,
		Config:     config,
		CreatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt:  time.Now().Format("2006-01-02 15:04:05"),
		CreatedBy:  "admin", // Temporary, should be set by the service layer
	}
}
