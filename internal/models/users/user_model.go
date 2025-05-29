package user

import (
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// IUser interface for abstraction and encapsulation
//
//go:generate mockgen -destination=../../mocks/mock_user.go -package=mocks . IUser
type IUser interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetName() string
	SetName(name string)
	GetUsername() string
	SetUsername(username string)
	GetPassword() string
	SetPassword(password string) error
	GetEmail() string
	SetEmail(email string)
	GetRoleID() string
	SetRoleID(roleID string)
	GetPhone() string
	SetPhone(phone string)
	GetActive() bool
	SetActive(active bool)
	CheckPasswordHash(password string) bool
	Sanitize()
	Validate() error
	GetDocument() string
	SetDocument(document string)
	GetUserObj() *UserModel
	GetLastLogin() time.Time
	SetLastLogin(lastLogin time.Time)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
}

type UserModel struct {
	ID        string    `gorm:"type:uuid;primaryKey" json:"id"`
	Name      string    `gorm:"type:varchar(255);not null" json:"name"`
	Username  string    `gorm:"type:varchar(255);unique;not null" json:"username"`
	Password  string    `gorm:"type:varchar(255);not null" json:"password"`
	Email     string    `gorm:"type:varchar(255);unique;not null" json:"email"`
	Phone     string    `gorm:"type:varchar(20)" json:"phone"`
	RoleID    string    `gorm:"type:uuid;not null" json:"role_id"`
	Document  string    `gorm:"type:varchar(255)" json:"document"`
	Active    bool      `gorm:"type:boolean;default:true" json:"active"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updated_at"`
	LastLogin time.Time `gorm:"type:timestamp" json:"last_login"`
}

func (um *UserModel) TableName() string {
	return "users"
}
func (um *UserModel) SetID(id string) {
	um.ID = id
}
func (um *UserModel) SetName(name string) {
	um.Name = name
}
func (um *UserModel) SetUsername(username string) {
	um.Username = username
}
func (um *UserModel) SetPassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	um.Password = string(bytes)
	return nil
}
func (um *UserModel) SetEmail(email string) {
	um.Email = email
}
func (um *UserModel) SetRoleID(roleID string) {
	um.RoleID = roleID
}
func (um *UserModel) SetPhone(phone string) {
	um.Phone = phone
}
func (um *UserModel) SetActive(active bool) {
	um.Active = active
}
func (um *UserModel) GetID() string {
	return um.ID
}
func (um *UserModel) GetName() string {
	return um.Name
}
func (um *UserModel) GetUsername() string {
	return um.Username
}
func (um *UserModel) GetPassword() string {
	return um.Password
}
func (um *UserModel) GetEmail() string {
	return um.Email
}
func (um *UserModel) GetRoleID() string {
	return um.RoleID
}
func (um *UserModel) GetPhone() string {
	return um.Phone
}
func (um *UserModel) GetActive() bool {
	return um.Active
}
func (um *UserModel) CheckPasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(um.Password), []byte(password))
	if err != nil {
		return false
	}
	return true
}
func (um *UserModel) Sanitize() {
	um.Password = ""
	um.Active = false
}
func (um *UserModel) Validate() error {
	if um.Username == "" {
		return fmt.Errorf("username is required")
	}
	if um.Email == "" {
		return fmt.Errorf("email is required")
	}
	if um.Password == "" {
		return fmt.Errorf("password is required")
	}
	if um.Name == "" {
		return fmt.Errorf("name is required")
	}
	if um.Phone == "" {
		return fmt.Errorf("phone is required")
	}
	return nil
}
func (um *UserModel) GetUserObj() *UserModel {
	return um
}
func (um *UserModel) SetDocument(document string) {
	// Implement the logic to set the document
}
func (um *UserModel) GetDocument() string {
	// Implement the logic to get the document
	return ""
}
func (um *UserModel) GetLastLogin() time.Time {
	return um.LastLogin
}
func (um *UserModel) SetLastLogin(lastLogin time.Time) {
	um.LastLogin = lastLogin
}
func (um *UserModel) GetCreatedAt() time.Time {
	return um.CreatedAt
}
func (um *UserModel) SetCreatedAt(createdAt time.Time) {
	um.CreatedAt = createdAt
}
func (um *UserModel) GetUpdatedAt() time.Time {
	return um.UpdatedAt
}
func (um *UserModel) SetUpdatedAt(updatedAt time.Time) {
	um.UpdatedAt = updatedAt
}
