package usecase

import (
	"fmt"
	"legato_server/api"
	"legato_server/authenticate"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	scenarioUC "legato_server/scenario/usecase"
	userUC "legato_server/user/usecase"
	"testing"
	"time"

	"github.com/spf13/viper"
)

type Position struct {
	X int `json:"x"`
	Y int `json:"y"`
}

var sshUseCase domain.SshUseCase
var userUseCase domain.UserUseCase
var scenarioUseCase domain.ScenarioUseCase

var createSsheNodes []api.NewServiceNode
var updateSsheNodes []api.NewServiceNode

func createSshNode() {
	// test1
	node1 := api.NewServiceNode{}
	node1.Name = "test1"
	node1.Position = api.Position{X: 30, Y: 30}
	subt := "ssh"
	node1.SubType = &subt
	node1.Type = "sshes"
	var jsonStr = []byte(`{
		"host" :"37.152.181.64",
		"username":"reza",
		"password":"-----------",
		"commands":["ls","echo hello world"]
	  }`)
	node1.Data = string(jsonStr)
	createSsheNodes = append(createSsheNodes, node1)
	//test2
	node2 := api.NewServiceNode{}
	node2.Name = "test2"
	node2.Position = api.Position{X: 30, Y: 30}
	subt = "ssh"
	node2.SubType = &subt
	node2.Type = "sshes"
	var jsonStr2 = []byte(`{
		"host" :"37.152.181.64",
		"username":"amin",
		"password":"----------",
		"commands":["ls","echo hello world"]
	  }`)
	node2.Data = string(jsonStr2)
	createSsheNodes = append(createSsheNodes, node2)
	//test3
	node3 := api.NewServiceNode{}
	node3.Name = "test3"
	node3.Position = api.Position{X: 30, Y: 30}
	subt = "ssh"
	node3.SubType = &subt
	node3.Type = "sshes"
	var jsonStr3 = []byte(`{
		"host" :"37.152.181.64",
		"username":"masoud",
		"password":"----------",
		"commands":["ls","echo hello"]
	  }`)
	node3.Data = string(jsonStr3)
	createSsheNodes = append(createSsheNodes, node3)
	//test4
	node4 := api.NewServiceNode{}
	node4.Name = "test4"
	node4.Position = api.Position{X: 30, Y: 30}
	subt = "ssh"
	node4.SubType = &subt
	node4.Type = "sshes"
	var jsonStr4 = []byte(`{
		"host" :"37.152.181.64",
		"username":"ali",
		"password":"----------",
		"commands":["ls","echo hello"]
	  }`)
	node4.Data = string(jsonStr4)
	createSsheNodes = append(createSsheNodes, node4)
}

func updateSshNode() {
	// test5
	node1 := api.NewServiceNode{}
	node1.Name = "updatetest1"
	node1.Position = api.Position{X: 20, Y: 30}
	var jsonStr = []byte(`{
		"host" :"37.152.181.64",
		"username":"reza",
		"password":"-----------",
		"commands":["ls","echo hello world"]
	  }`)
	node1.Data = string(jsonStr)
	updateSsheNodes = append(updateSsheNodes, node1)
	//test6
	node2 := api.NewServiceNode{}
	node2.Name = "updatetest2"
	node2.Position = api.Position{X: 20, Y: 30}
	var jsonStr2 = []byte(`{
		"host" :"37.152.181.64",
		"username":"amin",
		"password":"----------",
		"commands":["ls","echo hello world"]
	  }`)
	node2.Data = string(jsonStr2)
	updateSsheNodes = append(updateSsheNodes, node2)
	//test7
	node3 := api.NewServiceNode{}
	node3.Name = "updatetest3"
	node3.Position = api.Position{X: 20, Y: 30}
	var jsonStr3 = []byte(`{
		"host" :"37.152.181.64",
		"username":"masoud",
		"password":"----------",
		"commands":["ls","echo hello"]
	  }`)
	node3.Data = string(jsonStr3)
	updateSsheNodes = append(updateSsheNodes, node3)
	//test8
	node4 := api.NewServiceNode{}
	node4.Name = "test4"
	node4.Position = api.Position{X: 30, Y: 30}
	var jsonStr4 = []byte(`{
		"host" :"37.152.181.64",
		"username":"ali",
		"password":"----------",
		"commands":["ls","echo hello"]
	  }`)
	node4.Data = string(jsonStr4)
	updateSsheNodes = append(updateSsheNodes, node4)
}
func TestSSH(t *testing.T) {

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
	sshUseCase = NewSshUseCase(appDB, timeoutContext)
	userUseCase = userUC.NewUserUseCase(appDB, timeoutContext)
	scenarioUseCase = scenarioUC.NewScenarioUseCase(appDB, timeoutContext)
	_ = userUseCase.CreateDefaultUser()
	user, _ := userUseCase.GetUserByUsername("legato")
	fmt.Println(user.Username)
	scenario := api.NewScenario{
		Name: "myscenario",
	}
	var x = true
	scenario.IsActive = &x
	s, err := scenarioUseCase.AddScenario(&user, &scenario)
	createSshNode()
	updateSshNode()
	var nodeID []uint
	for _, node := range createSsheNodes {
		s, _ := sshUseCase.AddToScenario(&user, s.ID, node)
		nodeID = append(nodeID, s.Id)
	}

	for i, node := range updateSsheNodes {
		sshUseCase.Update(&user, s.ID, nodeID[i], node)
	}

}
