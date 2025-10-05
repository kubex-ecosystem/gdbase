package mcp

import (
	"context"
	"time"

	"github.com/google/uuid"
	m "github.com/kubex-ecosystem/gdbase/internal/models/mcp/preferences"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
)

type PreferencesModelType = m.PreferencesModel
type PreferencesModel = m.IPreferencesModel
type PreferencesService = m.IPreferencesService
type PreferencesRepo = m.IPreferencesRepo

func NewPreferencesService(preferencesRepo PreferencesRepo) PreferencesService {
	return m.NewPreferencesService(preferencesRepo)
}

func NewPreferencesRepo(ctx context.Context, dbService *svc.DBServiceImpl) PreferencesRepo {
	return m.NewPreferencesRepo(ctx, dbService)
}

func NewPreferencesModel(
	scope string,
	config t.JSONB,
) PreferencesModel {
	return &m.PreferencesModel{
		ID:        uuid.New().String(),
		Scope:     scope,
		Config:    config,
		CreatedAt: time.Now().Format("2006-01-02 15:04:05"),
		UpdatedAt: time.Now().Format("2006-01-02 15:04:05"),
		CreatedBy: "admin", // Temporary, should be set by the service layer
	}
}
