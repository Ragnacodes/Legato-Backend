package converter

import (
	legatoDb "legato_server/db"
	"legato_server/models"
)

func ServiceDbToService(service *legatoDb.Service) *models.Service {
	if service == nil {
		return nil
	}

	var s models.Service
	s.Name = service.Name
	s.Type = service.Type
	s.Position = models.Position{X: service.Position.X, Y: service.Position.Y}
	s.Data = struct{}{}

	if len(service.Children) == 0 {
		s.Children = []models.Service{}
		return &s
	}

	var children []models.Service
	for _, child := range service.Children {
		childSubGraph := ServiceDbToService(&child)
		children = append(children, *childSubGraph)
	}

	s.Children = children

	return &s
}

func ServiceToServiceDb(service *models.Service) legatoDb.Service {
	var s legatoDb.Service
	s.Name = service.Name
	s.Type = service.Type
	s.Position = legatoDb.Position{X: service.Position.X, Y: service.Position.Y}
	//s.Data = struct{}{}

	if len(service.Children) == 0 {
		s.Children = []legatoDb.Service{}
		return s
	}

	var children []legatoDb.Service
	for _, child := range service.Children {
		childSubGraph := ServiceToServiceDb(&child)
		children = append(children, childSubGraph)
	}

	s.Children = children

	return s
}
