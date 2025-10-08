package models

import (
	"context"
	"time"

	svc "github.com/kubex-ecosystem/gdbase/factory"
	auth "github.com/kubex-ecosystem/gdbase/internal/models/auth"
)

// AuthRequestDTO represents authentication request data
type AuthRequestDTO struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Token type aliases for external use
type RefreshTokenModel = auth.RefreshTokenModel
type IRefreshToken = auth.IRefreshToken
type ITokenRepo = auth.ITokenRepo
type ITokenService = auth.ITokenService

// NewTokenService creates a new token service with the given repository
func NewTokenService(tokenRepo ITokenRepo) ITokenService {
	return auth.NewTokenService(tokenRepo)
}

// NewTokenRepo creates a new token repository using the provided DBService
func NewTokenRepo(ctx context.Context, dbService *svc.DBServiceImpl) ITokenRepo {
	return auth.NewTokenRepo(ctx, dbService)
}

// NewRefreshTokenModel creates a new refresh token model instance
func NewRefreshTokenModel(userID, tokenID string, expiresAt time.Time) *RefreshTokenModel {
	return auth.NewRefreshTokenModel(userID, tokenID, expiresAt)
}
