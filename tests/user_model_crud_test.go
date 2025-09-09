package tests

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	um "github.com/kubex-ecosystem/gdbase/internal/models/users"
)

func setupUserTestDB(t *testing.T) *gorm.DB {
	db, err := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	if err != nil {
		t.Fatalf("failed to connect to test db: %v", err)
	}
	db.AutoMigrate(&um.UserModel{})
	return db
}

func TestUserModel_CRUD(t *testing.T) {
	db := setupUserTestDB(t)
	user := &um.UserModel{
		ID:       "user1",
		Name:     "Usuário Teste",
		Username: "usertest",
		Email:    "user@test.com",
		Phone:    "123456789",
		Active:   true,
	}
	err := db.Create(user).Error
	assert.NoError(t, err)

	var found um.UserModel
	err = db.First(&found, "id = ?", user.ID).Error
	assert.NoError(t, err)
	assert.Equal(t, user.ID, found.ID)

	user.Name = "Usuário Alterado"
	err = db.Save(user).Error
	assert.NoError(t, err)

	err = db.Delete(&um.UserModel{}, user.ID).Error
	assert.NoError(t, err)
}
