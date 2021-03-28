package router

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"legato_server/models"
	"net/http"
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

func getdatabase() []models.UserCredential {
	type users []models.UserCredential
	user1 := models.UserCredential{
		Username: "reza",
		Password: "1234",
	}
	user2 := models.UserCredential{
		Username: "ali",
		Password: "1234",
	}
	user3 := models.UserCredential{
		Username: "nima",
		Password: "1234",
	}
	user4 := models.UserCredential{
		Username: "mina",
		Password: "12151",
	}
	Users := users{
		user1,
		user2,
		user3,
		user4,
	}
	return Users
}
func isExists(username string, products []models.UserCredential) (result bool) {
	result = false
	for _, product := range products {
		if product.Username == username {
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
		Username: "ali",
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
	Users := users{
		user1,
		user2,
		user3,
		user4,
		user5,
	}
	return Users
}
func TestRegisterUnauthenticated(t *testing.T) {

	// type users []models.UserCredential
	users := getLoginPOSTPayload()
	for i, user := range users {

		if strings.EqualFold(user.Username, " ") == true || strings.EqualFold(user.Password, " ") == true {

			t.Errorf("test number %d failed (empty field) in login", i+1)
			t.Fail()
		} else {
			db := getdatabase()
			if isExists(user.Username, db) {

				payloadBuf := new(bytes.Buffer)
				json.NewEncoder(payloadBuf).Encode(user)
				// fmt.Print()

				// var jsonStr = []byte(`{"username":username,"password" : "1234"}`)
				req, err := http.NewRequest("POST", "http://localhost:8080/api/auth/login", payloadBuf)
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					// panic(err)
					t.Errorf("can't find this url in login")
				}

				fmt.Println("response Headers:", resp.Header)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("response Body:", string(body))
				t.Errorf("test numbetr %d passed in login", i+1)
				defer resp.Body.Close()

			} else {
				t.Errorf("test number %d failed (username does not exit) in login", i+1)
			}

		}
	}

	// type users []models.UserCredential
	users_s := getSignupPOSTPayload()
	for i, user := range users_s {

		if strings.EqualFold(user.Username, " ") == true || strings.EqualFold(user.Password, " ") == true {

			t.Errorf("test number %d failed (empty field) in signup", i+6)
			t.Fail()
		} else {
			db := getdatabase()
			if !isExists(user.Username, db) {

				payloadBuf := new(bytes.Buffer)
				json.NewEncoder(payloadBuf).Encode(user)
				// fmt.Print()

				// var jsonStr = []byte(`{"username":username,"password" : "1234"}`)
				req, err := http.NewRequest("POST", "http://localhost:8080/api/auth/signup", payloadBuf)
				req.Header.Set("Content-Type", "application/json")

				client := &http.Client{}
				resp, err := client.Do(req)
				if err != nil {
					// panic(err)
					t.Errorf("can't find this url in signup")
				}

				fmt.Println("response Headers:", resp.Header)
				body, _ := ioutil.ReadAll(resp.Body)
				fmt.Println("response Body:", string(body))
				t.Errorf("test numbetr %d passed in signup", i+6)
				defer resp.Body.Close()

			} else {
				t.Errorf("test number %d failed (username already exit) in signup", i+6)
			}

		}
	}

}
