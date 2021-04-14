package domain

import (
	"legato_server/api"
)

type UserUseCase interface {
	RegisterNewUser(newUser api.NewUser) error
	Login(credential api.UserCredential) (string, error)
	RefreshUserToken(accessToken string) (api.RefreshToken, error)
	GetUserByEmail(email string) (api.UserInfo, error)
	GetUserByUsername(email string) (api.UserInfo, error)
	GetAllUsers() ([]*api.UserInfo, error)
	CreateDefaultUser() error
}
