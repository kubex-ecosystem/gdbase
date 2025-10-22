// Package registrationtokens provides the repository for registration tokens
package registrationtokens

import (
	"context"
	"fmt"
	"log"

	svc "github.com/kubex-ecosystem/gdbase/internal/services"
	"gorm.io/gorm"
)

// IRegistrationTokenRepo defines the interface for the registration token repository.
type IRegistrationTokenRepo interface {
	TableName() string
	Create(token IRegistrationToken) (IRegistrationToken, error)
	FindOne(where ...interface{}) (IRegistrationToken, error)
	Delete(where ...interface{}) error
}

// RegistrationTokenRepo implements the IRegistrationTokenRepo interface.
type RegistrationTokenRepo struct {
	g *gorm.DB
}

// NewRegistrationTokenRepo creates a new instance of the registration token repository.
func NewRegistrationTokenRepo(ctx context.Context, dbService *svc.DBServiceImpl) (IRegistrationTokenRepo, error) {
	if dbService == nil {
		log.Printf("ERROR: RegistrationToken repository: dbService is nil")
		return nil, fmt.Errorf("dbService is nil")
	}
	db, err := svc.GetDB(ctx, dbService)
	if err != nil {
		log.Printf("ERROR: RegistrationToken repository: failed to get DB from dbService: %v", err)
		return nil, err
	}

	// Auto-migrate the schema
	err = db.AutoMigrate(&RegistrationTokenModel{})
	if err != nil {
		log.Printf("ERROR: RegistrationToken repository: failed to auto-migrate schema: %v", err)
		return nil, err
	}

	return &RegistrationTokenRepo{g: db}, nil
}

func (r *RegistrationTokenRepo) TableName() string {
	return "registration_tokens"
}

func (r *RegistrationTokenRepo) Create(token IRegistrationToken) (IRegistrationToken, error) {
	if token == nil {
		return nil, fmt.Errorf("token is nil")
	}

	tokenModel := token.GetTokenObj()
	err := r.g.Create(tokenModel).Error
	if err != nil {
		return nil, fmt.Errorf("failed to create registration token: %w", err)
	}
	return tokenModel, nil
}

func (r *RegistrationTokenRepo) FindOne(where ...interface{}) (IRegistrationToken, error) {
	var tokenModel RegistrationTokenModel
	err := r.g.Where(where[0], where[1:]...).First(&tokenModel).Error
	if err != nil {
		return nil, fmt.Errorf("failed to find registration token: %w", err)
	}
	return &tokenModel, nil
}

func (r *RegistrationTokenRepo) Delete(where ...interface{}) error {
	err := r.g.Where(where[0], where[1:]...).Delete(&RegistrationTokenModel{}).Error
	if err != nil {
		return fmt.Errorf("failed to delete registration token: %w", err)
	}
	return nil
}
