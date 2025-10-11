package mcp

import (
	"context"
	"time"

	"github.com/google/uuid"
	m "github.com/kubex-ecosystem/gdbase/internal/models/mcp/providers"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type ProvidersModelType = m.ProvidersModel
type ProvidersModel = m.IProvidersModel
type ProvidersService = m.IProvidersService
type ProvidersRepo = m.IProvidersRepo

func NewProvidersService(providersRepo ProvidersRepo) ProvidersService {
	return m.NewProvidersService(providersRepo)
}

func NewProvidersRepo(ctx context.Context, dbService *svc.DBServiceImpl) ProvidersRepo {
	return m.NewProvidersRepo(ctx, dbService)
}

func NewProvidersModel(
	provider string,
	orgOrGroup string,
	config t.JSONBImpl,
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
