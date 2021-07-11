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

var githubUseCase domain.GitUseCase

var creategitNodes []api.NewServiceNode
var updategitNodes []api.NewServiceNode

func createGitNode() {
	// test1
	node1 := api.NewServiceNode{}
	node1.Name = "test1"
	node1.Position = api.Position{X: 30, Y: 30}
	subt := "createPullRequest"
	node1.SubType = &subt
	node1.Type = "githubs"
	var jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"title":"merge to develop",
		"base":"main",
		"head":"develop",
		"owner":"rezamnkh79",
		"body":"develop"
	}`)
	node1.Data = string(jsonStr)
	creategitNodes = append(creategitNodes, node1)
	//test2
	node2 := api.NewServiceNode{}
	node2.Name = "test2"
	node2.Position = api.Position{X: 30, Y: 30}
	subt = "createPullRequest"
	node2.SubType = &subt
	node2.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"title":"merge to develop",
		"base":"main",
		"head":"develop",
		"owner":"rezamnkh79",
		"body":"develop"
	}`)
	node2.Data = string(jsonStr)
	creategitNodes = append(creategitNodes, node2)
	//test3
	node3 := api.NewServiceNode{}
	node3.Name = "test3"
	node3.Position = api.Position{X: 30, Y: 30}
	subt = "createIssue"
	node3.SubType = &subt
	node3.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"title":"merge to develop",
		"base":"master",
		"head":"develop",
		"owner":"rezamnkh79",
		"body":"develop"
	}`)
	node3.Data = string(jsonStr)
	creategitNodes = append(creategitNodes, node3)
	//test4
	node4 := api.NewServiceNode{}
	node4.Name = "test4"
	node3.Position = api.Position{X: 30, Y: 30}
	subt = "createIssue"
	node4.SubType = &subt
	node4.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"body":"hello",
		"title":"ffffff",
		"owner":"rezamnkh79",
		"labels" :["bug","invalid"],
		"assignee" :["rezamnkh79"],
		"state":"open"
	}`)
	node4.Data = string(jsonStr)
	creategitNodes = append(creategitNodes, node4)
	//test5
	node5 := api.NewServiceNode{}
	node5.Name = "test5"
	node5.Position = api.Position{X: 30, Y: 30}
	subt = "createIssue"
	node5.SubType = &subt
	node5.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"body":"hello",
		"title":"ffffff",
		"owner":"rezamnkh79",
		"labels" :["bug","invalid"],
		"assignee" :["rezamnkh79"],
		"state":"open"
	}`)
	node5.Data = string(jsonStr)
	creategitNodes = append(creategitNodes, node5)
	//test6
	node6 := api.NewServiceNode{}
	node6.Name = "test6"
	node6.Position = api.Position{X: 30, Y: 30}
	subt = "createPullRequest"
	node6.SubType = &subt
	node6.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"title":"merge to master",
		"base":"master",
		"head":"develop",
		"owner":"test",
		"body":"develop"
	}`)
	node6.Data = string(jsonStr)
	creategitNodes = append(creategitNodes, node6)
}

func updateGitNode() {
	// test7
	node1 := api.NewServiceNode{}
	node1.Name = "updatetest6"
	node1.Position = api.Position{X: 30, Y: 30}
	subt := "createPullRequest"
	node1.SubType = &subt
	node1.Type = "githubs"
	var jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"title":"merge to develop",
		"base":"main",
		"head":"develop",
		"owner":"rezamnkh79",
		"body":"develop"
	}`)
	node1.Data = string(jsonStr)
	updategitNodes = append(updategitNodes, node1)
	//test8
	node2 := api.NewServiceNode{}
	node2.Name = "updatetest2"
	node2.Position = api.Position{X: 30, Y: 30}
	subt = "createPullRequest"
	node2.SubType = &subt
	node2.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"title":"merge to develop",
		"base":"main",
		"head":"develop",
		"owner":"rezamnkh79",
		"body":"develop"
	}`)
	node2.Data = string(jsonStr)
	updategitNodes = append(updategitNodes, node2)
	//test9
	node3 := api.NewServiceNode{}
	node3.Name = "updatetest3"
	node3.Position = api.Position{X: 30, Y: 30}
	subt = "createIssue"
	node3.SubType = &subt
	node3.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"title":"merge to develop",
		"base":"master",
		"head":"develop",
		"owner":"rezamnkh79",
		"body":"develop"
	}`)
	node3.Data = string(jsonStr)
	updategitNodes = append(updategitNodes, node3)
	//test10
	node4 := api.NewServiceNode{}
	node4.Name = "updatetest4"
	node3.Position = api.Position{X: 30, Y: 30}
	subt = "createIssue"
	node4.SubType = &subt
	node4.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"body":"hello",
		"title":"ffffff",
		"owner":"rezamnkh79",
		"labels" :["bug","invalid"],
		"assignee" :["rezamnkh79"],
		"state":"open"
	}`)
	node4.Data = string(jsonStr)
	updategitNodes = append(updategitNodes, node4)
	//test11
	node5 := api.NewServiceNode{}
	node5.Name = "updatetest5"
	node5.Position = api.Position{X: 30, Y: 30}
	subt = "createIssue"
	node5.SubType = &subt
	node5.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"body":"hello",
		"title":"ffffff",
		"owner":"rezamnkh79",
		"labels" :["bug","invalid"],
		"assignee" :["rezamnkh79"],
		"state":"open"
	}`)
	node5.Data = string(jsonStr)
	updategitNodes = append(updategitNodes, node5)
	//test12
	node6 := api.NewServiceNode{}
	node6.Name = "updatetest6"
	node6.Position = api.Position{X: 30, Y: 30}
	subt = "createPullRequest"
	node6.SubType = &subt
	node6.Type = "githubs"
	jsonStr = []byte(`{
		"repositoryName":"eeeee",
		"title":"merge to master",
		"base":"master",
		"head":"develop",
		"owner":"test",
		"body":"develop"
	}`)
	node6.Data = string(jsonStr)
	updategitNodes = append(updategitNodes, node6)
}

type GitUse struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func TestGithub(t *testing.T) {

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
	githubUseCase = NewGithubUseCase(appDB, timeoutContext)
	userUseCase := userUC.NewUserUseCase(appDB, timeoutContext)
	user, _ := userUseCase.GetUserByUsername("legato")
	fmt.Println(user.Username)
	scenario := api.NewScenario{
		Name: "myscenario",
	}
	scenarioUseCase := scenarioUC.NewScenarioUseCase(appDB, timeoutContext)
	var x = true
	scenario.IsActive = &x
	// g.db.AddScenario()
	s, err := scenarioUseCase.AddScenario(&user, &scenario)
	createGitNode()
	updateGitNode()
	var nodeID []uint
	for _, node := range creategitNodes {
		s, _ := githubUseCase.AddToScenario(&user, s.ID, node)
		nodeID = append(nodeID, s.Id)
	}
	for i, node := range updategitNodes {
		githubUseCase.Update(&user, s.ID, nodeID[i], node)
	}

}
