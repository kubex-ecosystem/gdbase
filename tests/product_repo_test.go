package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	products "github.com/kubex-ecosystem/gdbase/internal/models/products"
)

func setupProductRepoTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	db.AutoMigrate(&products.Product{}, &products.ProductCategory{})
	return db
}

func TestProductRepo_CRUD(t *testing.T) {
	db := setupProductRepoTestDB(t)
	repo := products.NewProductRepo(db)
	cat := &products.ProductCategory{ID: "cat1", Name: "Categoria Teste"}
	db.Create(cat)
	prod := &products.Product{
		ID:         "prod2",
		Code:       "P002",
		SKU:        "SKU002",
		Name:       "Produto Teste 2",
		CategoryID: cat.ID,
		Unit:       "un",
		IsActive:   true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Stock:      products.Stock{Available: 5, Reserved: 0, Virtual: 0},
	}
	// Create
	created, err := repo.Create(prod)
	assert.NoError(t, err)
	assert.Equal(t, prod.ID, created.ID)

	// FindOne
	found, err := repo.FindOne("id = ?", prod.ID)
	assert.NoError(t, err)
	assert.Equal(t, prod.ID, found.ID)

	// Update
	prod.Name = "Produto Alterado 2"
	updated, err := repo.Update(prod)
	assert.NoError(t, err)
	assert.Equal(t, "Produto Alterado 2", updated.Name)

	// FindAll
	list, err := repo.FindAll("category_id = ?", cat.ID)
	assert.NoError(t, err)
	assert.Len(t, list, 1)

	// Delete
	err = repo.Delete(prod.ID)
	assert.NoError(t, err)
}
