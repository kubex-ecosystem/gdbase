package interfaces

import "context"

type IDBConfig interface {
	GetDBName() string
	GetDBType() string
	GetEnvironment() string
	GetPostgresConfig() any
	GetMySQLConfig() any
	GetSQLiteConfig() any
	GetMongoDBConfig() any
	GetRedisConfig() any
	GetRabbitMQConfig() any
	IsAutoMigrate() bool
	IsDebug() bool
	GetLogger() any
	GetConfig(context.Context) map[string]any
}
