package converter

import (
	"legato_server/api"
	"legato_server/db"
	"math/rand"
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

	//if s.RootService != nil {
	//	bs.DigestNodes = []string{}
	//} else {
	//	bs.DigestNodes = []string{}
	//}

	return bs
}

func ScenarioDbToFullScenarioGraph(s legatoDb.Scenario) api.FullScenarioGraph {
	fsg := api.FullScenarioGraph{}
	fsg.ID = s.ID
	fsg.Name = s.Name
	fsg.IsActive = s.IsActive
	//fsg.Graph = ServiceDbToService(s.RootService)

	return fsg
}

func FullScenarioGraphToScenarioDb(fsg api.FullScenarioGraph, userID uint) legatoDb.Scenario {
	s := legatoDb.Scenario{}
	s.Name = fsg.Name
	s.IsActive = fsg.IsActive
	// Graph
	//if fsg.Graph != nil {
	//	root := ServiceToServiceDb(fsg.Graph, userID)
	//	s.RootService = &root
	//	s.RootServiceID = &root.ID
	//} else {
	//	s.RootService = nil
	//}

	return s
}

func ScenarioDbToFullScenario(s legatoDb.Scenario) api.FullScenario {
	fs := api.FullScenario{}
	fs.ID = s.ID
	fs.Name = s.Name
	fs.IsActive = s.IsActive
	fs.Interval = rand.Intn(2)
	// Services
	var services []api.Service
	services = []api.Service{}
	for _, s := range s.Services {
		services = append(services, *ServiceDbToService(&s))
	}
	fs.Services = services

	return fs
}

//func FullScenarioToScenarioDb(fs api.FullScenario, userID uint) legatoDb.Scenario {
//	s := legatoDb.Scenario{}
//	s.Name = fs.Name
//	s.IsActive = fs.IsActive
//	// Graph
//
//	return s
//}
