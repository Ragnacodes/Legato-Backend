package legatoDb

import (
	"fmt"
	"legato_server/models"
	"strings"

	"gorm.io/gorm"
)

type Connection struct {
	gorm.Model
	Token      string
	Token_type string
	UserID     uint
	Name       string
}

func (ldb *LegatoDB) AddToken(u *User, c Connection) error {
	c.UserID = u.ID
	ldb.db.Create(&c)
	ldb.db.Save(&c)
	return nil
}

func (ldb *LegatoDB) GetUserTokens(u *User, ut models.UserGetToken) ([]Connection, error) {
	user, _ := ldb.GetUserByUsername(u.Username)

	var connections []Connection
	ldb.db.Model(&user).Association("Connections").Find(&connections)

	return connections, nil
}

func (ldb *LegatoDB) GetUserToken(u *User, ut models.UserGetToken) (Connection, error) {
	user, _ := ldb.GetUserByUsername(u.Username)

	var connections []Connection
	var connection Connection
	ldb.db.Model(&user).Association("Connections").Find(&connections)
	for _, con := range connections {
		if strings.EqualFold(con.Name, ut.Name) && strings.EqualFold(con.Token_type, ut.Token_type) {
			connection = con
		} else {
			connection.Token = "can not find token"
		}
	}
	return connection, nil
}

func (ldb *LegatoDB) GetUserTokenById(u *User, name string) (Connection, error) {
	var con Connection
	err := ldb.db.
		Where(&Connection{UserID: u.ID}).
		Where("Name = ?", name).
		Preload("Token").Find(&con).Error
	if err != nil {
		conect := Connection{}
		conect.Token = "couldn' find toke"
		return Connection{}, fmt.Errorf("can not find token")
	}

	return con, nil
}

func (ldb *LegatoDB) UpdateUserTokenById(u *User, name string, ut models.UserGetToken, id uint) error {
	var connection Connection
	var connections []Connection
	ldb.db.Model(&u).Association("Connections").Find(&connections)
	flag := false
	for _, con := range connections {
		if con.ID == id {
			connection = con
			flag = true
		}
	}
	if flag == false {

		return fmt.Errorf("can not find token")

	}
	connection.Name = name
	ldb.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&connection)
	ldb.db.Model(&connection).Update(connection.Name, name)
	ldb.db.Save(&connection)
	return nil
}

func (ldb *LegatoDB) CheckTokenByID(user *User, id uint) error {

	var connections []Connection
	ldb.db.Model(&user).Association("Connections").Find(&connections)
	flag := false
	for _, con := range connections {
		if con.ID == id {
			flag = true
		}
	}
	if flag == false {

		return fmt.Errorf("can not find token")

	}
	return nil
}

func (ldb *LegatoDB) DeleteConnectionByID(u *User, ut models.UserGetToken, id uint) error {
	var connection Connection
	var connections []Connection
	ldb.db.Model(&u).Association("Connections").Find(&connections)
	flag := false
	for _, con := range connections {
		if con.ID == id {
			connection = con
			flag = true
		}
	}
	if flag == false {

		return fmt.Errorf("can not find token")

	}
	ldb.db.Delete(&Connection{}, id)
	ldb.db.Save(&connection)
	return nil
}

func (ldb *LegatoDB) UpdateUserTokenByName(u *User, Token string, ut models.UserGetToken, id uint) error {
	var connection Connection
	var connections []Connection
	ldb.db.Model(&u).Association("Connections").Find(&connections)
	flag := false
	for _, con := range connections {
		if con.ID == id {
			connection = con
			flag = true
		}
	}
	if flag == false {

		return fmt.Errorf("can not find token")

	}
	connection.Token = Token
	ldb.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&connection)
	ldb.db.Model(&connection).Update(connection.Token, Token)
	ldb.db.Save(&connection)
	return nil
}
