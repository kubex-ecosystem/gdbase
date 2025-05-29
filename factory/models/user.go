package models

import (
	m "github.com/rafa-mori/gdbase/internal/models/users"
	"gorm.io/gorm"
)

type UserModelType = m.UserModel
type UserModel = m.IUser
type UserService = m.IUserService
type UserRepo = m.IUserRepo

func NewUserService(userRepo UserRepo) UserService {
	return m.NewUserService(userRepo)
}

func NewUserRepo(db *gorm.DB) UserRepo {
	return m.NewUserRepo(db)
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
