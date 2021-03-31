package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"legato_server/models"
	"net/http"
	"regexp"
	"strings"
	"testing"
)

func getLoginPOSTPayload() []models.UserCredential {
	type users []models.UserCredential
	user1 := models.UserCredential{
		Username: "reza",
		Password: "1234",
	}
	user2 := models.UserCredential{
		Username: " ",
		Password: "1234",
	}
	user3 := models.UserCredential{
		Username: "reza",
		Password: " ",
	}
	user4 := models.UserCredential{
		Username: " ",
		Password: " ",
	}

	user5 := models.UserCredential{
		Username: "mahdi",
		Password: "12356",
	}
	Users := users{
		user1,
		user2,
		user3,
		user4,
		user5,
	}
	return Users
}

func getdatabase() []models.NewUser {
	type users []models.NewUser
	user1 := models.NewUser{
		Username: "reza",
		Password: "1234",
		Email:    "reza@example.com",
	}
	user2 := models.NewUser{
		Username: "ali",
		Password: "1234",
		Email:    "ali@example.com",
	}
	user3 := models.NewUser{
		Username: "nima",
		Password: "1234",
		Email:    "al@examp.com",
	}
	user4 := models.NewUser{
		Username: "mina",
		Password: "12151",
		Email:    "aliexample.com",
	}
	user5 := models.NewUser{
		Username: "mahsa",
		Password: "12151",
		Email:    "mahsa@example.com",
	}

	Users := users{
		user1,
		user2,
		user3,
		user4,
		user5,
	}
	return Users
}

// check if a username is in a struct return True
func isExists_username(username string, products []models.NewUser) (result bool) {

	result = false
	for _, product := range products {
		if strings.EqualFold(product.Username, username) {
			result = true
			break
		}
	}
	return result
}

// check if a email is in a struct return True
func isExists_email(email string, products []models.NewUser) (result bool) {

	result = false
	for _, product := range products {
		if strings.EqualFold(product.Email, email) {
			result = true
			break
		}
	}
	return result
}

func getSignupPOSTPayload() []models.NewUser {
	type users []models.NewUser
	user1 := models.NewUser{
		Username: "reza",
		Email:    "reza@example.com",
		Password: "1234",
	}
	user2 := models.NewUser{
		Username: " ",
		Email:    "eee@example.com",
		Password: "1234",
	}
	user3 := models.NewUser{
		Username: "aliaa",
		Email:    "ali@example.com",
		Password: "1234",
	}
	user4 := models.NewUser{
		Username: "milad",
		Email:    "reza@example.com",
		Password: "1336544",
	}

	user5 := models.NewUser{
		Username: "danial",
		Email:    "mahdi@example.com",
		Password: "12356",
	}
	user6 := models.NewUser{
		Username: "mahdi",
		Email:    "mahdiexample.com",
		Password: "12356",
	}
	user7 := models.NewUser{
		Username: "melika",
		Email:    "melika@example.com",
		Password: "12356",
	}
	user8 := models.NewUser{
		Username: "melika",
		Email:    " ",
		Password: " ",
	}
	user9 := models.NewUser{
		Username: " ",
		Email:    " ",
		Password: " ",
	}
	Users := users{
		user1,
		user2,
		user3,
		user4,
		user5,
		user6,
		user7,
		user8,
		user9,
	}
	return Users
}

func sendrequest(req *http.Request, err error, i int, t *testing.T) {
	client := &http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		t.Errorf("test number %d can't find this url in login failed test ", i+1)
	} else {
		t.Errorf("test numbetr %d passed edwsin login", i+1)

	}

	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	defer resp.Body.Close()
}

func signup_check(user models.NewUser, t *testing.T, i int) {
	payloadBuf := new(bytes.Buffer)
	json.NewEncoder(payloadBuf).Encode(user)
	req, err := http.NewRequest("POST", "http://localhost:8080/api/auth/signup", payloadBuf)
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		t.Errorf("can't find this url in signup")
	}

	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))
	t.Errorf("test numbetr %d passed in signup", i+6)
	defer resp.Body.Close()
}

func TestRegisterUnauthenticated(t *testing.T) {

	users := getLoginPOSTPayload()
	for i, user := range users {
		fmt.Printf(" number : %d", i)
		if strings.EqualFold(user.Username, " ") == true || strings.EqualFold(user.Password, " ") == true {

			t.Errorf("test number %d failed (empty field) in login", i+1)
			t.Fail()
		} else {
			db := getdatabase()
			if isExists_username(user.Username, db) {

				payloadBuf := new(bytes.Buffer)
				json.NewEncoder(payloadBuf).Encode(user)
				url1 := "http://localhost:8080/api/auth/login"
				url2 := "http://localhost:8080/api/auth/loging"
				req1, err1 := http.NewRequest("POST", url1, payloadBuf)
				req1.Header.Set("Content-Type", "application/json")
				req2, err2 := http.NewRequest("POST", url2, payloadBuf)
				req2.Header.Set("Content-Type", "application/json")
				sendrequest(req2, err2, i-1, t)
				sendrequest(req1, err1, i, t)

			} else {
				t.Errorf("test number %d failed (username does not exit) in login", i+1)
			}

		}
	}

	users_s := getSignupPOSTPayload()
	for i, user := range users_s {
		// i++
		if strings.EqualFold(user.Username, " ") == true || strings.EqualFold(user.Password, " ") == true || strings.EqualFold(user.Email, " ") == true {

			t.Errorf("test number %d failed (empty field) in signup", i+6)
			t.Fail()
		} else {
			db := getdatabase()
			// if username already exist in database test will fail
			if !isExists_username(user.Username, db) && !isExists_email(user.Email, db) {
				pattern := regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
				matches := pattern.MatchString(user.Email)
				if matches == true {
					signup_check(user, t, i)

				} else {
					t.Errorf("test number %d failed (format of email is incorrect", i+6)
				}
			} else {
				if isExists_email(user.Email, db) {
					t.Errorf("test number %d failed (email already exit) in signup", i+6)
				} else {
					t.Errorf("test number %d failed (username already exit) in signup", i+6)
				}
			}

		}
	}

}
