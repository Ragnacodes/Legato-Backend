package models

type NewUser struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type RefreshToken struct {
	RefreshToken string `json:"refresh_token"`
}

type UserInfo struct {
	Email    string `json:"email"`
	Username string `json:"username"`
}

type UserCredential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserAddToken struct {
	Token      string `json:"token"`
	Token_type string `json:"token_type"`
	UserID     uint   `json:"userid"`
	Name       string `json:"name"`
}
type UserGetToken struct {
	Token_type string `json:"token_type"`
	UserID     uint   `json:"userid"`
	Name       string `json:"name"`
	Token_id   uint   `json:"id"`
}
