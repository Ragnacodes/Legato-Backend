package legatoDb

import (
	// "fmt"
	"log"

	"gorm.io/gorm"
)

// Each Scenario describes a schema that includes Handler and Events.
// Name is the title of that Scenario.
// Root is the first Service of the schema that start the scenario.
type Scenario struct {
	gorm.Model
	UserID uint
	Name   string
	RootID int 
	Root Service `gorm:"foreignKey:RootID"`
}

func(l *LegatoDB) CreateScenario(sc Scenario) *Scenario{
	l.Db.Create(&sc)
	return &sc
}
// To Start scenario
func (s *Scenario) Start() error {
	log.Printf("Scenario root %s is Executing:", s.Root.Name)
	s.Root.LoadOwner().Execute()
	return nil
}

