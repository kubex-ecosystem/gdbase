package factory

import (
	l "github.com/kubex-ecosystem/logz"
	dbAbs "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/types"
)

type DBConfig = t.DBConfig

type DBService = t.IDBService
type IDBConfig = t.DBConfig

func NewDatabaseService(config *t.DBConfig, logger l.Logger) (DBService, error) {
	return dbAbs.NewDatabaseService(config, logger)
}

func SetupDatabaseServices(d dbAbs.IDockerService, config *t.DBConfig) error {
	return dbAbs.SetupDatabaseServices(d, config)
}
