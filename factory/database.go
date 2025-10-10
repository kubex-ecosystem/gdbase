package factory

import (
	"context"
	"embed"

	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
	it "github.com/kubex-ecosystem/gdbase/internal/types"

	l "github.com/kubex-ecosystem/logz"
)

type DBService = svc.IDBService
type IDBService interface {
	svc.IDBService
}
type Rows interface {
	Next() bool
	Scan(dest ...interface{}) error
	Close() error
	Err() error
}
type DirectDatabase interface {
	Query(context.Context, string, ...interface{}) (any, error)
}

type DBServiceImpl = svc.DBServiceImpl

type DBConfig = svc.IDBConfig
type IDBConfig interface {
	svc.IDBConfig
}
type DBConfigImpl = svc.DBConfig

func NewDatabaseService(ctx context.Context, config *DBConfigImpl, logger l.Logger) (*DBServiceImpl, error) {
	return svc.NewDatabaseService(ctx, config, logger)
}

func SetupDatabaseServices(ctx context.Context, d svc.IDockerService, config *DBConfigImpl) error {
	return svc.SetupDatabaseServices(ctx, d, config)
}

func NewDBConfigWithArgs(ctx context.Context, dbName, dbConfigFilePath string, autoMigrate bool, logger l.Logger, debug bool) *DBConfigImpl {
	return svc.NewDBConfigWithArgs(ctx, dbName, dbConfigFilePath, autoMigrate, logger, debug)
}
func NewDBConfigFromFile(ctx context.Context, dbConfigFilePath string, autoMigrate bool, logger l.Logger, debug bool) (*DBConfigImpl, error) {
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
type Messagery = it.Messagery

type Environment = ci.IEnvironment
type EnvironmentType = it.Environment

func NewEnvironment(configFile string, isConfidential bool, logger l.Logger) (*EnvironmentType, error) {
	return it.NewEnvironmentType(configFile, isConfidential, logger)
}

type MongoDB = it.MongoDB
type Redis = it.Redis
type RabbitMQ = it.RabbitMQ

// type IDBService = dbAbs.IDBService

type IDockerService = svc.IDockerService
type DockerService = svc.DockerService

type DatabaseImpl = it.Database
type RabbitMQConfig = it.RabbitMQ
type PostgresConfig = it.Database
type MySQLConfig = it.Database
type SQLiteConfig = it.Database
type MongoDBConfig = it.MongoDB
type RedisConfig = it.Redis

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

type JWT = it.JWT
type JWTImpl = it.JWT

func NewJWT() *JWTImpl {
	return &it.JWT{}
}

type Reference = it.Reference
type IReference interface {
	ci.IReference
}
type ReferenceImpl = it.Reference

func NewReference(name string) IReference {
	return it.NewReference(name)
}
