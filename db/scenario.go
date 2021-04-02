package legatoDb

import (
	"fmt"
	"gorm.io/gorm"
	"legato_server/services"
	"log"
)

// Each Scenario describes a schema that includes Handler and Events.
// Name is the title of that Scenario.
// Root is the first Service of the schema that start the scenario.
type Scenario struct {
	gorm.Model
	UserID        uint
	Name          string
	IsActive      bool
	RootServiceID *uint
	RootService   *Service         `gorm:"RootServiceID:"`
	Root          services.Service `gorm:"-"`
}

func (s *Scenario) String() string {
	return fmt.Sprintf("(@Scenario: %+v)", *s)
}

// Scenario methods
// To Start scenario
func (s *Scenario) Start() error {
	s.Root.Execute()

	return nil
}

// Scenario database
func (ldb *LegatoDB) AddScenario(u *User, s *Scenario) error {
	s.UserID = u.ID

	ldb.db.Create(&s)
	ldb.db.Save(&s)

	return nil
}

func (ldb *LegatoDB) GetUserScenarios(u *User) ([]Scenario, error) {
	user, _ := ldb.GetUserByUsername(u.Username)

	var scenarios []Scenario
	ldb.db.Model(&user).Association("Scenarios").Find(&scenarios)

	return scenarios, nil
}

func (ldb *LegatoDB) GetUserScenarioById(u *User, scenarioId string) (Scenario, error) {
	var sc Scenario
	err := ldb.db.
		Where(&Scenario{UserID: u.ID}).
		Where("id = ?", scenarioId).
		Preload("RootService").Find(&sc).Error
	if err != nil {
		return Scenario{}, err
	}

	log.Println(sc.String())

	return sc, nil
}

func (ldb *LegatoDB) GetScenarioByName(u *User, name string) (Scenario, error) {
	var sc Scenario
	err := ldb.db.Where(&Scenario{Name: name, UserID: u.ID}).Preload("RootService").Find(&sc).Error
	if err != nil {
		return Scenario{}, err
	}

	return sc, nil
}
