package domain

import (
	"legato_server/api"
	legatoDb "legato_server/db"
)

type UserUseCase interface {
	RegisterNewUser(newUser api.NewUser) error
	Login(credential api.UserCredential) (string, error)
	RefreshUserToken(accessToken string) (api.RefreshToken, error)
	GetUserByEmail(email string) (api.UserInfo, error)
	GetUserByUsername(email string) (api.UserInfo, error)
	GetAllUsers() ([]*api.UserInfo, error)
	CreateDefaultUser() error
	AddConnectionToDB(name string, ut api.Connection) (api.Connection, error)
	GetConnectionByID(username string, id uint) (legatoDb.Connection, error)
	GetConnections(username string) ([]legatoDb.Connection, error)
	UpdateUserConnectionNameById(username string, ut api.Connection) error
	CheckConnectionByID(username string, id uint) error
	DeleteUserConnectionById(username string, id uint) error
	UpdateTokenFieldByID(username string, ut api.Connection) error
}
