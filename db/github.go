package legatoDb

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log"

	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
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
	ConnectionId uint   `json:"connectionid"`
	GitUsername  string `json:"username"`
}

func (g *Github) String() string {
	return fmt.Sprintf("(@Git: %+v)", *g)
}

type CreateIssueData struct {
	Owner     string       `json:"owner"`
	Token     oauth2.Token `json:"token"`
	RepoName  string       `json:"repositoryName"`
	Body      string       `json:"body"`
	Title     string       `json:"title"`
	Labels    []string     `json:"labels"`
	Assignees []string     `json:"assignee"`
	State     string       `json:"state"`
}

type CreatePullRequestData struct {
	Owner    string       `json:"owner"`
	Token    oauth2.Token `json:"token"`
	RepoName string       `json:"repositoryName"`
	Base     string       `json:"base"`
	Head     string       `json:"head"`
	Title    string       `json:"title"`
	Body     string       `json:"body"`
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
	fmt.Println(gn.Service.Data)
	fmt.Println(a.ConnectionId)
	if a.ConnectionId != 0 {
		gn.ConnectionID = a.ConnectionId
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
func (g Github) Execute(...interface{}) {
	log.Println("*******Starting Git Service*******")

	err := legatoDb.db.Preload("Service").Find(&g).Error
	if err != nil {
		panic(err)
	}

	switch g.Service.SubType {
	case "create issue":
		var data CreateIssueData
		err = json.Unmarshal([]byte(g.Service.Data), &data)
		if err != nil {
			log.Print(err)
		}
		client := createClientForGit(&data.Token)
		NewIssue := &github.IssueRequest{
			Title:     github.String(data.Title),
			Body:      github.String(data.Body),
			Assignees: &data.Assignees,
			Labels:    &data.Labels,
			State:     &data.State,
		}
		err = createIusse(NewIssue, data.RepoName, client, data.Owner)
		if err != nil {
			log.Println(err)
		} else {
			log.Print("issue created successfully")
		}

	case "create pull request":
		var data CreatePullRequestData
		err = json.Unmarshal([]byte(g.Service.Data), &data)
		if err != nil {
			log.Print(err)
		}
		client := createClientForGit(&data.Token)
		newPR := &github.NewPullRequest{
			Title:               github.String(data.Title),
			Head:                github.String(data.Head),
			Base:                github.String(data.Base),
			Body:                github.String(data.Body),
			MaintainerCanModify: github.Bool(true),
		}
		err = CreatePullRequest(newPR, data.RepoName, client, data.Owner)
		if err != nil {
			log.Println(err)
		} else {
			log.Print("pull request created successfully")
		}

	}

	log.Printf("Executing type (%s) : %s\n", gitType, g.Service.Name)

	g.Next()
}

func (g Github) Post() {
	log.Printf("Executing type (%s) node in background : %s\n", gitType, g.Service.Name)
}

func (g Github) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&g).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing \"%s\" Children \n", g.Service.Name)

	for _, node := range g.Service.Children {
		serv, err := node.Load()
		if err != nil {
			log.Println("error in loading services in Next()")
			return
		}
		serv.Execute()
	}

	log.Printf("*******End of \"%s\"*******", g.Service.Name)
}
func createClientForGit(token *oauth2.Token) *github.Client {
	oauthConf := &oauth2.Config{
		ClientID:     "a87b311ff0542babc5bd",
		ClientSecret: "0d24ae8ec82ca068984ee25e0b6285be9e9c211c",
		Scopes:       []string{"user:email", "repo", "public_repo", "repo_deployment", "write:org", "delete_repo", "read:org"},
		Endpoint:     githuboauth.Endpoint,
	}

	oauthClient := oauthConf.Client(context.Background(), token)
	client := github.NewClient(oauthClient)
	return client
}
func createIusse(NewIssue *github.IssueRequest, repName string, client *github.Client, owner string) error {
	_, _, err := client.Issues.Create(context.Background(), owner, repName, NewIssue)
	return err
}

func CreatePullRequest(newPR *github.NewPullRequest, repName string, client *github.Client, owner string) error {
	_, _, err := client.PullRequests.Create(context.Background(), owner, repName, newPR)
	return err
}
