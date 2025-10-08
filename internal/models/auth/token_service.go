// Package auth provides authentication service implementations
package auth

import (
	"context"
	"fmt"
	"time"
)

// ITokenService defines the interface for token service operations
type ITokenService interface {
	CreateRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	GetRefreshToken(ctx context.Context, tokenID string) (*RefreshTokenModel, error)
	ValidateRefreshToken(ctx context.Context, tokenID string) (*RefreshTokenModel, error)
	RevokeRefreshToken(ctx context.Context, userID string, tokenID string) error
	RevokeAllUserTokens(ctx context.Context, userID string) error
	CleanupExpiredTokens(ctx context.Context) error
}

// TokenServiceImpl implements ITokenService
type TokenServiceImpl struct {
	repo ITokenRepo
}

// NewTokenService creates a new token service instance
func NewTokenService(repo ITokenRepo) ITokenService {
	if repo == nil {
		panic("TokenRepo cannot be nil")
	}

	return &TokenServiceImpl{repo: repo}
}

// CreateRefreshToken creates a new refresh token for a user
func (s *TokenServiceImpl) CreateRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error {
	if userID == "" {
		return fmt.Errorf("user_id cannot be empty")
	}
	if tokenID == "" {
		return fmt.Errorf("token_id cannot be empty")
	}
	if expiresIn <= 0 {
		return fmt.Errorf("expiresIn must be positive")
	}

	if err := s.repo.SetRefreshToken(ctx, userID, tokenID, expiresIn); err != nil {
		return fmt.Errorf("service: failed to create refresh token: %w", err)
	}

	return nil
}

// GetRefreshToken retrieves a refresh token by its ID
func (s *TokenServiceImpl) GetRefreshToken(ctx context.Context, tokenID string) (*RefreshTokenModel, error) {
	if tokenID == "" {
		return nil, fmt.Errorf("token_id cannot be empty")
	}

	token, err := s.repo.GetRefreshToken(ctx, tokenID)
	if err != nil {
		return nil, fmt.Errorf("service: failed to get refresh token: %w", err)
	}

	if token == nil {
		return nil, nil
	}

	return token, nil
}

// ValidateRefreshToken validates a refresh token and checks if it's expired
func (s *TokenServiceImpl) ValidateRefreshToken(ctx context.Context, tokenID string) (*RefreshTokenModel, error) {
	token, err := s.GetRefreshToken(ctx, tokenID)
	if err != nil {
		return nil, err
	}

	if token == nil {
		return nil, fmt.Errorf("token not found")
	}

	if token.IsExpired() {
		return nil, fmt.Errorf("token has expired")
	}

	return token, nil
}

// RevokeRefreshToken revokes a specific refresh token
func (s *TokenServiceImpl) RevokeRefreshToken(ctx context.Context, userID string, tokenID string) error {
	if userID == "" {
		return fmt.Errorf("user_id cannot be empty")
	}
	if tokenID == "" {
		return fmt.Errorf("token_id cannot be empty")
	}

	if err := s.repo.DeleteRefreshToken(ctx, userID, tokenID); err != nil {
		return fmt.Errorf("service: failed to revoke refresh token: %w", err)
	}

	return nil
}

// RevokeAllUserTokens revokes all refresh tokens for a specific user
func (s *TokenServiceImpl) RevokeAllUserTokens(ctx context.Context, userID string) error {
	if userID == "" {
		return fmt.Errorf("user_id cannot be empty")
	}

	if err := s.repo.DeleteUserRefreshTokens(ctx, userID); err != nil {
		return fmt.Errorf("service: failed to revoke all user tokens: %w", err)
	}

	return nil
}

// CleanupExpiredTokens removes all expired tokens from the database
func (s *TokenServiceImpl) CleanupExpiredTokens(ctx context.Context) error {
	if err := s.repo.CleanupExpiredTokens(ctx); err != nil {
		return fmt.Errorf("service: failed to cleanup expired tokens: %w", err)
	}

	return nil
}
