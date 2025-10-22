// Package models provides factory functions for the models in the application.
package models

import (
	"context"
	"time"

	"github.com/kubex-ecosystem/gdbase/internal/models/registrationtokens"
	svc "github.com/kubex-ecosystem/gdbase/internal/services"
)

// IRegistrationToken is an alias for the internal interface.
type IRegistrationToken = registrationtokens.IRegistrationToken

// RegistrationTokenModel is an alias for the internal model.
type RegistrationTokenModel = registrationtokens.RegistrationTokenModel

// NewRegistrationToken is a factory function to create a new registration token model.
func NewRegistrationToken(userID, token string, expiresAt time.Time) *RegistrationTokenModel {
	return registrationtokens.NewRegistrationToken(userID, token, expiresAt)
}

// IRegistrationTokenRepo is an alias for the internal interface.
type IRegistrationTokenRepo = registrationtokens.IRegistrationTokenRepo

// NewRegistrationTokenRepo is a factory function to create a new registration token repository.
func NewRegistrationTokenRepo(ctx context.Context, dbService *svc.DBServiceImpl) (IRegistrationTokenRepo, error) {
	return registrationtokens.NewRegistrationTokenRepo(ctx, dbService)
}

// IRegistrationTokenService is an alias for the internal interface.
type IRegistrationTokenService = registrationtokens.IRegistrationTokenService

// NewRegistrationTokenService is a factory function to create a new registration token service.
func NewRegistrationTokenService(repo IRegistrationTokenRepo) IRegistrationTokenService {
	return registrationtokens.NewRegistrationTokenService(repo)
}
