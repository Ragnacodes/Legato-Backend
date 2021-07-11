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

var gmailUseCase domain.GmailUseCase

var createGmialNodes []api.NewServiceNode
var updateGmailNodes []api.NewServiceNode

func createGmailNode() {
	// test1
	node1 := api.NewServiceNode{}
	node1.Name = "test1"
	node1.Position = api.Position{X: 30, Y: 30}
	subt := "sendEmail"
	node1.SubType = &subt
	node1.Type = "gmials"
	var jsonStr = []byte(`{
		"body":"hello",
		"subject":"test",
		"to":["mansourikhahreza@gmail.com"],
		"email":"rezamansourikhah@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node1.Data = string(jsonStr)
	createGmialNodes = append(createGmialNodes, node1)
	// test2
	node2 := api.NewServiceNode{}
	node2.Name = "test2"
	node2.Position = api.Position{X: 30, Y: 30}
	subt = "sendEmail"
	node2.SubType = &subt
	node2.Type = "gmials"
	jsonStr = []byte(`{
		"body":"hello world",
		"subject":"test email",
		"to":["mansourikhah@gmail.com"],
		"email":"reza@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node2.Data = string(jsonStr)
	createGmialNodes = append(createGmialNodes, node2)
	// test3
	node3 := api.NewServiceNode{}
	node3.Name = "test3"
	node3.Position = api.Position{X: 30, Y: 30}
	subt = "sendEmail"
	node3.SubType = &subt
	node3.Type = "gmials"
	jsonStr = []byte(`{
		"body":"hello world",
		"subject":"test email",
		"to":["example@gmail.com"],
		"email":"reza@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node3.Data = string(jsonStr)
	createGmialNodes = append(createGmialNodes, node3)
	// test4
	node4 := api.NewServiceNode{}
	node4.Name = "test4"
	node4.Position = api.Position{X: 30, Y: 30}
	subt = "sendEmail"
	node4.SubType = &subt
	node4.Type = "gmials"
	jsonStr = []byte(`{
		"body":"hello world",
		"subject":"test email",
		"to":["example@gmail.com"],
		"email":"reza@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node4.Data = string(jsonStr)
	createGmialNodes = append(createGmialNodes, node4)
	// test5
	node5 := api.NewServiceNode{}
	node5.Name = "test4"
	node5.Position = api.Position{X: 30, Y: 30}
	subt = "sendEmail"
	node5.SubType = &subt
	node5.Type = "gmials"
	jsonStr = []byte(`{
		"body":"hello world",
		"subject":"test email",
		"to":["example@gmail.com"],
		"email":"reza@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node5.Data = string(jsonStr)
	createGmialNodes = append(createGmialNodes, node5)

}

func updateGmailNode() {
	// test6
	node1 := api.NewServiceNode{}
	node1.Name = "updatetest6"
	node1.Position = api.Position{X: 30, Y: 30}
	subt := "sendEmail"
	node1.SubType = &subt
	node1.Type = "gmials"
	var jsonStr = []byte(`{
		"body":"hello",
		"subject":"test",
		"to":["mansourikhahreza@gmail.com"],
		"email":"rezamansourikhah@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node1.Data = string(jsonStr)
	updateGmailNodes = append(updateGmailNodes, node1)
	// test7
	node2 := api.NewServiceNode{}
	node2.Name = "updatetest7"
	node2.Position = api.Position{X: 30, Y: 30}
	subt = "sendEmail"
	node2.SubType = &subt
	node2.Type = "gmials"
	jsonStr = []byte(`{
		"body":"hello world",
		"subject":"test email",
		"to":["mansourikhah@gmail.com"],
		"email":"reza@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node2.Data = string(jsonStr)
	updateGmailNodes = append(updateGmailNodes, node2)
	// test8
	node3 := api.NewServiceNode{}
	node3.Name = "updatetest8"
	node3.Position = api.Position{X: 30, Y: 30}
	subt = "sendEmail"
	node3.SubType = &subt
	node3.Type = "gmials"
	jsonStr = []byte(`{
		"body":"hello this is update",
		"subject":"test email",
		"to":["example@gmail.com"],
		"email":"reza@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node3.Data = string(jsonStr)
	updateGmailNodes = append(updateGmailNodes, node3)
	// test9
	node4 := api.NewServiceNode{}
	node4.Name = "update test9"
	node4.Position = api.Position{X: 30, Y: 30}
	subt = "sendEmail"
	node4.SubType = &subt
	node4.Type = "gmials"
	jsonStr = []byte(`{
		"body":"hello this is update",
		"subject":"test email",
		"to":["example@gmail.com"],
		"email":"reza@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node4.Data = string(jsonStr)
	updateGmailNodes = append(updateGmailNodes, node4)
	// test10
	node5 := api.NewServiceNode{}
	node5.Name = "test10"
	node5.Position = api.Position{X: 30, Y: 30}
	subt = "sendEmail"
	node5.SubType = &subt
	node5.Type = "gmials"
	jsonStr = []byte(`{
		"body":"hello this is update",
		"subject":"test email",
		"to":["example@gmail.com"],
		"email":"reza@gmail.com",
		"password":"XXXXXXXXXXXXXXXX"
	}`)
	node5.Data = string(jsonStr)
	updateGmailNodes = append(updateGmailNodes, node5)

}
func TestGmail(t *testing.T) {

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
	gmailUseCase = NewGmailUseCase(appDB, timeoutContext)
	userUseCase := userUC.NewUserUseCase(appDB, timeoutContext)
	scenarioUseCase := scenarioUC.NewScenarioUseCase(appDB, timeoutContext)
	_ = userUseCase.CreateDefaultUser()
	user, _ := userUseCase.GetUserByUsername("legato")
	fmt.Println(user.Username)
	scenario := api.NewScenario{
		Name: "myscenario",
	}
	var x = true
	scenario.IsActive = &x
	s, err := scenarioUseCase.AddScenario(&user, &scenario)
	createGmailNode()
	updateGmailNode()
	var nodeID []uint
	for _, node := range createGmialNodes {
		s, _ := gmailUseCase.AddToScenario(&user, s.ID, node)
		nodeID = append(nodeID, s.Id)
	}
	for i, node := range updateGmailNodes {
		gmailUseCase.Update(&user, s.ID, nodeID[i], node)
	}

}
