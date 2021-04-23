package legatoDb

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"legato_server/services"
)

type Service struct {
	gorm.Model
	Name       string
	OwnerID    int
	OwnerType  string
	ParentID   *uint
	Children   []Service `gorm:"foreignkey:ParentID"`
	PosX       int
	PosY       int
	UserID     uint
	ScenarioID *uint
}

func (s *Service) String() string {
	return fmt.Sprintf("(@Service: %+v)", *s)
}

func (ldb *LegatoDB) GetServiceById(scenario *Scenario, serviceId uint) (*Service, error) {
	var srv *Service
	err := ldb.db.
		Where(&Service{ScenarioID: &scenario.ID}).
		Where("id = ?", serviceId).
		Find(&srv).Error
	if err != nil {
		return nil, err
	}

	return srv, nil
}

func (ldb *LegatoDB) DeleteServiceById(scenario *Scenario, serviceId uint) error {
	var srv *Service
	ldb.db.Where(&Service{ScenarioID: &scenario.ID}).Where("id = ?", serviceId).Find(&srv)
	if srv.ID != serviceId {
		return errors.New("the service is not in this scenario")
	}

	// Note: webhook and http records should be deleted here, too
	ldb.db.Delete(srv)

	// Attach children to the parent
	parentId := srv.ParentID
	ldb.db.Where(&Service{ParentID: &srv.ID}).Updates(Service{ParentID: parentId})

	return nil
}

func (ldb *LegatoDB) GetServicesGraph(root *Service) (*Service, error) {
	if root == nil {
		return nil, nil
	}

	err := ldb.db.Preload("Children").Preload("Position").Find(&root).Error
	if err != nil {
		return nil, err
	}

	if len(root.Children) == 0 {
		return root, nil
	}

	var children []Service
	for _, child := range root.Children {
		childSubGraph, err := ldb.GetServicesGraph(&child)
		if err != nil {
			return nil, err
		}

		children = append(children, *childSubGraph)
	}

	root.Children = children

	return root, nil
}

// Load
// It Load the service entity to a services.Service
// so that we can execute the scenario for them.
func (s *Service) Load() (services.Service, error) {
	var serv services.Service
	var err error
	switch s.OwnerType {
	case webhookType:
		serv, err = legatoDb.GetWebhookByService(*s)
		break
	case httpType:
		serv, err = legatoDb.GetHttpByService(*s)
		break
	}
	if err != nil {
		return nil, err
	}

	return serv, nil
}

// GetServiceData
// Each one of services have some special data. By giving the Service model
// this function returns a map of those data.
func (s *Service) GetServiceData() (map[string]interface{}, error) {
	data := make(map[string]interface{})

	switch s.OwnerType {
	case webhookType:
		break
	case httpType:
		h, err := legatoDb.GetHttpByService(*s)
		if err != nil {
			return nil, err
		}

		data["url"] = h.Url
		data["method"] = h.Method
		break
	}

	return data, nil
}

func (ldb *LegatoDB) AppendChildren(parent *Service, children []Service) {
	parent.Children = append(parent.Children, children...)
	ldb.db.Save(&parent)
}
