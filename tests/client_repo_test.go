package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	clients "github.com/rafa-mori/gdbase/internal/models/clients"
)

func setupClientRepoTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	db.AutoMigrate(&clients.ClientDetailed{})
	return db
}

func TestClientRepo_CRUD(t *testing.T) {
	db := setupClientRepoTestDB(t)
	repo := clients.NewClientRepo(db)
	client := &clients.ClientDetailed{
		ID:           "client2",
		Code:         ptrStr("C002"),
		TradingName:  ptrStr("Empresa Teste 2"),
		DocumentType: clients.Company,
		Status:       clients.Active,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}
	// Create
	created, err := repo.Create(client)
	assert.NoError(t, err)
	assert.Equal(t, client.ID, created.ID)

	// FindOne
	found, err := repo.FindOne("id = ?", client.ID)
	assert.NoError(t, err)
	assert.Equal(t, client.ID, found.ID)

	// Update
	client.Status = clients.Blocked
	updated, err := repo.Update(client)
	assert.NoError(t, err)
	assert.Equal(t, clients.Blocked, updated.Status)

	// FindAll
	list, err := repo.FindAll("status = ?", clients.Blocked)
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	// Delete
	err = repo.Delete(client.ID)
	assert.NoError(t, err)
}

func ptrStr(s string) *string { return &s }
