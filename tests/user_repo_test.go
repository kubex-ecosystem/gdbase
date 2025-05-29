package tests

import (
	"testing"

	um "github.com/rafa-mori/gdbase/internal/models/users"
	"github.com/stretchr/testify/assert"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type MockUserModel struct {
	ID       string
	Name     string
	Username string
	Email    string
	Phone    string
	Active   bool
}

func (m *MockUserModel) GetID() string       { return m.ID }
func (m *MockUserModel) GetName() string     { return m.Name }
func (m *MockUserModel) GetUsername() string { return m.Username }
func (m *MockUserModel) GetEmail() string    { return m.Email }
func (m *MockUserModel) GetPhone() string    { return m.Phone }
func (m *MockUserModel) GetActive() bool     { return m.Active }

func TestUserRepo_Create(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&um.UserModel{})
	repo := um.NewUserRepo(db)

	mockUser := &um.UserModel{ID: "1", Name: "John Doe", Username: "johndoe", Email: "john@example.com", Phone: "1234567890", Active: true}
	createdUser, err := repo.Create(mockUser)

	assert.NoError(t, err)
	assert.Equal(t, mockUser.ID, createdUser.GetID())
}

func TestUserRepo_FindOne(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&um.UserModel{})
	repo := um.NewUserRepo(db)

	mockUser := &um.UserModel{ID: "1", Name: "John Doe"}
	db.Create(mockUser)

	foundUser, err := repo.FindOne("id = ?", "1")

	assert.NoError(t, err)
	assert.Equal(t, mockUser.ID, foundUser.GetID())
}

func TestUserRepo_FindAll(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&um.UserModel{})
	repo := um.NewUserRepo(db)

	mockUser1 := &um.UserModel{ID: "1", Name: "John Doe", Email: "j@d.com", Username: "johndoe"}
	mockUser2 := &um.UserModel{ID: "2", Name: "Jane Doe", Email: "a@j.com", Username: "janedoe"}
	db.Create(mockUser1)
	db.Create(mockUser2)

	users, err := repo.FindAll("id IN(?)", []string{"1", "2"})

	assert.NoError(t, err)
	assert.Len(t, users, 2)
}

func TestUserRepo_Update(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&um.UserModel{})
	repo := um.NewUserRepo(db)

	mockUser := &um.UserModel{ID: "1", Name: "John Doe"}
	db.Create(mockUser)

	mockUser.Name = "John Updated"
	updatedUser, err := repo.Update(mockUser)

	assert.NoError(t, err)
	assert.Equal(t, "John Updated", updatedUser.GetName())
}

func TestUserRepo_Delete(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&um.UserModel{})
	repo := um.NewUserRepo(db)

	mockUser := &um.UserModel{ID: "1", Name: "John Doe"}
	db.Create(mockUser)

	err := repo.Delete("1")

	assert.NoError(t, err)
	var count int64
	db.Model(&um.UserModel{}).Where("id = ?", "1").Count(&count)
	assert.Equal(t, int64(0), count)
}

func TestUserRepo_List(t *testing.T) {
	db, _ := gorm.Open(sqlite.Open(":memory:"), &gorm.Config{})
	db.AutoMigrate(&um.UserModel{})
	repo := um.NewUserRepo(db)

	mockUser1 := &um.UserModel{ID: "1", Name: "John Doe", Username: "johndoe", Email: "john@example.com", Phone: "1234567890", Active: true}
	mockUser2 := &um.UserModel{ID: "2", Name: "Jane Doe", Username: "janedoe", Email: "jane@example.com", Phone: "0987654321", Active: false}
	db.Create(mockUser1)
	db.Create(mockUser2)

	tableHandler, err := repo.List("id IN ?", []string{"1", "2"})

	assert.NoError(t, err)
	assert.Len(t, tableHandler.Rows, 2)
	assert.Equal(t, "John Doe", tableHandler.Rows[0][2])
}
