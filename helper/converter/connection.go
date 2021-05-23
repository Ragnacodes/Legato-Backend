package converter

import (
	"encoding/json"
	"legato_server/api"
	"strings"
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
	condata := api.Connection{}
	err := json.Unmarshal([]byte(data), &condata)
	switch Type {
	case "sshes":
		flag := false
		if strings.Contains(string(data), "password") == true {
			flag = true
		}

		if flag == true {
			data := map[string]interface{}{"host": condata.Data["host"], "password": condata.Data["password"], "username": condata.Data["username"]}
			return data, err
		} else {
			data := map[string]interface{}{"host": condata.Data["host"], "sshKey": condata.Data["sshKey"], "username": condata.Data["username"]}
			return data, err
		}

	}
	return nil, nil
}
