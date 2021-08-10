package legatoDb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"legato_server/services"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githubOAuth "golang.org/x/oauth2/github"
	"gorm.io/gorm"
)

const gitType string = "githubs"

type Github struct {
	gorm.Model
	ConnectionID uint
	Connection   *Connection `gorm:"foreignkey:id;references:ConnectionID"`
	GitUsername  string
	Token        string
	Service      Service `gorm:"polymorphic:Owner;"`
}
type updateGitData struct {
	ConnectionId uint `json:"connectionId"`
}

func (g *Github) String() string {
	return fmt.Sprintf("(@Github: %+v)", *g)
}

type createIssueData struct {
	Owner     string   `json:"owner"`
	RepoName  string   `json:"repositoryName"`
	Body      string   `json:"body"`
	Title     string   `json:"title"`
	Labels    []string `json:"labels"`
	Assignees []string `json:"assignee"`
	State     string   `json:"state"`
}

type createPullRequestData struct {
	Owner    string `json:"owner"`
	RepoName string `json:"repositoryName"`
	Base     string `json:"base"`
	Head     string `json:"head"`
	Title    string `json:"title"`
	Body     string `json:"body"`
}

// Database methods
func (ldb *LegatoDB) CreateGitForScenario(s *Scenario, g Github) (*Github, error) {
	g.Service.UserID = s.UserID
	g.Service.ScenarioID = &s.ID

	ldb.db.Create(&g)
	ldb.db.Save(&g)

	return &g, nil
}

func (ldb *LegatoDB) UpdateGit(s *Scenario, servId uint, gn Github) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var g Github
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&g).Error
	if err != nil {
		return err
	}
	if g.Service.ID != servId {
		return errors.New("the ssh service is not in this scenario")
	}
	var a updateGitData
	err = json.Unmarshal([]byte(gn.Service.Data), &a)
	if err != nil {
		log.Println("con not update ssh")
	}
	if a.ConnectionId != 0 {
		gn.ConnectionID = a.ConnectionId
		user, _ := ldb.GetUserById(g.Service.UserID)
		con, _ := ldb.GetUserConnectionById(&user, gn.ConnectionID)
		gn.Token = con.Data
	}

	ldb.db.Model(&serv).Updates(gn.Service)
	ldb.db.Model(&g).Updates(gn)

	if gn.Service.ParentID == nil {
		legatoDb.db.Model(&serv).Select("parent_id").Update("parent_id", nil)
	}

	return nil
}

func (ldb *LegatoDB) GetGitByID(id uint, u *User) (Github, error) {
	var g Github
	err := ldb.db.Where(&Connection{UserID: u.ID}).Where("ID = ?", id).Find(&g).Error
	if err != nil {
		return Github{}, err
	}
	return g, nil
}

func (ldb *LegatoDB) GetGitByService(serv Service) (*Github, error) {
	var g Github
	err := ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&g).Error
	if err != nil {
		return nil, err
	}
	if g.ID != uint(serv.OwnerID) {
		return nil, errors.New("the Git service is not in this scenario")
	}

	return &g, nil
}

// Service Interface for Git
func (g Github) Execute(Odata *services.Pipe) {
	err := legatoDb.db.Preload("Service").Find(&g).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		g.Next(Odata)
		return
	}
	SendLogMessage("*******Starting Github Service*******", *g.Service.ScenarioID, nil)

	logData := fmt.Sprintf("Executing type (%s) : %s\n", gitType, g.Service.Name)
	SendLogMessage(logData, *g.Service.ScenarioID, nil)

	switch g.Service.SubType {
	case "createIssue":
		var data createIssueData
		err = json.Unmarshal([]byte(g.Service.Data), &data)
		if err != nil {
			log.Print(err)
		}
		token := *&oauth2.Token{}
		err = json.Unmarshal([]byte(g.Token), &token)
		client := createClientForGit(&token)
		NewIssue := &github.IssueRequest{
			Title:     github.String(data.Title),
			Body:      github.String(data.Body),
			Assignees: &data.Assignees,
			Labels:    &data.Labels,
			State:     &data.State,
		}
		// send log
		logData := fmt.Sprintf("Creating a new github issue")
		SendLogMessage(logData, *g.Service.ScenarioID, &g.Service.ID)
		b, err := json.Marshal(NewIssue)
		SendLogMessage(string(b), *g.Service.ScenarioID, &g.Service.ID)

		err = createIssue(NewIssue, data.RepoName, client, data.Owner)
		if err != nil {
			log.Println(err)
		}

	case "createPullRequest":
		var data createPullRequestData
		err = json.Unmarshal([]byte(g.Service.Data), &data)
		if err != nil {
			log.Print(err)
		}
		token := *&oauth2.Token{}
		err = json.Unmarshal([]byte(g.Token), &token)
		client := createClientForGit(&token)
		newPR := &github.NewPullRequest{
			Title:               github.String(data.Title),
			Head:                github.String(data.Head),
			Base:                github.String(data.Base),
			Body:                github.String(data.Body),
			MaintainerCanModify: github.Bool(true),
		}
		// send log
		logData := fmt.Sprintf("Creating a new pull request")
		SendLogMessage(logData, *g.Service.ScenarioID, &g.Service.ID)
		b, err := json.Marshal(newPR)
		SendLogMessage(string(b), *g.Service.ScenarioID, &g.Service.ID)

		err = CreatePullRequest(newPR, data.RepoName, client, data.Owner)
		if err != nil {
			log.Println(err)
		}

	}

	g.Next(Odata)
}

func (g Github) Post(Odata *services.Pipe) {
	log.Printf("Executing type (%s) node in background : %s\n", gitType, g.Service.Name)
}


func (g Github) Resume(data ...interface{}){

}

func (g Github) Next(Odata *services.Pipe) {
	err := legatoDb.db.Preload("Service.Children").Find(&g).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		return
	}

	log.Printf("Executing \"%s\" Children \n", g.Service.Name)

	for _, node := range g.Service.Children {
		go func(n Service) {
			serv, err := n.Load()
			if err != nil {
				log.Println("error in loading services in Next()")
				return
			}

			serv.Execute(Odata)
		}(node)
	}

	logData := fmt.Sprintf("*******End of \"%s\"*******", g.Service.Name)
	SendLogMessage(logData, *g.Service.ScenarioID, nil)
}
func createClientForGit(token *oauth2.Token) *github.Client {
	oauthConf := &oauth2.Config{
		ClientID:     "a87b311ff0542babc5bd",
		ClientSecret: "0d24ae8ec82ca068984ee25e0b6285be9e9c211c",
		Scopes:       []string{"user:email", "repo", "public_repo", "repo_deployment", "write:org", "delete_repo", "read:org"},
		Endpoint:     githubOAuth.Endpoint,
	}

	oauthClient := oauthConf.Client(context.Background(), token)
	client := github.NewClient(oauthClient)
	return client
}
func createIssue(NewIssue *github.IssueRequest, repName string, client *github.Client, owner string) error {
	_, _, err := client.Issues.Create(context.Background(), owner, repName, NewIssue)
	return err
}

func CreatePullRequest(newPR *github.NewPullRequest, repName string, client *github.Client, owner string) error {
	_, _, err := client.PullRequests.Create(context.Background(), owner, repName, newPR)
	return err
}
