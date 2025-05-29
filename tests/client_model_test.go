package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	clients "github.com/rafa-mori/gdbase/internal/models/clients"
)

func setupClientTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	db.AutoMigrate(&clients.ClientDetailed{})
	return db
}

func TestClientModel_CRUD(t *testing.T) {
	db := setupClientTestDB(t)
	client := &clients.ClientDetailed{
		ID:           "client1",
		Code:         ptrStr("C001"),
		TradingName:  ptrStr("Empresa Teste"),
		DocumentType: clients.Company,
		Status:       clients.Active,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	err := db.Create(client).Error
	assert.NoError(t, err)

	var found clients.ClientDetailed
	err = db.First(&found, "id = ?", client.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, client.ID, found.ID)

	client.Status = clients.Blocked
	err = db.Save(client).Error
	assert.NoError(t, err)

	err = db.Delete(&clients.ClientDetailed{}, client.ID).Error
	assert.NoError(t, err)
}
