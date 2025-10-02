package factory

import (
	"context"
	"embed"

	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
	it "github.com/kubex-ecosystem/gdbase/internal/types"

	l "github.com/kubex-ecosystem/logz"
)

type DBService = ci.IDBService
type IDBService interface {
	ci.IDBService
}
type DBConfig = ci.IDBConfig
type IDBConfig interface {
	ci.IDBConfig
}

func NewDatabaseService(ctx context.Context, config DBConfig, logger l.Logger) (DBService, error) {
	return svc.NewDatabaseService(ctx, config.(*svc.DBConfig), logger)
}

func SetupDatabaseServices(ctx context.Context, d svc.IDockerService, config DBConfig) error {
	return svc.SetupDatabaseServices(ctx, d, config.(*svc.DBConfig))
}

func NewDBConfigWithArgs(ctx context.Context, dbName, dbConfigFilePath string, autoMigrate bool, logger l.Logger, debug bool) DBConfig {
	return svc.NewDBConfigWithArgs(ctx, dbName, dbConfigFilePath, autoMigrate, logger, debug)
}
func NewDBConfigFromFile(ctx context.Context, dbConfigFilePath string, autoMigrate bool, logger l.Logger, debug bool) (DBConfig, error) {
	return svc.NewDBConfigFromFile(ctx, dbConfigFilePath, autoMigrate, logger, debug)
}

var migrationFiles embed.FS

func SetMigrationFiles(mf embed.FS) {
	migrationFiles = mf
}

func GetMigrationFiles() embed.FS {
	return migrationFiles
}

type Database = it.Database
type Environment = it.Environment
type MongoDB = it.MongoDB
type Redis = it.Redis
type RabbitMQ = it.RabbitMQ

// type IDBService = dbAbs.IDBService

type IDockerService = svc.IDockerService
type DockerService = svc.DockerService

type DatabaseType = it.Database
type RabbitMQConfig = it.RabbitMQ
type PostgresConfig = it.Database
type MySQLConfig = it.Database
type SQLiteConfig = it.Database
type MongoDBConfig = it.MongoDB
type RedisConfig = it.Redis
type EnvironmentType = it.Environment

type JSONB = ci.IJSONB
type IJSONB interface {
	ci.IJSONB
}
type JSONBImpl = it.JSONB

func NewJSONB() IJSONB {
	return &it.JSONB{}
}

type JSONBData = it.JSONBData
type IJSONBData interface{ ci.IJSONB }

func NewJSONBData() IJSONBData { return it.NewJSONBData() }
