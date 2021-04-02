package converter

import (
	legatoDb "legato_server/db"
	"legato_server/models"
)

func ServiceDbToService(service *legatoDb.Service) models.Service {
	var s models.Service
	s.Name = service.Name
	s.Type = service.Type
	s.Data = struct{}{}

	if len(service.Children) == 0 {
		s.Children = []models.Service{}
		return s
	}

	var children []models.Service
	for _, child := range service.Children {
		childSubGraph := ServiceDbToService(&child)
		children = append(children, childSubGraph)
	}

	s.Children = children

	return s

}
