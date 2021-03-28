package legatoDb

import (
	"fmt"
	"gorm.io/gorm"
)

type Service struct {
	gorm.Model
	Name     string
	Type     string
	ParentID *uint
	Children []Service `gorm:"foreignkey:ParentID"`
}

func (s *Service) String() string {
	return fmt.Sprintf("(@Service: %+v)", *s)
}

func (ldb *LegatoDB) GetServicesGraph(root *Service) (*Service, error) {
	err := ldb.db.Preload("Children").Find(&root).Error
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