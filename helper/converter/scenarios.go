package converter

import (
	"legato_server/api"
	"legato_server/db"
)

func NewScenarioToScenarioDb(ns api.NewScenario) legatoDb.Scenario {
	s := legatoDb.Scenario{}
	s.Name = ns.Name
	s.IsActive = ns.IsActive

	return s
}

func ScenarioDbToBriefScenario(s legatoDb.Scenario) api.BriefScenario {
	bs := api.BriefScenario{}
	bs.ID = s.ID
	bs.Name = s.Name
	bs.IsActive = s.IsActive

	if s.RootService != nil {
		bs.DigestNodes = []string{}
	} else {
		bs.DigestNodes = []string{}
	}

	return bs
}

func ScenarioDbToFullScenario(s legatoDb.Scenario) api.FullScenario {
	fs := api.FullScenario{}
	fs.ID = s.ID
	fs.Name = s.Name
	fs.IsActive = s.IsActive
	fs.Graph = ServiceDbToService(s.RootService)

	return fs
}

func FullScenarioToScenarioDb(fs api.FullScenario, userID uint) legatoDb.Scenario {
	s := legatoDb.Scenario{}
	s.Name = fs.Name
	s.IsActive = fs.IsActive
	// Graph
	if fs.Graph != nil {
		root := ServiceToServiceDb(fs.Graph, userID)
		s.RootService = &root
	} else {
		s.RootService = nil
	}

	return s
}
