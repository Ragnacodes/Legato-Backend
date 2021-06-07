package api

import "golang.org/x/oauth2"

type GitInfo struct {
	Id               uint          `json:"id"`
	GitUsername      string        `json:"username"`
	Token            *oauth2.Token `json:"token"`
	ConnectionID     uint          `json:"connectionId"`
	RepositoriesName string        `json:"repositoryName"`
}
