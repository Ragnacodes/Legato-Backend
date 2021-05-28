package domain

import (
	"legato_server/api"
	// legatoDb "legato_server/db"
)

type GitUseCase interface {
	GetGitWithId(id uint, username string) (api.GitInfo, error)
	AddToScenario(u *api.UserInfo, scenarioId uint, ns api.NewServiceNode) (api.ServiceNode, error)
	Update(u *api.UserInfo, scenarioId uint, serviceId uint, ns api.NewServiceNode) error
}
