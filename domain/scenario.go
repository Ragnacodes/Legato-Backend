package domain

import (
	"legato_server/api"
)

type ScenarioUseCase interface {
	AddScenario(userInfo *api.UserInfo, newScenario *api.NewScenario) (api.BriefScenario, error)
	GetUserScenarios(userInfo *api.UserInfo) ([]api.BriefScenario, error)
	GetUserScenarioById(userInfo *api.UserInfo, scenarioId uint) (api.FullScenario, error)
	UpdateUserScenarioById(userInfo *api.UserInfo, scenarioId uint, updated api.NewScenario) error
	DeleteUserScenarioById(u *api.UserInfo, scenarioId uint) error
	StartScenario(u *api.UserInfo, scenarioId uint) error
}
