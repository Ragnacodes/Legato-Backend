package usecase

import (
	"legato_server/authenticate"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/models"
	"legato_server/user/helper"
	"log"
	"time"
)

var defaultUser = legatoDb.User{
	Username: "legato",
	Email:    "legato@gmail.com",
	Password: "1234qwer",
}

type userUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewUserUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.UserUseCase {
	return &userUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

// Register new user and add it in our database
func (u *userUseCase) RegisterNewUser(nu models.NewUser) error {
	err := u.db.AddUser(helper.NewUserToUserEntity(nu))
	if err != nil {
		return err
	}

	return nil
}

// Login the user
// return the access token
func (u *userUseCase) Login(cred models.UserCredential) (t string, e error) {
	// Check username validation
	expectedUser, err := u.db.GetUserByUsername(cred.Username)
	if err != nil {
		return "", err
	}

	// Check credentials
	token, err := authenticate.Login(cred, expectedUser)
	if err != nil {
		return "", err
	}

	t = token.TokenString

	return t, nil
}

// Returns user that has the email address
func (u *userUseCase) GetUserByEmail(s string) (user models.UserInfo, e error) {
	ue, err := u.db.GetUserByEmail(s)
	user = helper.UserEntityToUser(ue)
	if err != nil {
		return models.UserInfo{}, err
	}

	return user, nil
}

// Returns user that has the username
func (u *userUseCase) GetUserByUsername(s string) (user models.UserInfo, e error) {
	ue, err := u.db.GetUserByUsername(s)
	user = helper.UserEntityToUser(ue)
	if err != nil {
		return models.UserInfo{}, err
	}

	return user, nil
}

// Returns a list of all of our users in database
func (u *userUseCase) GetAllUsers() (users []*models.UserInfo, e error) {
	us, err := u.db.FetchAllUsers()
	if err != nil {
		return users, err
	}

	for _, u := range us {
		user := helper.UserEntityToUser(u)
		users = append(users, &user)
	}

	return users, nil
}

func (u *userUseCase) RefreshUserToken(at string) (models.RefreshToken, error) {
	t, err := authenticate.Refresh(at)
	if err != nil {
		return models.RefreshToken{}, err
	}

	return models.RefreshToken{RefreshToken: t.TokenString}, nil
}

// This is for testing purposes
// It puts default user in the database.
func (u *userUseCase) CreateDefaultUser() error {
	err := u.db.AddUser(defaultUser)
	if err != nil {
		log.Printf("%v\n", err)
		return err
	}
	log.Printf("Default user created: %v\n", defaultUser)

	return nil
}
