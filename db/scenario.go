package legatoDb

import (
	"fmt"
	"gorm.io/gorm"
	"legato_server/services"
)

// Each Scenario describes a schema that includes Handler and Events.
// Name is the title of that Scenario.
// Root is the first Service of the schema that start the scenario.
type Scenario struct {
	gorm.Model
	UserID uint
	Name   string
	RootServiceID *uint
	RootService   *Service         `gorm:"RootServiceID:"`
	Root          services.Service `gorm:"-"`
}

func (s *Scenario) String() string {
	return fmt.Sprintf("(@Scenario: %+v)", *s)
}

func(ldb *LegatoDB) CreateScenario(sc Scenario) *Scenario{
	ldb.db.Create(&sc)
	return &sc
}

func (ldb *LegatoDB) AddScenario(s *Scenario) error {
	ldb.db.Create(&s)
	ldb.db.Save(&s)

	return nil
}

func (ldb *LegatoDB) GetScenarioByName(u *User, name string) (Scenario, error) {
	var sc Scenario
	err := ldb.db.Where(&Scenario{Name: name, UserID: u.ID}).Preload("RootService").Find(&sc).Error
	if err != nil {
		return Scenario{}, err
	}

	return sc, nil
}

// Scenario methods

// To Start scenario
func (s *Scenario) Start() error {
	s.RootService.LoadOwner().Execute()
	return nil
}