package legatoDb

import (
	"log"
	"fmt"
	"gorm.io/gorm"
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
}

func (s *Scenario) String() string {
	return fmt.Sprintf("(@Scenario: %+v)", *s)

}

func(l *LegatoDB) CreateScenario(sc Scenario) *Scenario{
	l.Db.Create(&sc)
	return &sc
}
// To Start scenario
func (s *Scenario) Start() error {
	log.Printf("Scenario root %s is Executing:", s.RootService.Name)
	s.RootService.LoadOwner().Execute()
	return nil
}


func (ldb *LegatoDB) AddScenario(s *Scenario) error {
	ldb.Db.Create(&s)
	ldb.Db.Save(&s)

	return nil
}

func (ldb *LegatoDB) GetScenarioByName(u *User, name string) (Scenario, error) {
	var sc Scenario
	err := ldb.Db.Where(&Scenario{Name: name, UserID: u.ID}).Preload("RootService").Find(&sc).Error
	if err != nil {
		return Scenario{}, err
	}

	return sc, nil
}
