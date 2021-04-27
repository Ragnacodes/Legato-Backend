package api

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type UserInfo struct {
	ID       uint
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
