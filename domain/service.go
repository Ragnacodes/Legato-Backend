package domain

import "legato_server/api"

type ServiceUseCase interface {
	GetServiceNodeById(u *api.UserInfo, scenarioId uint, serviceId uint) (api.ServiceNode, error)
	DeleteServiceNodeById(u *api.UserInfo, scenarioId uint, serviceId uint) error
}
