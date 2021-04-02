package converter

import (
	"legato_server/db"
	"legato_server/models"
	"strings"
)

func NewUserToUserDb(newUser models.NewUser) legatoDb.User {
	u := legatoDb.User{}
	u.Username = strings.ToLower(newUser.Username)
	u.Email = strings.ToLower(newUser.Email)
	u.Password = newUser.Password

	return u
}

func UserDbToUser(ue legatoDb.User) models.UserInfo {
	u := models.UserInfo{}
	u.Email = strings.ToLower(ue.Email)
	u.Username = strings.ToLower(ue.Username)

	return u
}

func UserInfoToUserDb(u models.UserInfo) legatoDb.User {
	ue := legatoDb.User{}
	ue.Email = strings.ToLower(u.Email)
	ue.Username = strings.ToLower(u.Username)

	return ue
}