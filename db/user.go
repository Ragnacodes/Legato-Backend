package legatoDb

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
	"strings"
	"time"
)

const FEMALE = "Female"
const MALE = "Male"

type User struct {
	gorm.Model
	UserID    uuid.UUID
	Username  string
	Email     string
	Password  string
	LastLogin time.Time
	FirstName string
	LastName  string
	Gender    string
}

func (u *User) String() string {
	return fmt.Sprintf("User: %v", u)
}

func (edb *LegatoDB) AddUser(u User) error {
	if pw, err := bcrypt.GenerateFromPassword([]byte(u.Password), 0); err != nil {
		return err
	} else {
		u.Password = string(pw)
	}

	user := &User{}
	edb.db.Where(&User{Username: u.Username}).First(&user)
	if user.Username == u.Username {
		return errors.New("this username is already taken")
	}

	user = &User{}
	edb.db.Where(&User{Email: u.Email}).First(&user)
	if user.Email == u.Email {
		return errors.New("this email is already taken")
	}

	edb.db.NewRecord(u) // => returns `true` as primary key is blank
	edb.db.Create(&u)

	return nil
}

func (edb *LegatoDB) GetUserByUsername(username string) (User, error) {
	user := User{}
	edb.db.Where(&User{Username: strings.ToLower(username)}).First(&user)
	if user.Username != username {
		return User{}, errors.New("username does not exist")
	}

	return user, nil
}

func (edb *LegatoDB) GetUserByEmail(email string) (User, error) {
	user := User{}
	edb.db.Where(&User{Email: strings.ToLower(email)}).First(&user)
	if user.Email != email {
		return User{}, errors.New("email does not exist")
	}

	return user, nil
}

func (edb *LegatoDB) FetchAllUsers() ([]User, error) {
	var users []User
	edb.db.Find(&users)

	if len(users) <= 0 {
		return users, errors.New("there is no user")
	}

	return users, nil
}
