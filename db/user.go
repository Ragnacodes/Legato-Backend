package legatoDb

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type User struct {
	gorm.Model
	Username  string
	Email     string
	Password  string
	LastLogin time.Time
	Scenarios []Scenario
	Services []Service
}

func (u *User) String() string {
	return fmt.Sprintf("(@User: %+v)", *u)
}

func (ldb *LegatoDB) AddUser(u User) error {
	var user *User

	// Encode the user password
	if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0); err != nil {
		return err
	} else {
		u.Password = string(pw)
	}

	// Check unique username
	user = &User{}
	ldb.db.Where(&User{Username: u.Username}).Find(&user)
	if user.Username == u.Username {
		return errors.New("this username is already taken")
	}

	// Check unique user email
	user = &User{}
	ldb.db.Where(&User{Email: u.Email}).Find(&user)
	if user.Email == u.Email {
		return errors.New("this email is already taken")
	}

	// Set initial values for new user
	u.LastLogin = time.Now()

	ldb.db.Create(&u)
	ldb.db.Save(u)

	return nil
}

func (ldb *LegatoDB) GetUserByUsername(username string) (User, error) {
	user := User{}
	ldb.db.Where(&User{Username: strings.ToLower(username)}).First(&user)
	if user.Username != username {
		return User{}, errors.New("username does not exist")
	}

	return user, nil
}

func (ldb *LegatoDB) GetUserByEmail(email string) (User, error) {
	user := User{}
	ldb.db.Where(&User{Email: strings.ToLower(email)}).First(&user)
	if user.Email != email {
		return User{}, errors.New("email does not exist")
	}

	return user, nil
}

func (ldb *LegatoDB) FetchAllUsers() ([]User, error) {
	var users []User
	ldb.db.Find(&users)

	if len(users) <= 0 {
		return users, errors.New("there is no user")
	}

	return users, nil
}
