// Package models provides factory functions for OAuth models
package models

import (
	"context"

	oauth "github.com/kubex-ecosystem/gdbase/internal/models/oauth"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
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
func NewOAuthClientRepo(ctx context.Context, dbService *svc.DBService, dbName string) OAuthClientRepo {
	return oauth.NewOAuthClientRepo(ctx, dbService, dbName)
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
func NewAuthCodeRepo(ctx context.Context, dbService *svc.DBService, dbName string) AuthCodeRepo {
	if dbService == nil {
		return nil
	}
	db, err := dbService.GetDB(ctx, dbName)
	if err != nil {
		return nil
	}
	if db == nil {
		return nil
	}
	return oauth.NewAuthCodeRepo(db)
}

// NewAuthCodeService creates a new authorization code service
func NewAuthCodeService(repo AuthCodeRepo) AuthCodeService {
	return oauth.NewAuthCodeService(repo)
}
