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
	FirstName: "Admin",
	LastName:  "Admin",
	Username:  "admin",
	Email:     "legato@gmail.com",
	Password:  "123qwe",
}

type userUseCase struct {
	//userRepo       domain.UserRepository
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
	expectedUser, loginErr := u.db.GetUserByUsername(cred.Username)
	if loginErr != nil {
		return "", loginErr
	}

	// Check credentials
	token, authErr := authenticate.Login(cred, expectedUser)
	if authErr != nil {
		return "", authErr
	}

	t = token.TokenString

	return t, nil
}

// Returns user that has the email address
func (u *userUseCase) GetUserByEmail(s string) (user models.User, e error) {
	ue, err := u.db.GetUserByEmail(s)
	user = helper.UserEntityToUser(ue)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Returns user that has the username
func (u *userUseCase) GetUserByUsername(s string) (user models.User, e error) {
	ue, err := u.db.GetUserByUsername(s)
	user = helper.UserEntityToUser(ue)
	if err != nil {
		return models.User{}, err
	}

	return user, nil
}

// Returns a list of all of our users in database
func (u *userUseCase) GetAllUsers() (users []*models.User, e error) {
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

func (u *userUseCase) RefreshUserToken(rft models.RefreshToken) (string, error) {
	t, err := authenticate.Refresh(rft.Token)
	if err != nil {
		return "", err
	}

	return t.TokenString, nil
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
