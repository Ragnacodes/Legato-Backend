package domain

import "legato_server/api"

type TelegramUseCase interface {
	AddToScenario(u *api.UserInfo, scenarioId uint, nh api.NewServiceNode) (api.ServiceNode, error)
}
