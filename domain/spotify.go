package domain

import (
	"legato_server/api"
	"golang.org/x/oauth2"
)

type SpotifyUseCase interface {
	AddToScenario(userInfo *api.UserInfo, scenarioId uint, nh api.NewServiceNode) (api.ServiceNode, error)
	Update(u *api.UserInfo, scenarioId uint, nodeId uint, nt api.NewServiceNode) error
	CreateSpotifyToken(userInfo api.UserInfo, token *oauth2.Token) (error)
	GetUserToken(userInfo api.UserInfo) (token *oauth2.Token, err error)
}