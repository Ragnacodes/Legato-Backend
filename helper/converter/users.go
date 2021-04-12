package converter

import (
	"legato_server/api"
	"legato_server/db"
	"strings"
)

func NewUserToUserDb(newUser api.NewUser) legatoDb.User {
	u := legatoDb.User{}
	u.Username = strings.ToLower(newUser.Username)
	u.Email = strings.ToLower(newUser.Email)
	u.Password = newUser.Password

	return u
}

func UserDbToUser(ue legatoDb.User) api.UserInfo {
	u := api.UserInfo{}
	u.ID = ue.ID
	u.Email = strings.ToLower(ue.Email)
	u.Username = strings.ToLower(ue.Username)

	return u
}

func UserInfoToUserDb(u api.UserInfo) legatoDb.User {
	ue := legatoDb.User{}
	ue.Email = strings.ToLower(u.Email)
	ue.Username = strings.ToLower(u.Username)

	return ue
}
