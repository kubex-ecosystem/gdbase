// Package registrationtokens provides the service for registration tokens
package registrationtokens

import (
	"fmt"
	"time"
)

// IRegistrationTokenService defines the interface for the registration token service.
type IRegistrationTokenService interface {
	CreateToken(token IRegistrationToken) (IRegistrationToken, error)
	GetToken(token string) (IRegistrationToken, error)
	DeleteToken(token string) error
	DeleteUserTokens(userID string) error
}

// RegistrationTokenService implements the IRegistrationTokenService interface.
type RegistrationTokenService struct {
	repo IRegistrationTokenRepo
}

// NewRegistrationTokenService creates a new instance of the registration token service.
func NewRegistrationTokenService(repo IRegistrationTokenRepo) IRegistrationTokenService {
	return &RegistrationTokenService{repo: repo}
}

func (s *RegistrationTokenService) CreateToken(token IRegistrationToken) (IRegistrationToken, error) {
	if token.GetUserID() == "" || token.GetToken() == "" || token.GetExpiresAt().IsZero() {
		return nil, fmt.Errorf("missing required fields")
	}

	// Ensure token is not expired
	if token.GetExpiresAt().Before(time.Now()) {
		return nil, fmt.Errorf("token expiration date cannot be in the past")
	}

	return s.repo.Create(token)
}

func (s *RegistrationTokenService) GetToken(token string) (IRegistrationToken, error) {
	return s.repo.FindOne("token = ?", token)
}

func (s *RegistrationTokenService) DeleteToken(token string) error {
	return s.repo.Delete("token = ?", token)
}

func (s *RegistrationTokenService) DeleteUserTokens(userID string) error {
	return s.repo.Delete("user_id = ?", userID)
}
