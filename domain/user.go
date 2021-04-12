package domain

import (
	legatoDb "legato_server/db"
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
	AddTokenDB(name string, ut models.UserAddToken) error
	GetTokenByUsername(username string, ut models.UserGetToken) (legatoDb.Connection, error)
	GetTokensByUsername(username string, ut models.UserGetToken) ([]legatoDb.Connection, error)
	UpdateUserTokenById(username string, ut models.UserGetToken) error
	CheckTokenByID(username string, ut models.UserGetToken) error
	DeleteUserTokenById(username string, ut models.UserGetToken) error
	UpdateUserTokenByName(username string, ut models.UserGetToken) error
}
