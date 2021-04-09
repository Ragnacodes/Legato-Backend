package usecase

import (
	"legato_server/api"
	"legato_server/authenticate"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
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
func (u *userUseCase) RegisterNewUser(nu api.NewUser) error {
	err := u.db.AddUser(converter.NewUserToUserDb(nu))
	if err != nil {
		return err
	}

	return nil
}

// Login the user
// return the access token
func (u *userUseCase) Login(cred api.UserCredential) (t string, e error) {
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
func (u *userUseCase) GetUserByEmail(s string) (user api.UserInfo, e error) {
	ue, err := u.db.GetUserByEmail(s)
	user = converter.UserDbToUser(ue)
	if err != nil {
		return api.UserInfo{}, err
	}

	return user, nil
}

// Returns user that has the username
func (u *userUseCase) GetUserByUsername(s string) (user api.UserInfo, e error) {
	ue, err := u.db.GetUserByUsername(s)
	user = converter.UserDbToUser(ue)
	if err != nil {
		return api.UserInfo{}, err
	}

	return user, nil
}

// Returns a list of all of our users in database
func (u *userUseCase) GetAllUsers() (users []*api.UserInfo, e error) {
	us, err := u.db.FetchAllUsers()
	if err != nil {
		return users, err
	}

	for _, u := range us {
		user := converter.UserDbToUser(u)
		users = append(users, &user)
	}

	return users, nil
}

func (u *userUseCase) RefreshUserToken(at string) (api.RefreshToken, error) {
	t, err := authenticate.Refresh(at)
	if err != nil {
		return api.RefreshToken{}, err
	}

	return api.RefreshToken{RefreshToken: t.TokenString}, nil
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
