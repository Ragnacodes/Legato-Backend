package converter

import (
	"legato_server/api"
	legatoDb "legato_server/db"
)

func ServiceDbToService(service *legatoDb.Service) *api.Service {
	if service == nil {
		return nil
	}

	var s api.Service
	s.Name = service.Name
	s.Type = service.OwnerType
	s.Position = api.Position{X: service.Position.X, Y: service.Position.Y}
	s.Data = struct{}{}

	if len(service.Children) == 0 {
		s.Children = []api.Service{}
		return &s
	}

	var children []api.Service
	for _, child := range service.Children {
		childSubGraph := ServiceDbToService(&child)
		children = append(children, *childSubGraph)
	}

	s.Children = children

	return &s
}

func ServiceToServiceDb(service *api.Service, userID uint) legatoDb.Service {
	var s legatoDb.Service
	s.Name = service.Name
	s.OwnerType = service.Type
	s.Position = legatoDb.Position{X: service.Position.X, Y: service.Position.Y}
	s.UserID = userID
	//s.Data = struct{}{}

	if len(service.Children) == 0 {
		s.Children = []legatoDb.Service{}
		return s
	}

	var children []legatoDb.Service
	for _, child := range service.Children {
		childSubGraph := ServiceToServiceDb(&child, userID)
		children = append(children, childSubGraph)
	}

	s.Children = children

	return s
}
