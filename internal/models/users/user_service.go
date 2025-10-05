package user

import (
	"errors"
	"fmt"

	is "github.com/kubex-ecosystem/gdbase/internal/services"
)

type IUserService interface {
	CreateUser(user IUser) (IUser, error)
	GetUserByID(id string) (IUser, error)
	UpdateUser(user IUser) (IUser, error)
	DeleteUser(id string) error
	ListUsers() ([]IUser, error)
	GetUserByEmail(email string) (IUser, error)
	GetUserByUsername(username string) (IUser, error)
	GetUserByPhone(phone string) (IUser, error)
	GetContextDBService() *is.DBServiceImpl
}

type UserService struct {
	repo IUserRepo
}

func NewUserService(repo IUserRepo) IUserService {
	return &UserService{repo: repo}
}

func (us *UserService) CreateUser(user IUser) (IUser, error) {
	if user.GetUsername() == "" || user.GetEmail() == "" || user.GetPassword() == "" {
		return nil, errors.New("missing required fields")
	}
	createdUser, err := us.repo.Create(user)
	if err != nil {
		return nil, fmt.Errorf("error creating user: %w", err)
	}
	return createdUser, nil
}

func (us *UserService) GetUserByID(id string) (IUser, error) {
	user, err := us.repo.FindOne("id = ?", id)
	if err != nil {
		return nil, fmt.Errorf("error fetching user: %w", err)
	}
	return user, nil
}

func (us *UserService) UpdateUser(user IUser) (IUser, error) {
	updatedUser, err := us.repo.Update(user)
	if err != nil {
		return nil, fmt.Errorf("error updating user: %w", err)
	}
	return updatedUser, nil
}

func (us *UserService) DeleteUser(id string) error {
	err := us.repo.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting user: %w", err)
	}
	return nil

}

func (us *UserService) ListUsers() ([]IUser, error) {
	users, err := us.repo.FindAll("active = ?", true)
	if err != nil {
		return nil, fmt.Errorf("error listing users: %w", err)
	}
	return users, nil
}

func (us *UserService) GetUserByEmail(email string) (IUser, error) {
	user, err := us.repo.FindOne("email = ?", email)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by email: %w", err)
	}
	return user, nil
}

func (us *UserService) GetUserByUsername(username string) (IUser, error) {
	user, err := us.repo.FindOne("username = ?", username)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by username: %w", err)
	}
	return user, nil
}

func (us *UserService) GetUserByPhone(phone string) (IUser, error) {
	user, err := us.repo.FindOne("phone = ?", phone)
	if err != nil {
		return nil, fmt.Errorf("error fetching user by phone: %w", err)
	}
	return user, nil
}

func (us *UserService) GetContextDBService() *is.DBServiceImpl {
	return us.repo.GetContextDBService()
}
