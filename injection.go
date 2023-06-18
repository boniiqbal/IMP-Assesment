//go:build wireinject
// +build wireinject

package main

import (
	"imp-backend/application/use_case/auth/login"
	"imp-backend/application/use_case/auth/signup"
	"imp-backend/application/use_case/user/list_user"

	"imp-backend/infrastructure/persistence/repository"

	"github.com/google/wire"
	"gorm.io/gorm"
)

func LoginHandler(db *gorm.DB) login.LoginHandler {
	wire.Build(login.NewLoginHandler, login.NewLoginService, repository.NewAuthRepository)
	return login.LoginHandler{}
}

func SignupHandler(db *gorm.DB) signup.SignupHandler {
	wire.Build(signup.NewSignupHandler, signup.NewSignupService, repository.NewAuthRepository)
	return signup.SignupHandler{}
}

func ListUserHandler(db *gorm.DB) list_user.ListUserHandler {
	wire.Build(list_user.NewListUserHandler, list_user.NewListUserService, repository.NewAuthRepository)
	return list_user.ListUserHandler{}
}
