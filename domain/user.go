package domain

import (
	"legato_server/models"
)

type UserUseCase interface {
	RegisterNewUser(newUser models.NewUser) error
	Login(credential models.UserCredential) (string, error)
	RefreshUserToken(accessToken string) (models.RefreshToken, error)
	GetUserByEmail(email string) (models.UserInfo, error)
	GetUserByUsername(email string) (models.UserInfo, error)
	GetAllUsers() ([]*models.UserInfo, error)
	CreateDefaultUser() error
}
