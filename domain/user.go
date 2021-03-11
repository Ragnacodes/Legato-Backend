package domain

import (
	"legato_server/models"
)

type UserUseCase interface {
	RegisterNewUser(newUser models.NewUser) error
	Login(credential models.UserCredential) (string, error)
	RefreshUserToken(refreshToken models.RefreshToken) (string, error)
	GetUserByEmail(email string) (models.User, error)
	GetUserByUsername(email string) (models.User, error)
	GetAllUsers() ([]*models.User, error)
	CreateDefaultUser() error
}
