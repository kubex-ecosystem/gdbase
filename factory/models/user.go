package models

import (
	"context"

	svc "github.com/kubex-ecosystem/gdbase/factory"
	m "github.com/kubex-ecosystem/gdbase/internal/models/users"
)

type UserModelType = m.UserModel
type UserModel = m.IUser
type UserService = m.IUserService
type UserRepo = m.IUserRepo

func NewUserService(userRepo UserRepo) UserService {
	return m.NewUserService(userRepo)
}

func NewUserRepo(ctx context.Context, dbService *svc.DBServiceImpl) UserRepo {
	return m.NewUserRepo(ctx, dbService)
}

func NewUserModel(username, name, email string) UserModel {
	return &m.UserModel{
		Username: username,
		Name:     name,
		Email:    email,
		Password: "",
		Active:   true,
	}
}
