package usecase

import (
	"legato_server/api"
	"legato_server/authenticate"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"testing"
	"time"

	"github.com/spf13/viper"
)

var connectionUseCase domain.GmailUseCase

var createConnections []api.Connection
var updateConnections []api.Connection

func createConnection() {
	// test1
	conn1 := api.Connection{}
	conn1.Name = "git"
	conn1.Type = "githubs"
	jsonStr := map[string]interface{}{
		"token": "mfxl'vmsdv';mfdb'fdamlmdsvfdkfnjn",
	}
	conn1.Data = jsonStr
	createConnections = append(createConnections, conn1)
	// test2
	conn2 := api.Connection{}
	conn2.Name = "git"
	conn2.Type = "githubs"
	jsonStr = map[string]interface{}{
		"token": "maaaaaaaaaaaaadddddddddddddddweeeeeee",
	}
	conn1.Data = jsonStr
	createConnections = append(createConnections, conn2)
	// test3
	conn3 := api.Connection{}
	conn3.Name = "ssh1"
	conn3.Type = "sshes"
	jsonStr = map[string]interface{}{
		"host":     "37.152.181.64",
		"username": "reza",
		"password": "--------------------",
		"commands": []string{"ls", "echo hello world"},
	}
	conn3.Data = jsonStr
	createConnections = append(createConnections, conn3)
	// test4
	conn4 := api.Connection{}
	conn4.Name = "ssh2"
	conn4.Type = "sshes"
	jsonStr = map[string]interface{}{
		"host":     "37.152.181.64",
		"username": "reza",
		"password": "------------------------",
		"commands": []string{"ls"},
	}
	conn4.Data = jsonStr
	createConnections = append(createConnections, conn4)
	// test5
	conn5 := api.Connection{}
	conn5.Name = "gmail"
	conn5.Type = "gmails"
	jsonStr = map[string]interface{}{
		"host":     "37.152.181.64",
		"username": "reza",
		"password": "sko192j3h",
		"commands": []string{"ls", "echo hello world"},
	}
	conn5.Data = jsonStr
	createConnections = append(createConnections, conn5)

}

func updateConnection() {
	// test6
	conn1 := api.Connection{}
	conn1.Name = "update_git"
	conn1.Type = "githubs"
	jsonStr := map[string]interface{}{
		"token": "mfxdddddddddfeeeeeeeeeeeee",
	}
	conn1.Data = jsonStr
	updateConnections = append(updateConnections, conn1)
	// test2
	conn2 := api.Connection{}
	conn2.Name = "update_git"
	conn2.Type = "githubs"
	jsonStr = map[string]interface{}{
		"token": "updateeeeeeeeeee",
	}
	conn1.Data = jsonStr
	updateConnections = append(updateConnections, conn1)
	// test3
	conn3 := api.Connection{}
	conn3.Name = "update_ssh"
	conn3.Type = "sshes"
	jsonStr = map[string]interface{}{
		"host":     "37.152.181.64",
		"username": "reza",
		"password": "sko192j3h",
		"commands": []string{"ls", "echo hello world"},
	}
	conn3.Data = jsonStr
	updateConnections = append(updateConnections, conn3)
	// test4
	conn4 := api.Connection{}
	conn4.Name = "update_ssh2"
	conn4.Type = "sshes"
	jsonStr = map[string]interface{}{
		"host":     "37.152.181.64",
		"username": "reza",
		"password": "------------------------",
		"commands": []string{"ls", "cd home"},
	}
	conn4.Data = jsonStr
	updateConnections = append(updateConnections, conn4)
	// test5
	conn5 := api.Connection{}
	conn5.Name = "update_gmail"
	conn5.Type = "gmails"
	jsonStr = map[string]interface{}{
		"host":     "37.152.181.64",
		"username": "reza",
		"password": "sko192j3h",
		"commands": []string{"ls", "echo update"},
	}
	conn5.Data = jsonStr
	updateConnections = append(updateConnections, conn5)

}
func TestConnection(t *testing.T) {

	env.LoadEnv()

	// Generate random jwt key
	authenticate.GenerateRandomKey()

	// Make server sent event

	// Connect to database
	appDB, err := legatoDb.Connect()
	if err != nil {
		panic(err)
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	userUseCase := NewUserUseCase(appDB, timeoutContext)
	_ = userUseCase.CreateDefaultUser()
	user, _ := userUseCase.GetUserByUsername("legato")
	createConnection()
	updateConnection()
	var nodeID []uint
	for _, con := range createConnections {
		s, _ := userUseCase.AddConnectionToDB(user.Username, con)
		nodeID = append(nodeID, s.ID)
	}
	for i, con := range updateConnections {
		con.ID = nodeID[i]
		userUseCase.UpdateDataConnectionByID(user.Username, con)
	}
	userUseCase.DeleteUserConnectionById(user.Username, nodeID[0])
	userUseCase.DeleteUserConnectionById(user.Username, nodeID[1])

}
