package domain

import "legato_server/models"

type ScenarioUseCase interface {
	AddScenario(userInfo *models.UserInfo, newScenario *models.NewScenario) error
	GetUserScenarios(userInfo *models.UserInfo) ([]models.BriefScenario, error)
	GetUserScenarioById(userInfo *models.UserInfo, scenarioId string) (models.FullScenario, error)
	TestScenario()
}
