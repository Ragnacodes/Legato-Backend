package domain

import (
	"legato_server/api"
)

type ScenarioUseCase interface {
	AddScenario(userInfo *api.UserInfo, newScenario *api.NewScenario) (api.BriefScenario, error)
	GetUserScenarios(userInfo *api.UserInfo) ([]api.BriefScenario, error)
	GetUserScenarioById(userInfo *api.UserInfo, scenarioId string) (api.FullScenario, error)
	UpdateUserScenarioById(userInfo *api.UserInfo, scenarioId string, updated api.FullScenario) error
	TestScenario()
}
