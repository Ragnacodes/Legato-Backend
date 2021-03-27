package helper

import (
	"legato_server/db"
	"legato_server/models"
)

func NewScenarioToScenarioEntity(newUser models.NewScenario) legatoDb.Scenario {
	u := legatoDb.Scenario{}
	u.Name = newUser.Name

	return u
}
