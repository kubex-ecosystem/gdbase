// Package oauth provides OAuth2 client and authorization code models
package oauth

import (
	"encoding/json"
	"fmt"
	"time"
)

// IOAuthClient interface for abstraction and encapsulation
//
//go:generate mockgen -destination=../../mocks/mock_oauth_client.go -package=mocks . IOAuthClient
type IOAuthClient interface {
	TableName() string
	GetID() string
	SetID(id string)
	GetClientID() string
	SetClientID(clientID string)
	GetClientName() string
	SetClientName(name string)
	GetRedirectURIs() []string
	SetRedirectURIs(uris []string)
	GetScopes() []string
	SetScopes(scopes []string)
	GetActive() bool
	SetActive(active bool)
	GetCreatedAt() time.Time
	SetCreatedAt(createdAt time.Time)
	GetUpdatedAt() time.Time
	SetUpdatedAt(updatedAt time.Time)
	Validate() error
	Sanitize()
	IsRedirectURIAllowed(uri string) bool
	HasScope(scope string) bool
}

// OAuthClientModel represents an OAuth2 client application
type OAuthClientModel struct {
	ID           string    `gorm:"type:uuid;primaryKey" json:"id"`
	ClientID     string    `gorm:"type:varchar(255);unique;not null;index" json:"client_id"`
	ClientName   string    `gorm:"type:varchar(255);not null" json:"client_name"`
	RedirectURIs string    `gorm:"type:text" json:"-"` // JSON array stored as text
	Scopes       string    `gorm:"type:text" json:"-"` // JSON array stored as text
	Active       bool      `gorm:"type:boolean;default:true" json:"active"`
	CreatedAt    time.Time `gorm:"type:timestamp;default:now()" json:"created_at"`
	UpdatedAt    time.Time `gorm:"type:timestamp;default:now()" json:"updated_at"`

	// Cached parsed values (not stored in DB)
	redirectURIsParsed []string `gorm:"-" json:"-"`
	scopesParsed       []string `gorm:"-" json:"-"`
}

func (oac *OAuthClientModel) TableName() string {
	return "oauth_clients"
}

func (oac *OAuthClientModel) GetID() string {
	return oac.ID
}

func (oac *OAuthClientModel) SetID(id string) {
	oac.ID = id
}

func (oac *OAuthClientModel) GetClientID() string {
	return oac.ClientID
}

func (oac *OAuthClientModel) SetClientID(clientID string) {
	oac.ClientID = clientID
}

func (oac *OAuthClientModel) GetClientName() string {
	return oac.ClientName
}

func (oac *OAuthClientModel) SetClientName(name string) {
	oac.ClientName = name
}

func (oac *OAuthClientModel) GetRedirectURIs() []string {
	if oac.redirectURIsParsed != nil {
		return oac.redirectURIsParsed
	}
	if oac.RedirectURIs == "" {
		return []string{}
	}
	var uris []string
	if err := json.Unmarshal([]byte(oac.RedirectURIs), &uris); err != nil {
		return []string{}
	}
	oac.redirectURIsParsed = uris
	return uris
}

func (oac *OAuthClientModel) SetRedirectURIs(uris []string) {
	oac.redirectURIsParsed = uris
	if len(uris) == 0 {
		oac.RedirectURIs = "[]"
		return
	}
	data, err := json.Marshal(uris)
	if err != nil {
		oac.RedirectURIs = "[]"
		return
	}
	oac.RedirectURIs = string(data)
}

func (oac *OAuthClientModel) GetScopes() []string {
	if oac.scopesParsed != nil {
		return oac.scopesParsed
	}
	if oac.Scopes == "" {
		return []string{}
	}
	var scopes []string
	if err := json.Unmarshal([]byte(oac.Scopes), &scopes); err != nil {
		return []string{}
	}
	oac.scopesParsed = scopes
	return scopes
}

func (oac *OAuthClientModel) SetScopes(scopes []string) {
	oac.scopesParsed = scopes
	if len(scopes) == 0 {
		oac.Scopes = "[]"
		return
	}
	data, err := json.Marshal(scopes)
	if err != nil {
		oac.Scopes = "[]"
		return
	}
	oac.Scopes = string(data)
}

func (oac *OAuthClientModel) GetActive() bool {
	return oac.Active
}

func (oac *OAuthClientModel) SetActive(active bool) {
	oac.Active = active
}

func (oac *OAuthClientModel) GetCreatedAt() time.Time {
	return oac.CreatedAt
}

func (oac *OAuthClientModel) SetCreatedAt(createdAt time.Time) {
	oac.CreatedAt = createdAt
}

func (oac *OAuthClientModel) GetUpdatedAt() time.Time {
	return oac.UpdatedAt
}

func (oac *OAuthClientModel) SetUpdatedAt(updatedAt time.Time) {
	oac.UpdatedAt = updatedAt
}

func (oac *OAuthClientModel) Validate() error {
	if oac.ClientID == "" {
		return fmt.Errorf("client_id is required")
	}
	if oac.ClientName == "" {
		return fmt.Errorf("client_name is required")
	}
	uris := oac.GetRedirectURIs()
	if len(uris) == 0 {
		return fmt.Errorf("at least one redirect_uri is required")
	}
	return nil
}

func (oac *OAuthClientModel) Sanitize() {
	// Parse JSON fields to ensure cache is populated
	oac.GetRedirectURIs()
	oac.GetScopes()
}

func (oac *OAuthClientModel) IsRedirectURIAllowed(uri string) bool {
	uris := oac.GetRedirectURIs()
	for _, allowed := range uris {
		if allowed == uri {
			return true
		}
	}
	return false
}

func (oac *OAuthClientModel) HasScope(scope string) bool {
	scopes := oac.GetScopes()
	for _, s := range scopes {
		if s == scope {
			return true
		}
	}
	return false
}
