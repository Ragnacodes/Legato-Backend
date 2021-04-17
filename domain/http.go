package domain

import (
	"legato_server/api"
)

type HttpUseCase interface {
	AddToScenario(userInfo *api.UserInfo, scenarioId uint, newWebhook api.NewServiceNode) (api.ServiceNode, error)
	Update(u *api.UserInfo, scenarioId uint, nodeId uint, nw api.NewServiceNode) error
}
