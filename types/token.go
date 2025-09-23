// Package types - token service types
package types

import (
	"context"
	"crypto/rsa"
	"time"

	ci "github.com/kubex-ecosystem/gdbase/internal/interfaces"
	sci "github.com/kubex-ecosystem/gdbase/internal/security/interfaces"
)

type TSConfig struct {
	TokenRepository       TokenRepo
	PrivKey               *rsa.PrivateKey
	PubKey                *rsa.PublicKey
	RefreshSecret         string
	IDExpirationSecs      int64
	RefreshExpirationSecs int64
	KeyringPass           string
	TokenClient           TokenClient
	DBService             *IDBService
	KeyringService        sci.IKeyringService
}
type TokenPair struct {
	IDToken
	RefreshToken
}
type RefreshToken struct {
	ID  string `json:"-"`
	UID string `json:"-"`
	SS  string
}
type IDToken struct {
	SS string
}

type TokenClient interface {
	LoadPrivateKey() (*rsa.PrivateKey, error)
	LoadPublicKey() *rsa.PublicKey
	LoadTokenCfg() (TokenService, int64, int64, error)
}
type TokenRepo interface {
	SetRefreshToken(ctx context.Context, userID string, tokenID string, expiresIn time.Duration) error
	DeleteRefreshToken(ctx context.Context, userID string, prevTokenID string) error
	DeleteUserRefreshTokens(ctx context.Context, userID string) error
}

type TokenService interface {
	NewPairFromUser(ctx context.Context, u ci.User, prevTokenID string) (*TokenPair, error)
	SignOut(ctx context.Context, uid string) error
	ValidateIDToken(tokenString string) (ci.User, error)
	ValidateRefreshToken(refreshTokenString string) (*RefreshToken, error)
	RenewToken(ctx context.Context, refreshToken string) (*TokenPair, error)
}
