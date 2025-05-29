package types

import (
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
	GetConfig() *DBConfig
}
