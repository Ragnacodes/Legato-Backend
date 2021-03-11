package models

type NewUser struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Username  string `json:"username"`
	Password  string `json:"password"`
}

type RefreshToken struct {
	Token string `json:"token"`
}

type User struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Username  string `json:"username"`
	Gender    string `json:"gender"`
}

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
