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
	case "discords":
		type DiscordData struct {
			GuildId string `json:"guildId"`
		}
		condata := DiscordData{}
		err := json.Unmarshal([]byte(data), &condata)
		json.Unmarshal([]byte(data), &condata)
		data := map[string]interface{}{"guildId": condata.GuildId}
		return data, err
	case "telegrams":
		type TelegramData struct {
			GuildId string `json:"key"`
		}
		condata := TelegramData{}
		err := json.Unmarshal([]byte(data), &condata)
		json.Unmarshal([]byte(data), &condata)
		data := map[string]interface{}{"key": condata.GuildId}
		return data, err

	}
	return nil, nil
}
func ExtractData(data interface{}, Type string, ut *api.Connection) (string, map[string]interface{}, error) {
	switch Type {
	case "discords":
		jsonString, err := json.Marshal(ut.Data)
		return string(jsonString), ut.Data, err

	case "telegrams":
		jsonString, err := json.Marshal(ut.Data)
		return string(jsonString), ut.Data, err

	case "sshes":
		jsonString, err := json.Marshal(ut.Data)
		return string(jsonString), ut.Data, err

	}
	return "", nil, nil
}
