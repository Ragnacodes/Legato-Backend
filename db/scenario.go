package legatoDb

import (
	"gorm.io/gorm"
	
)

// Each Scenario describes a schema that includes Handler and Events.
// Name is the title of that Scenario.
// Root is the first Service of the schema that start the scenario.
type Scenario struct {
	gorm.Model
	UserID uint
	Name   string
	
}

// To Start scenario
func (s *Scenario) Start() error {
	
	return nil
}
