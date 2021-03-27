package domain

import "legato_server/models"

type ScenarioUseCase interface {
	AddScenario(u *models.UserInfo, ns *models.NewScenario) error
	TestScenario()
}
