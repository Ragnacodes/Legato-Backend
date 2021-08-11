package authenticate

import (
	"crypto/rand"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"legato_server/api"
	legatoDb "legato_server/db"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Token
type Token struct {
	TokenString string
	Expiration  time.Time
}

// JWT SECRET KEY
var jwtKey []byte

// Each time that server is initialized, GenerateRandomKey is called to
// generate another key
func GenerateRandomKey() []byte {
	jwtKey = make([]byte, 32)
	if _, err := rand.Read(jwtKey); err != nil {
		panic(err)
	}
	fmt.Printf("JWT Key: %b \n", jwtKey)

	return jwtKey
}

// Login check input details with database.
// If everything was ok then it creates JWT token.
// Returns JWT token
func Login(cred api.UserCredential, user legatoDb.User) (t Token, e error) {

	// Check Password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(cred.Password))
	if err != nil {
		return Token{}, errors.New("wrong password")
	}

	//expirationTime := time.Now().Add(30 * time.Minute)
	expirationTime := time.Now().Add(30 * 24 * time.Hour)

	claims := &Claims{
		Username: cred.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	tokenJWT := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := tokenJWT.SignedString(jwtKey)
	if err != nil {
		return Token{}, errors.New("internal server error")
	}

	t = Token{
		TokenString: tokenString,
		Expiration:  expirationTime,
	}

	return t, nil
}

func Refresh(tokenStr string) (t Token, e error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return t, errors.New("invalid token")
		}
		return t, errors.New("bad request")
	}

	if !tkn.Valid {
		return t, errors.New("internal server error")
	}

	// TODO: discuss @masoud about setting this time
	// Ensure that a new token is not issued until enough time has elapsed
	//log.Println(time.Unix(claims.ExpiresAt, 0).Sub(time.Now()))
	//if time.Unix(claims.ExpiresAt, 0).Sub(time.Now()) > 30*time.Hour {
	//	return t, errors.New("bad request time")
	//}

	expirationTime := time.Now().Add(30 * time.Minute)
	claims.ExpiresAt = expirationTime.Unix()
	newToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := newToken.SignedString(jwtKey)
	if err != nil {
		return t, errors.New("internal server error")
	}

	t = Token{
		TokenString: tokenString,
		Expiration:  expirationTime,
	}

	return t, nil
}

// CheckToken check validation for incoming tokens
func CheckToken(tokenStr string) (*Claims, error) {
	claims := &Claims{}

	tkn, err := jwt.ParseWithClaims(tokenStr, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			return nil, errors.New("invalid token")
		}
		return nil, errors.New("bad request")
	}

	if !tkn.Valid {
		return nil, errors.New("unauthorized")
	}

	return claims, nil
}
