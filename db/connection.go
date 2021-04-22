package legatoDb

import (
	"fmt"
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

func (ldb *LegatoDB) AddConnection(u *User, c Connection) (Connection, error) {
	c.UserID = u.ID
	con := Connection{}

	ldb.db.
		Where(&Connection{UserID: u.ID}).
		Where("Token = ?", c.Token).Find(&con)
	// check token does not exist
	if !strings.EqualFold(con.Token, c.Token) {
		ldb.db.Create(&c)
		ldb.db.Save(&c)
		return c, nil
	}
	return Connection{}, fmt.Errorf("this connection is already exist")
}

func (ldb *LegatoDB) GetUserConnections(u *User) ([]Connection, error) {
	user, _ := ldb.GetUserByUsername(u.Username)
	var connections []Connection
	ldb.db.Model(&user).Association("Connections").Find(&connections)

	return connections, nil
}

// func (ldb *LegatoDB) GetUserConnectionByName(u *User, name string) (Connection, error) {
// 	var connection Connection
// 	err := ldb.db.
// 		Where(&Scenario{UserID: u.ID}).
// 		Where("Name = ?", name).
// 		Find(&connection).Error
// 	return connection, err
// }

func (ldb *LegatoDB) GetUserConnectionById(u *User, id uint) (Connection, error) {
	var con Connection
	err := ldb.db.
		Where(&Connection{UserID: u.ID}).
		Where("ID = ?", id).Find(&con).Error
	if err != nil {
		conect := Connection{}
		conect.Token = "couldn' find connection"
		return Connection{}, fmt.Errorf("can not find connection")
	}
	return con, nil
}

func (ldb *LegatoDB) UpdateUserConnectionNameByID(u *User, name string, id uint) error {
	var connection Connection

	err := ldb.db.Take(&Connection{}).
		Where("ID = ?", id).Find(&connection).Error

	if err == nil {
		if connection.UserID == u.ID {
			connection.Name = name
			ldb.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&connection)
			ldb.db.Model(&connection).Update(connection.Token, name)
			ldb.db.Save(&connection)
		} else {
			return fmt.Errorf("there is no connection with this id for this user")
		}

	}
	return err
}

func (ldb *LegatoDB) CheckConnectionByID(u *User, id uint) error {

	var connection Connection
	err := ldb.db.Take(&Connection{}).
		Where("ID = ?", id).Find(&connection).Error
	if connection.UserID == u.ID && err == nil {
		return nil
	}

	return fmt.Errorf("there is no connection with this id for this user")
}

func (ldb *LegatoDB) DeleteConnectionByID(u *User, id uint) error {
	var connection Connection
	err := ldb.db.Take(&Connection{}).
		Where("ID = ?", id).Find(&connection).Error
	if u.ID == connection.UserID {
		ldb.db.Delete(&connection)
	} else {
		return fmt.Errorf("there is not a connection with this id for this user")
	}

	return err
}

func (ldb *LegatoDB) UpdateTokenFieldByID(u *User, Token string, id uint) error {
	var connection Connection

	err := ldb.db.Take(&Connection{}).
		Where("ID = ?", id).Find(&connection).Error

	if err == nil {
		if connection.UserID == u.ID {
			connection.Token = Token
			ldb.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&connection)
			ldb.db.Model(&connection).Update(connection.Token, Token)
			ldb.db.Save(&connection)
		} else {
			return fmt.Errorf("there is no connection with this id for this user")
		}

	}
	return err
}
