package converter

import (
	"legato_server/db"
	"legato_server/models"
)

func NewScenarioToScenarioDb(ns models.NewScenario) legatoDb.Scenario {
	s := legatoDb.Scenario{}
	s.Name = ns.Name
	s.IsActive = ns.IsActive

	return s
}

func ScenarioDbToBriefScenario(s legatoDb.Scenario) models.BriefScenario {
	bs := models.BriefScenario{}
	bs.ID = s.ID
	bs.Name = s.Name
	bs.IsActive = s.IsActive

	if s.RootService != nil {
		bs.DigestNodes = []string{s.RootService.Type}
	} else {
		bs.DigestNodes = []string{}
	}

	return bs
}
