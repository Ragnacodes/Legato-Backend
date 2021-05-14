package converter

import (
	"errors"
	"legato_server/api"
	legatoDb "legato_server/db"
	"time"
)

func NewScenarioToScenarioDb(ns api.NewScenario) (legatoDb.Scenario, error) {
	s := legatoDb.Scenario{}
	s.Name = ns.Name
	s.IsActive = ns.IsActive
	if s.IsActive == nil {
		return legatoDb.Scenario{}, errors.New("isActive can not be null")
	}
	s.Services = []legatoDb.Service{}

	return s, nil
}

func ScenarioDbToBriefScenario(s legatoDb.Scenario) api.BriefScenario {
	bs := api.BriefScenario{}
	bs.ID = s.ID
	bs.Name = s.Name
	bs.Interval = s.Interval
	bs.IsActive = s.IsActive
	bs.DigestNodes = []string{}

	return bs
}

func ScenarioDbToFullScenario(s legatoDb.Scenario) api.FullScenario {
	fs := api.FullScenario{}
	fs.ID = s.ID
	fs.Name = s.Name
	fs.IsActive = s.IsActive
	fs.LastScheduledTime = s.LastScheduledTime.Format(time.RFC3339)
	fs.Interval = s.Interval
	// Services
	var services []api.ServiceNode
	services = []api.ServiceNode{}
	for _, s := range s.Services {
		services = append(services, ServiceDbToServiceNode(s))
	}
	fs.Services = services

	return fs
}
