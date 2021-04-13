package legatoDb

import (
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
	PosX         int
	PosY         int
	UserID     uint
	ScenarioID uint
}

func (s *Service) String() string {
	return fmt.Sprintf("(@Service: %+v)", *s)
}

func (ldb *LegatoDB) GetServiceById(scenario *Scenario, serviceId uint) (*Service, error) {
	var srv *Service
	err := ldb.db.
		Where(&Service{ScenarioID: scenario.ID}).
		Where("id = ?", serviceId).
		Find(&srv).Error
	if err != nil {
		return nil, err
	}

	return srv, nil
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

func (s *Service) LoadOwner() services.Service {
	var wh Webhook
	query := fmt.Sprintf("SELECT * FROM %s WHERE id = %d", s.OwnerType, s.OwnerID)
	err := legatoDb.db.Raw(query).Scan(&wh).Error
	if err != nil {
		panic(err)
	}

	return wh
}

func (ldb *LegatoDB) AppendChildren(parent *Service, children []Service) {
	parent.Children = append(parent.Children, children...)
	ldb.db.Save(&parent)
}
