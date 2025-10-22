// Package registrationtokens provides the model for registration tokens
package registrationtokens

import (
	"time"

	"github.com/google/uuid"
)

// IRegistrationToken defines the interface for a registration token.
type IRegistrationToken interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetUserID() string
	SetUserID(userID string)
	GetToken() string
	SetToken(token string)
	GetExpiresAt() time.Time
	SetExpiresAt(expiresAt time.Time)
	IsExpired() bool
	GetTokenObj() *RegistrationTokenModel
}

// RegistrationTokenModel represents the database model for a registration token.
type RegistrationTokenModel struct {
	ID        string    `gorm:"type:uuid;primaryKey"`
	UserID    string    `gorm:"type:uuid;not null"`
	Token     string    `gorm:"type:varchar(255);unique;not null"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null"`
}

// NewRegistrationToken creates a new instance of a registration token.
func NewRegistrationToken(userID, token string, expiresAt time.Time) *RegistrationTokenModel {
	return &RegistrationTokenModel{
		ID:        uuid.New().String(),
		UserID:    userID,
		Token:     token,
		ExpiresAt: expiresAt,
	}
}

func (rtm *RegistrationTokenModel) TableName() string {
	return "registration_tokens"
}

func (rtm *RegistrationTokenModel) GetID() string {
	return rtm.ID
}

func (rtm *RegistrationTokenModel) SetID(id string) {
	rtm.ID = id
}

func (rtm *RegistrationTokenModel) GetUserID() string {
	return rtm.UserID
}

func (rtm *RegistrationTokenModel) SetUserID(userID string) {
	rtm.UserID = userID
}

func (rtm *RegistrationTokenModel) GetToken() string {
	return rtm.Token
}

func (rtm *RegistrationTokenModel) SetToken(token string) {
	rtm.Token = token
}

func (rtm *RegistrationTokenModel) GetExpiresAt() time.Time {
	return rtm.ExpiresAt
}

func (rtm *RegistrationTokenModel) SetExpiresAt(expiresAt time.Time) {
	rtm.ExpiresAt = expiresAt
}

func (rtm *RegistrationTokenModel) IsExpired() bool {
	return time.Now().After(rtm.ExpiresAt)
}

func (rtm *RegistrationTokenModel) GetTokenObj() *RegistrationTokenModel {
	return rtm
}
