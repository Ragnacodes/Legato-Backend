package helper

import (
	appdb "legato_server/db"
	"legato_server/models"
	"strings"
)

func NewUserToUserEntity(newUser models.NewUser) appdb.User {
	u := appdb.User{}
	u.Username = strings.ToLower(newUser.Username)
	u.Email = strings.ToLower(newUser.Email)
	u.Password = newUser.Password

	return u
}

func UserEntityToUser(ue appdb.User) models.UserInfo {
	u := models.UserInfo{}
	u.Email = strings.ToLower(ue.Email)
	u.Username = strings.ToLower(ue.Username)

 	return u
}
