// Package services provides an interface for database service operations.
package services

import (
	t "github.com/rafa-mori/gdbase/types"
	"gorm.io/gorm"
)

type IDBService interface {
	Initialize() error
	GetDB() (*gorm.DB, error)
	CloseDBConnection() error
	CheckDatabaseHealth() error
	IsConnected() error
	Reconnect() error
	GetHost() (string, error)
	GetConfig() *t.DBConfig
}
