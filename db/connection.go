package legatoDb

import (
	"fmt"

	"gorm.io/gorm"
)

type Connection struct {
	gorm.Model
	UserID uint
	Name   string
	Data   string
	Type   string
}

func (ldb *LegatoDB) AddConnection(u *User, c Connection) (Connection, error) {
	c.UserID = u.ID
	ldb.db.Create(&c)
	err := ldb.db.Save(&c).Error
	return c, err
}

func (ldb *LegatoDB) GetUserConnections(u *User) ([]Connection, error) {
	user, _ := ldb.GetUserByUsername(u.Username)
	var connections []Connection
	_ = ldb.db.Model(&user).Association("Connections").Find(&connections)

	return connections, nil
}

func (ldb *LegatoDB) GetUserConnectionById(u *User, id uint) (Connection, error) {
	var con Connection
	err := ldb.db.
		Where(&Connection{UserID: u.ID}).
		Where("ID = ?", id).Find(&con).Error
	if err != nil {
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

func (ldb *LegatoDB) UpdateDataFieldByID(u *User, data string, id uint) error {
	var connection Connection

	err := ldb.db.Take(&Connection{}).
		Where("ID = ?", id).Find(&connection).Error

	if err == nil {
		if connection.UserID == u.ID {
			connection.Data = data
			ldb.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&connection)
			ldb.db.Save(&connection)
		} else {
			return fmt.Errorf("there is no connection with this id for this user")
		}

	}
	return err
}
