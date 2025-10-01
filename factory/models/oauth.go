// Package models provides factory functions for OAuth models
package models

import (
	oauth "github.com/kubex-ecosystem/gdbase/internal/models/oauth"
	"gorm.io/gorm"
)

// OAuth Client type aliases - exported for external use
type (
	OAuthClient        = oauth.IOAuthClient
	OAuthClientModel   = oauth.OAuthClientModel
	OAuthClientRepo    = oauth.IOAuthClientRepo
	OAuthClientService = oauth.IOAuthClientService
)

// Auth Code type aliases - exported for external use
type (
	AuthCode        = oauth.IAuthCode
	AuthCodeModel   = oauth.AuthCodeModel
	AuthCodeRepo    = oauth.IAuthCodeRepo
	AuthCodeService = oauth.IAuthCodeService
)

// NewOAuthClientModel creates a new OAuth client model
func NewOAuthClientModel(clientID, clientName string, redirectURIs, scopes []string) OAuthClient {
	model := &oauth.OAuthClientModel{
		ClientID:   clientID,
		ClientName: clientName,
		Active:     true,
	}
	model.SetRedirectURIs(redirectURIs)
	model.SetScopes(scopes)
	return model
}

// NewOAuthClientRepo creates a new OAuth client repository
func NewOAuthClientRepo(db *gorm.DB) OAuthClientRepo {
	return oauth.NewOAuthClientRepo(db)
}

// NewOAuthClientService creates a new OAuth client service
func NewOAuthClientService(repo OAuthClientRepo) OAuthClientService {
	return oauth.NewOAuthClientService(repo)
}

// NewAuthCodeModel creates a new authorization code model
func NewAuthCodeModel(code, clientID, userID, redirectURI, codeChallenge, method string) AuthCode {
	return &oauth.AuthCodeModel{
		Code:                code,
		ClientID:            clientID,
		UserID:              userID,
		RedirectURI:         redirectURI,
		CodeChallenge:       codeChallenge,
		CodeChallengeMethod: method,
	}
}

// NewAuthCodeRepo creates a new authorization code repository
func NewAuthCodeRepo(db *gorm.DB) AuthCodeRepo {
	return oauth.NewAuthCodeRepo(db)
}

// NewAuthCodeService creates a new authorization code service
func NewAuthCodeService(repo AuthCodeRepo) AuthCodeService {
	return oauth.NewAuthCodeService(repo)
}
