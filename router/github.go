package router

import (
	"context"
	"encoding/json"
	"legato_server/api"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	githuboauth "golang.org/x/oauth2/github"
)

var GitRG = routeGroup{
	name: "Github",
	routes: routes{
		route{
			"Show user repositories name",
			POST,
			"/users/:username/services/github/repositories",
			getUserRepositoryList,
		},
		route{
			"show list of branches of a repository",
			POST,
			"/users/:username/services/github/branches",
			getUserRepositoryBranches,
		},
	},
}

func getUserRepositoryList(c *gin.Context) {
	username := c.Param("username")
	githubData := api.GitInfo{}
	_ = c.BindJSON(&githubData)

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}

	oauthConf := &oauth2.Config{
		ClientID:     "a87b311ff0542babc5bd",
		ClientSecret: "0d24ae8ec82ca068984ee25e0b6285be9e9c211c",
		Scopes:       []string{"user:email", "repo", "public_repo", "repo_deployment", "write:org", "delete_repo", "read:org"},
		Endpoint:     githuboauth.Endpoint,
	}
	conn, _ := resolvers.UserUseCase.GetConnectionByID(username, githubData.ConnectionID)
	token := *&oauth2.Token{}
	_ = json.Unmarshal([]byte(conn.Data), &token)

	oauthClient := oauthConf.Client(context.Background(), &token)
	client := github.NewClient(oauthClient)
	repos, _, err := client.Repositories.List(context.Background(), "", nil)
	if err != nil {
		c.JSON(503, err)
	}

	var repoName []string
	for _, repo := range repos {
		repoName = append(repoName, *repo.FullName)
	}

	c.JSON(http.StatusOK, gin.H{
		"repositoriesName": repoName,
	})
}

func getUserRepositoryBranches(c *gin.Context) {
	username := c.Param("username")
	githubData := api.GitInfo{}
	_ = c.BindJSON(&githubData)

	// Auth
	loginUser := checkAuth(c, []string{username})
	if loginUser == nil {
		return
	}
	oauthConf := &oauth2.Config{
		ClientID:     "a87b311ff0542babc5bd",
		ClientSecret: "0d24ae8ec82ca068984ee25e0b6285be9e9c211c",
		Scopes:       []string{"user:email", "repo", "public_repo", "repo_deployment", "write:org", "delete_repo", "read:org"},
		Endpoint:     githuboauth.Endpoint,
	}
	conn, _ := resolvers.UserUseCase.GetConnectionByID(username, githubData.ConnectionID)
	token := *&oauth2.Token{}
	_ = json.Unmarshal([]byte(conn.Data), &token)

	oauthClient := oauthConf.Client(context.Background(), &token)
	client := github.NewClient(oauthClient)
	repoAndOwner := strings.Split(githubData.RepositoriesName, "/")
	repoName := strings.Replace(repoAndOwner[1], "/", "", 1)
	owner := strings.Replace(repoAndOwner[0], "/", "", 1)
	branches, _, err := client.Repositories.ListBranches(context.Background(), owner, repoName, nil)
	var branchName []string
	for _, branch := range branches {
		branchName = append(branchName, *branch.Name)
	}
	if err != nil {
		c.JSON(503, gin.H{
			"branchesName": []string{},
		})
	}

	c.JSON(http.StatusOK, gin.H{
		"branchesName": branchName,
	})
}
