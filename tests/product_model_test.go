package tests

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	products "github.com/kubex-ecosystem/gdbase/internal/models/products"
)

func setupProductTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	db.AutoMigrate(&products.Product{}, &products.ProductCategory{})
	return db
}

func TestProductModel_CRUD(t *testing.T) {
	db := setupProductTestDB(t)
	cat := &products.ProductCategory{ID: "cat1", Name: "Categoria Teste"}
	db.Create(cat)
	prod := &products.Product{
		ID:         "prod1",
		Code:       "P001",
		SKU:        "SKU001",
		Name:       "Produto Teste",
		CategoryID: cat.ID,
		Unit:       "un",
		IsActive:   true,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
		Stock:      products.Stock{Available: 10, Reserved: 0, Virtual: 0},
	}
	err := db.Create(prod).Error
	assert.NoError(t, err)

	var found products.Product
	err = db.First(&found, "id = ?", prod.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, prod.ID, found.ID)

	prod.Name = "Produto Alterado"
	err = db.Save(prod).Error
	assert.NoError(t, err)

	err = db.Delete(&products.Product{}, prod.ID).Error
	assert.NoError(t, err)
}
