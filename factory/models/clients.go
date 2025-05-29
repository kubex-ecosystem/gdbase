package models

import (
	m "github.com/rafa-mori/gdbase/internal/models/clients"
	"gorm.io/gorm"
)

type ClientModel = m.ClientDetailed
type ClientService = m.IClientService
type ClientRepo = m.IClientRepo

func NewClientService(clientRepo ClientRepo) ClientService {
	return m.NewClientService(clientRepo)
}

func NewClientRepo(db *gorm.DB) ClientRepo {
	return m.NewClientRepo(db)
}
