package legatoDb

import (
	"fmt"

	"gorm.io/gorm"
)

type Connection struct {
	gorm.Model
	Token      string
	Token_type string
	UserID     uint
	Name       string
}

func (ldb *LegatoDB) AddConnection(u *User, c Connection) error {
	c.UserID = u.ID
	ldb.db.Create(&c)
	ldb.db.Save(&c)
	return nil
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

func (ldb *LegatoDB) UpdateUserConnectionByID(u *User, name string, id uint) error {
	var connection Connection

	err := ldb.db.Take(&Connection{}).
		Where("ID = ?", id).Find(&connection).Error

	if err == nil && connection.UserID == u.ID {
		connection.Name = name
		ldb.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&connection)
		ldb.db.Model(&connection).Update(connection.Token, name)
		ldb.db.Save(&connection)
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

func (ldb *LegatoDB) DeleteConnectionByID(id uint) error {
	// ldb.db.Delete(&Connection{}, id)
	// err := ldb.db.Take(&Connection{}).Where("id = ?", id).Delete(&Connection{}).Error
	var connection Connection
	err := ldb.db.Take(&Connection{}).
		Where("ID = ?", id).Find(&connection).Error
	ldb.db.Delete(&connection)
	return err
}

func (ldb *LegatoDB) UpdateTokenFieldByID(u *User, Token string, id uint) error {
	var connection Connection

	err := ldb.db.Take(&Connection{}).
		Where("ID = ?", id).Find(&connection).Error

	if err == nil && connection.UserID == u.ID {
		connection.Token = Token
		ldb.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&connection)
		ldb.db.Model(&connection).Update(connection.Token, Token)
		ldb.db.Save(&connection)
	}
	return err
}
