package models

import (
	"context"

	m "github.com/kubex-ecosystem/gdbase/internal/models/clients"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

type ClientModel = m.ClientDetailed
type ClientService = m.IClientService
type ClientRepo = m.IClientRepo

func NewClientService(clientRepo ClientRepo) ClientService {
	return m.NewClientService(clientRepo)
}

func NewClientRepo(ctx context.Context, dbService *svc.DBServiceImpl) ClientRepo {
	return m.NewClientRepo(ctx, dbService)
}
