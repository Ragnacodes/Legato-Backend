package converter

import (
	"legato_server/api"
	legatoDb "legato_server/db"
)

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
	sn.Data, _ = s.GetServiceData()

	return sn
}
