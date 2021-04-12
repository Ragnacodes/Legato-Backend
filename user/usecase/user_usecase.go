package usecase

import (
	"fmt"
	"legato_server/authenticate"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"legato_server/models"
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
	err := u.db.AddUser(converter.NewUserToUserDb(nu))
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
	user = converter.UserDbToUser(ue)
	if err != nil {
		return models.UserInfo{}, err
	}

	return user, nil
}

// Returns user that has the username
func (u *userUseCase) GetUserByUsername(s string) (user models.UserInfo, e error) {
	ue, err := u.db.GetUserByUsername(s)
	user = converter.UserDbToUser(ue)
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
		user := converter.UserDbToUser(u)
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

func (u *userUseCase) AddTokenDB(name string, ut models.UserAddToken) error {
	user, _ := u.db.GetUserByUsername(name)
	con := converter.NewTokenDb(ut)
	err := u.db.AddToken(&user, con)
	if err != nil {
		return fmt.Errorf("can not add token")
	}

	return nil
}

func (u *userUseCase) GetTokenByUsername(username string, ut models.UserGetToken) (legatoDb.Connection, error) {
	user, _ := u.db.GetUserByUsername(username)
	connection, err := u.db.GetUserToken(&user, ut)
	if err != nil {
		return legatoDb.Connection{}, fmt.Errorf("can not find token")
	}

	return connection, nil
}

func (u *userUseCase) GetTokensByUsername(username string, ut models.UserGetToken) ([]legatoDb.Connection, error) {
	user, _ := u.db.GetUserByUsername(username)
	connections, err := u.db.GetUserTokens(&user, ut)
	if err != nil {
		return []legatoDb.Connection{}, fmt.Errorf("can not find token")
	}

	return connections, nil
}

func (u *userUseCase) UpdateUserTokenById(username string, ut models.UserGetToken) error {
	user, _ := u.db.GetUserByUsername(username)

	err := u.db.UpdateUserTokenById(&user, ut.Name, ut, ut.Token_id)

	if err != nil {
		return err
	}

	return nil
}

func (u *userUseCase) CheckTokenByID(username string, ut models.UserGetToken) error {
	user, _ := u.db.GetUserByUsername(username)
	err := u.db.CheckTokenByID(&user, ut.Token_id)
	if err != nil {
		return nil
	}
	return err
}

func (u *userUseCase) DeleteUserTokenById(username string, ut models.UserGetToken) error {
	user, _ := u.db.GetUserByUsername(username)

	err := u.db.DeleteConnectionByID(&user, ut, ut.Token_id)

	if err != nil {
		return err
	}

	return nil
}
func (u *userUseCase) UpdateUserTokenByName(username string, ut models.UserGetToken) error {
	user, _ := u.db.GetUserByUsername(username)

	err := u.db.UpdateUserTokenByName(&user, ut.Name, ut, ut.Token_id)

	if err != nil {
		return err
	}

	return nil
}
