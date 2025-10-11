// Package auth provides authentication and token management models
package auth

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

// IRefreshToken defines the interface for refresh token operations
type IRefreshToken interface {
	TableName() string
	GetID() uint
	SetID(id uint)
	GetUserID() string
	SetUserID(userID string)
	GetTokenID() string
	SetTokenID(tokenID string)
	GetExpiresAt() time.Time
	SetExpiresAt(expiresAt time.Time)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	Validate() error
	Sanitize()
	IsExpired() bool
}

// RefreshTokenModel represents a refresh token in the database
type RefreshTokenModel struct {
	ID        uint      `gorm:"primaryKey;autoIncrement" json:"id"`
	UserID    string    `gorm:"type:uuid;index;not null" json:"user_id"`
	TokenID   string    `gorm:"type:uuid;uniqueIndex;not null" json:"token_id"`
	ExpiresAt time.Time `gorm:"type:timestamp;not null" json:"expires_at"`
	CreatedAt time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt time.Time `gorm:"type:timestamp;default:now()" json:"updated_at"`
}

// TableName returns the table name for GORM
func (r *RefreshTokenModel) TableName() string {
	return "refresh_tokens"
}

// GetID returns the primary key ID
func (r *RefreshTokenModel) GetID() uint {
	return r.ID
}

// SetID sets the primary key ID
func (r *RefreshTokenModel) SetID(id uint) {
	r.ID = id
}

// GetUserID returns the user ID associated with this token
func (r *RefreshTokenModel) GetUserID() string {
	return r.UserID
}

// SetUserID sets the user ID
func (r *RefreshTokenModel) SetUserID(userID string) {
	r.UserID = userID
}

// GetTokenID returns the unique token identifier
func (r *RefreshTokenModel) GetTokenID() string {
	return r.TokenID
}

// SetTokenID sets the token ID
func (r *RefreshTokenModel) SetTokenID(tokenID string) {
	r.TokenID = tokenID
}

// GetExpiresAt returns the expiration timestamp
func (r *RefreshTokenModel) GetExpiresAt() time.Time {
	return r.ExpiresAt
}

// SetExpiresAt sets the expiration timestamp
func (r *RefreshTokenModel) SetExpiresAt(expiresAt time.Time) {
	r.ExpiresAt = expiresAt
}

// GetCreatedAt returns the creation timestamp
func (r *RefreshTokenModel) GetCreatedAt() time.Time {
	return r.CreatedAt
}

// SetCreatedAt sets the creation timestamp
func (r *RefreshTokenModel) SetCreatedAt(createdAt time.Time) {
	r.CreatedAt = createdAt
}

// GetUpdatedAt returns the last update timestamp
func (r *RefreshTokenModel) GetUpdatedAt() time.Time {
	return r.UpdatedAt
}

// SetUpdatedAt sets the last update timestamp
func (r *RefreshTokenModel) SetUpdatedAt(updatedAt time.Time) {
	r.UpdatedAt = updatedAt
}

// Validate checks if the model has valid data
func (r *RefreshTokenModel) Validate() error {
	if r.UserID == "" {
		return fmt.Errorf("user_id cannot be empty")
	}
	if _, err := uuid.Parse(r.UserID); err != nil {
		return fmt.Errorf("user_id must be a valid UUID: %w", err)
	}
	if r.TokenID == "" {
		return fmt.Errorf("token_id cannot be empty")
	}
	if _, err := uuid.Parse(r.TokenID); err != nil {
		return fmt.Errorf("token_id must be a valid UUID: %w", err)
	}
	if r.ExpiresAt.IsZero() {
		return fmt.Errorf("expires_at cannot be zero")
	}
	if r.ExpiresAt.Before(time.Now()) {
		return fmt.Errorf("token has already expired")
	}
	return nil
}

// Sanitize prepares the model for safe output (removes sensitive data if needed)
func (r *RefreshTokenModel) Sanitize() {
	r.UpdatedAt = time.Now()
}

// IsExpired checks if the token has expired
func (r *RefreshTokenModel) IsExpired() bool {
	return time.Now().After(r.ExpiresAt)
}

// NewRefreshTokenModel creates a new refresh token model instance
func NewRefreshTokenModel(userID, tokenID string, expiresAt time.Time) *RefreshTokenModel {
	now := time.Now()
	return &RefreshTokenModel{
		UserID:    userID,
		TokenID:   tokenID,
		ExpiresAt: expiresAt,
		CreatedAt: now,
		UpdatedAt: now,
	}
}
