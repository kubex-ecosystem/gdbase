package types

import (
	it "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
	t "github.com/kubex-ecosystem/gdbase/internal/types"
	s "github.com/kubex-ecosystem/gdbase/services"
	l "github.com/kubex-ecosystem/logz"
)

const (
	// DefaultConfigDir is the default directory for configuration files
	DefaultConfigDir = svc.DefaultConfigDir

	// DefaultConfigFile is the default configuration file path
	DefaultConfigFile     = svc.DefaultConfigFile
	DefaultVolumesDir     = svc.DefaultVolumesDir
	DefaultRedisVolume    = svc.DefaultRedisVolume
	DefaultPostgresVolume = svc.DefaultPostgresVolume
	DefaultMongoVolume    = svc.DefaultMongoVolume
	DefaultRabbitMQVolume = svc.DefaultRabbitMQVolume
)

type DBConfig = svc.DBConfig
type IDBService = s.IDBService
type DBService = svc.DBService
type DatabaseType = t.Database
type EnvironmentType = it.IEnvironment
type RabbitMQConfig = t.RabbitMQ
type PostgresConfig = t.Database
type MySQLConfig = t.Database
type SQLiteConfig = t.Database
type MongoDBConfig = t.Database
type RedisConfig = t.Database

// NewDBConfig creates a new DBConfig instance

func NewDBConfig(name, filePath string, enabled bool, logger l.Logger, debug bool) *DBConfig {
	return nil
}
