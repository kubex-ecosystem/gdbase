package oauth

import (
	"fmt"
	"time"
)

// IAuthCode interface for abstraction and encapsulation
//
//go:generate mockgen -destination=../../mocks/mock_auth_code.go -package=mocks . IAuthCode
type IAuthCode interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetCode() string
	SetCode(code string)
	GetClientID() string
	SetClientID(clientID string)
	GetUserID() string
	SetUserID(userID string)
	GetRedirectURI() string
	SetRedirectURI(uri string)
	GetCodeChallenge() string
	SetCodeChallenge(challenge string)
	GetCodeChallengeMethod() string
	SetCodeChallengeMethod(method string)
	GetScope() string
	SetScope(scope string)
	GetExpiresAt() time.Time
	SetExpiresAt(expiresAt time.Time)
	GetUsed() bool
	SetUsed(used bool)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	Validate() error
	IsExpired() bool
	IsValid() bool
}

// AuthCodeModel represents an OAuth2 authorization code with PKCE support
type AuthCodeModel struct {
	ID                  string    `gorm:"type:uuid;primaryKey" json:"id"`
	Code                string    `gorm:"type:varchar(255);unique;not null;index" json:"code"`
	ClientID            string    `gorm:"type:varchar(255);not null;index" json:"client_id"`
	UserID              string    `gorm:"type:uuid;not null;index" json:"user_id"`
	RedirectURI         string    `gorm:"type:text;not null" json:"redirect_uri"`
	CodeChallenge       string    `gorm:"type:varchar(255);not null" json:"code_challenge"`
	CodeChallengeMethod string    `gorm:"type:varchar(10);default:'S256'" json:"code_challenge_method"`
	Scope               string    `gorm:"type:text" json:"scope"`
	ExpiresAt           time.Time `gorm:"type:timestamp;not null;index" json:"expires_at"`
	Used                bool      `gorm:"type:boolean;default:false;index" json:"used"`
	CreatedAt           time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
}

func (ac *AuthCodeModel) TableName() string {
	return "oauth_authorization_codes"
}

func (ac *AuthCodeModel) GetID() string {
	return ac.ID
}

func (ac *AuthCodeModel) SetID(id string) {
	ac.ID = id
}

func (ac *AuthCodeModel) GetCode() string {
	return ac.Code
}

func (ac *AuthCodeModel) SetCode(code string) {
	ac.Code = code
}

func (ac *AuthCodeModel) GetClientID() string {
	return ac.ClientID
}

func (ac *AuthCodeModel) SetClientID(clientID string) {
	ac.ClientID = clientID
}

func (ac *AuthCodeModel) GetUserID() string {
	return ac.UserID
}

func (ac *AuthCodeModel) SetUserID(userID string) {
	ac.UserID = userID
}

func (ac *AuthCodeModel) GetRedirectURI() string {
	return ac.RedirectURI
}

func (ac *AuthCodeModel) SetRedirectURI(uri string) {
	ac.RedirectURI = uri
}

func (ac *AuthCodeModel) GetCodeChallenge() string {
	return ac.CodeChallenge
}

func (ac *AuthCodeModel) SetCodeChallenge(challenge string) {
	ac.CodeChallenge = challenge
}

func (ac *AuthCodeModel) GetCodeChallengeMethod() string {
	return ac.CodeChallengeMethod
}

func (ac *AuthCodeModel) SetCodeChallengeMethod(method string) {
	ac.CodeChallengeMethod = method
}

func (ac *AuthCodeModel) GetScope() string {
	return ac.Scope
}

func (ac *AuthCodeModel) SetScope(scope string) {
	ac.Scope = scope
}

func (ac *AuthCodeModel) GetExpiresAt() time.Time {
	return ac.ExpiresAt
}

func (ac *AuthCodeModel) SetExpiresAt(expiresAt time.Time) {
	ac.ExpiresAt = expiresAt
}

func (ac *AuthCodeModel) GetUsed() bool {
	return ac.Used
}

func (ac *AuthCodeModel) SetUsed(used bool) {
	ac.Used = used
}

func (ac *AuthCodeModel) GetCreatedAt() time.Time {
	return ac.CreatedAt
}

func (ac *AuthCodeModel) SetCreatedAt(createdAt time.Time) {
	ac.CreatedAt = createdAt
}

func (ac *AuthCodeModel) Validate() error {
	if ac.Code == "" {
		return fmt.Errorf("code is required")
	}
	if ac.ClientID == "" {
		return fmt.Errorf("client_id is required")
	}
	if ac.UserID == "" {
		return fmt.Errorf("user_id is required")
	}
	if ac.RedirectURI == "" {
		return fmt.Errorf("redirect_uri is required")
	}
	if ac.CodeChallenge == "" {
		return fmt.Errorf("code_challenge is required")
	}
	if ac.CodeChallengeMethod != "S256" && ac.CodeChallengeMethod != "plain" {
		return fmt.Errorf("code_challenge_method must be 'S256' or 'plain'")
	}
	if ac.ExpiresAt.IsZero() {
		return fmt.Errorf("expires_at is required")
	}
	return nil
}

func (ac *AuthCodeModel) IsExpired() bool {
	return time.Now().After(ac.ExpiresAt)
}

func (ac *AuthCodeModel) IsValid() bool {
	return !ac.Used && !ac.IsExpired()
}
