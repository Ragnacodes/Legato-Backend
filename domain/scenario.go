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
	Schedule(u *api.UserInfo, scenarioId uint, schedule *api.NewStartScenarioSchedule) error
	StartScenarioInstantly(u *api.UserInfo, scenarioId uint) error
	ForceStartScenario(scenarioId uint) error
	SetInterval(userInfo *api.UserInfo, scenarioId uint, interval *api.NewScenarioInterval) error
}
