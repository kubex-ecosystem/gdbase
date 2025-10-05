// Package services provides an interface for database service operations.
package services

import (
	"context"

	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
	l "github.com/kubex-ecosystem/logz"
)

type DBConfig = svc.DBConfig
type IDBConfig interface {
	svc.DBConfig
}

type IDBService interface {
	svc.IDBService
}

type DBService = svc.DBServiceImpl

type IDockerService = svc.IDockerService
type DockerService = svc.DockerService

func NewDatabaseService(ctx context.Context, config *svc.DBConfig, logger l.Logger) (svc.IDBService, error) {
	return svc.NewDatabaseService(ctx, config, logger)
}

func SetupDatabaseServices(ctx context.Context, d svc.IDockerService, config *svc.DBConfig) error {
	return svc.SetupDatabaseServices(ctx, d, config)
}

func NewDBConfigWithArgs(ctx context.Context, dbName, dbConfigFilePath string, autoMigrate bool, logger l.Logger, debug bool) *svc.DBConfig {
	return svc.NewDBConfigWithArgs(ctx, dbName, dbConfigFilePath, autoMigrate, logger, debug)
}
func NewDBConfigFromFile(ctx context.Context, dbConfigFilePath string, autoMigrate bool, logger l.Logger, debug bool) (*svc.DBConfig, error) {
	return svc.NewDBConfigFromFile(ctx, dbConfigFilePath, autoMigrate, logger, debug)
}

type DatabaseType = t.Database
type IDatabase interface {
	t.Database
}
type DatabaseObj = *t.Database

type RabbitMQ = t.RabbitMQ
type IRabbitMQ interface {
	t.RabbitMQ
}
type RabbitMQObj = *t.RabbitMQ

type Postgres = t.Database
type IPostgres interface {
	t.Database
}
type PostgresObj = *t.Database

type MySQL = t.Database
type IMySQL interface {
	t.Database
}
type MySQLObj = *t.Database

type SQLite = t.Database
type ISQLite interface {
	t.Database
}
type SQLiteObj = *t.Database

type MongoDB = t.MongoDB
type IMongoDB interface {
	t.MongoDB
}
type MongoDBObject = *t.MongoDB

type Redis = t.Redis
type IRedis interface {
	t.Redis
}
type RedisObj = *t.Redis

type EnvironmentType = t.Environment

type IJSONB interface {
	ci.IJSONB
}
type JSONBObj = ci.IJSONB
type JSONBImpl = t.JSONB
