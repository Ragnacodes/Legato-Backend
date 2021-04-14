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
	s.Id = service.ID
	s.UserId = &service.UserID
	// ParentID
	if service.ParentID != nil {
		s.ParentId = service.ParentID
	} else {
		s.ParentId = nil
	}
	s.Name = service.Name
	s.Type = service.OwnerType
	s.Position = api.Position{X: service.PosX, Y: service.PosY}
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
	s.ID = service.Id
	s.Name = service.Name
	s.OwnerType = service.Type
	s.PosX, s.PosY = service.Position.X, service.Position.Y
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

func NewServiceNodeToServiceDb(sn api.NewServiceNode) legatoDb.Service {
	var s legatoDb.Service
	s.ParentID = sn.ParentId
	s.OwnerType = sn.Type
	s.Name = sn.Name
	s.PosX, s.PosY = sn.Position.X, sn.Position.Y

	return s
}

func ServiceDbToServiceNode(s legatoDb.Service) api.ServiceNode {
	var sn api.ServiceNode
	sn.Id = s.ID
	sn.ParentId = s.ParentID
	sn.Type = s.OwnerType
	sn.Name = s.Name
	sn.Position = api.Position{X: s.PosX, Y: s.PosY}

	return sn
}
