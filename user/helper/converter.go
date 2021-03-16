package helper

import (
	appdb "legato_server/db"
	"legato_server/models"
	"strings"
)

func NewUserToUserEntity(newUser models.NewUser) appdb.User {
	u := appdb.User{}
	u.FirstName = newUser.FirstName
	u.LastName = newUser.LastName
	u.Username = strings.ToLower(newUser.Username)
	u.Password = newUser.Password

	return u
}

func UserEntityToUser(ue appdb.User) models.User {
	u := models.User{}
	u.FirstName = ue.FirstName
	u.LastName = ue.LastName
	u.Email = strings.ToLower(ue.Email)
	u.Username = strings.ToLower(ue.Username)
	u.Gender = ue.Gender

	return u
}
