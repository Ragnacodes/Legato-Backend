package converter

import (
	"encoding/json"
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

	}
	return nil, nil
}
