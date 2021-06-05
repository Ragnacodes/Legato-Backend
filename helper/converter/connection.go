package converter

import (
	"context"
	"encoding/json"
	"legato_server/api"
	"strings"

	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
	"golang.org/x/oauth2/google"
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
		js, _ := json.Marshal(condata.Token)
		token := oauth2.Token{}
		json.Unmarshal([]byte(js), &token)
		data := &map[string]interface{}{
			"token": token,
		}
		return *data, err
	case "gmail":
		type Tokenaouth struct {
			Token map[string]interface{} `json:"token"`
		}
		condata := Tokenaouth{}
		err := json.Unmarshal([]byte(data), &condata)
		js, _ := json.Marshal(condata.Token)
		token := oauth2.Token{}
		json.Unmarshal([]byte(js), &token)
		data := &map[string]interface{}{
			"token": token,
		}
		return *data, err

	}
	return nil, nil
}
func getGitToken(data string) (interface{}, error) {
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
func getGmailToken(data string) (interface{}, error) {
	type extractdata struct {
		Token string `json:"token"`
	}
	var d extractdata
	oauthConf := oauth2.Config{
		ClientID:     "906955768602-u0nu3ruckq6pcjvune1tulkq3n0kfvrl.apps.googleusercontent.com",
		ClientSecret: "VoXRAy3fWRcqFi10rlo31HB2",
		Endpoint:     google.Endpoint,
		RedirectURL:  "http://localhost:3000/redirect/gmail/",
	}

	err := json.Unmarshal([]byte(data), &d)
	token, err := oauthConf.Exchange(context.Background(), d.Token)
	return token, err
}
func ExtractData(data interface{}, Type string, ut *api.Connection) (string, map[string]interface{}, error) {
	switch Type {
	case "github":
		js, _ := json.Marshal(ut.Data)
		token, err := getGitToken(string(js))
		data := &map[string]interface{}{
			"token": token,
		}

		jsonString, err := json.Marshal(token)
		return string(jsonString), *data, err

	case "sshes":
		jsonString, err := json.Marshal(ut.Data)
		return string(jsonString), ut.Data, err
	case "gmail":
		js, _ := json.Marshal(ut.Data)
		token, err := getGmailToken(string(js))
		data := &map[string]interface{}{
			"token": token,
		}

		jsonString, err := json.Marshal(token)
		return string(jsonString), *data, err
	}
	return "", nil, nil
}
