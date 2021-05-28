package converter

import (
	"context"
	"encoding/json"
	"fmt"
	"legato_server/api"
	"strings"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

type loginWithPasswordData struct {
	Commands []string `json:"commands"`
	Host     string   `json:"host"`
	Password string   `json:"password"`
	Username string   `json:"username"`
}
type loginWithSshKeyData struct {
	Username string   `json:"username"`
	Host     string   `json:"host"`
	SshKey   string   `json:"sshKey"`
	Commands []string `json:"commands"`
}

func BindConnectionData(data string, Type string) (map[string]interface{}, error) {

	switch Type {
	case "sshes":
		flag := false
		if strings.Contains(string(data), "password") == true {
			flag = true
		}

		if flag == true {
			condata := loginWithPasswordData{}
			err := json.Unmarshal([]byte(data), &condata)
			data := map[string]interface{}{"host": condata.Host, "password": condata.Password, "username": condata.Username}
			return data, err
		} else {
			condata := loginWithSshKeyData{}
			err := json.Unmarshal([]byte(data), &condata)
			data := map[string]interface{}{"host": condata.Host, "sshKey": condata.SshKey, "username": condata.Username}
			return data, err
		}
	case "github":
		type Tokenaouth struct {
			Token map[string]interface{} `json:"token"`
		}
		condata := Tokenaouth{}
		err := json.Unmarshal([]byte(data), &condata)
		fmt.Print(condata.Token["access_token"])
		data := map[string]interface{}{"access_token": condata.Token["access_token"], "token_type": condata.Token["token_type"], "expiry": condata.Token["expiry"]}
		return data, err

	}
	return nil, nil
}
func getToken(data string) (interface{}, error) {
	type extractdata struct {
		Token string `json:"token"`
	}
	var d extractdata
	oauthConf := &oauth2.Config{
		ClientID:     "a87b311ff0542babc5bd",
		ClientSecret: "0d24ae8ec82ca068984ee25e0b6285be9e9c211c",
		Scopes:       []string{"user:email", "repo", "public_repo", "repo_deployment", "write:org", "delete_repo", "read:org"},
		Endpoint:     githuboauth.Endpoint,
	}
	err := json.Unmarshal([]byte(data), &d)
	token, err := oauthConf.Exchange(context.Background(), d.Token)
	return token, err
}

func ExtractData(data interface{}, Type string, ut *api.Connection) (string, error) {
	switch Type {
	case "github":
		s, _ := json.Marshal(ut.Data)
		token, err := getToken(string(s))
		data := &map[string]interface{}{
			"token": token,
		}
		jsonString, err := json.Marshal(data)
		return string(jsonString), err

	case "sshes":
		jsonString, err := json.Marshal(ut.Data)
		return string(jsonString), err
	}
	return "", nil
}
