// Package interfaces defines the interfaces for database configuration and services.
package interfaces

import (
	"context"
	"database/sql"
	"time"

	"gorm.io/gorm"
)

type IDBService interface {
	Initialize(ctx context.Context) error
	InitializeFromEnv(ctx context.Context, env IEnvironment) error
	GetDB(ctx context.Context) (*gorm.DB, error)
	CloseDBConnection(ctx context.Context) error
	CheckDatabaseHealth(ctx context.Context) error
	GetConnection(ctx context.Context, database string, timeout time.Duration) (*sql.Conn, error)
	IsConnected(ctx context.Context) error
	Reconnect(ctx context.Context) error
	GetName(ctx context.Context) (string, error)
	GetHost(ctx context.Context) (string, error)
	GetConfig(ctx context.Context) map[string]any
	RunMigrations(ctx context.Context, files map[string]string) (int, int, error)
}
