package tests

// import (
// 	"testing"

// 	um "github.com/rafa-mori/gdbase/internal/models/users"
// 	"github.com/stretchr/testify/assert"
// 	"golang.org/x/crypto/bcrypt"
// )

// func TestUserModel_SetName(t *testing.T) {
// 	user := &um.UserModel{}
// 	user.SetName("John Doe")
// 	assert.Equal(t, "John Doe", user.Name)
// }

// func TestUserModel_SetUsername(t *testing.T) {
// 	user := &um.UserModel{}
// 	user.SetUsername("johndoe")
// 	assert.Equal(t, "johndoe", user.Username)
// }

// func TestUserModel_SetPassword(t *testing.T) {
// 	user := &um.UserModel{}
// 	err := user.SetPassword("securepassword")
// 	assert.NoError(t, err)

// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte("securepassword"))
// 	assert.NoError(t, err)
// }

// func TestUserModel_SetEmail(t *testing.T) {
// 	user := &um.UserModel{}
// 	user.SetEmail("johndoe@example.com")
// 	assert.Equal(t, "johndoe@example.com", user.Email)
// }

// func TestUserModel_SetRoleID(t *testing.T) {
// 	user := &um.UserModel{}
// 	user.SetRoleID(1)
// 	assert.Equal(t, uint(1), user.RoleID)
// }

// func TestUserModel_SetPhone(t *testing.T) {
// 	user := &um.UserModel{}
// 	user.SetPhone("1234567890")
// 	assert.Equal(t, "1234567890", user.Phone)
// }

// func TestUserModel_SetActive(t *testing.T) {
// 	user := &um.UserModel{}
// 	user.SetActive(true)
// 	assert.True(t, user.Active)
// }

// func TestUserModel_Getters(t *testing.T) {
// 	user := &um.UserModel{
// 		ID:       "123",
// 		Name:     "John Doe",
// 		Username: "johndoe",
// 		Password: "hashedpassword",
// 		Email:    "johndoe@example.com",
// 		Phone:    "1234567890",
// 		RoleID:   1,
// 		Active:   true,
// 	}

// 	assert.Equal(t, "123", user.GetID())
// 	assert.Equal(t, "John Doe", user.GetName())
// 	assert.Equal(t, "johndoe", user.GetUsername())
// 	assert.Equal(t, "hashedpassword", user.GetPassword())
// 	assert.Equal(t, "johndoe@example.com", user.GetEmail())
// 	assert.Equal(t, uint(1), user.GetRoleID())
// 	assert.Equal(t, "1234567890", user.GetPhone())
// 	assert.True(t, user.GetActive())
// }
