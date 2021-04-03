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

func ScenarioDbToFullScenario(s legatoDb.Scenario) models.FullScenario {
	fs := models.FullScenario{}
	fs.ID = s.ID
	fs.Name = s.Name
	fs.IsActive = s.IsActive
	fs.Graph = ServiceDbToService(s.RootService)

	return fs
}

func FullScenarioToScenarioDb(fs models.FullScenario) legatoDb.Scenario {
	s := legatoDb.Scenario{}
	s.Name = fs.Name
	s.IsActive = fs.IsActive
	// Graph
	if fs.Graph != nil {
		root := ServiceToServiceDb(fs.Graph)
		s.RootService = &root
	} else {
		s.RootService = nil
	}

	return s
}
