package legatoDb

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
)

// Each Scenario describes a schema that includes Handler and Events.
// Name is the title of that Scenario.
// Root is the first Service of the schema that start the scenario.
type Scenario struct {
	gorm.Model
	UserID        uint
	Name          string
	IsActive      *bool
	//RootServiceID *uint
	//RootService   *Service `gorm:"RootServiceID:"`
	Services      []Service
}

func (s *Scenario) String() string {
	return fmt.Sprintf("(@Scenario: %+v)", *s)
}

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

func (ldb *LegatoDB) GetUserScenarioById(u *User, scenarioId uint) (Scenario, error) {
	var sc Scenario
	err := ldb.db.
		Where(&Scenario{UserID: u.ID}).
		Where("id = ?", scenarioId).
		Preload("Services").
		//Preload("RootService").
		Find(&sc).Error
	if err != nil {
		return Scenario{}, err
	}

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

func (ldb *LegatoDB) UpdateUserScenarioById(u *User, scenarioID uint, updatedScenario Scenario) error {
	var scenario Scenario
	ldb.db.Where(&Scenario{UserID: u.ID}).Where("id = ?", scenarioID).Find(&scenario)
	if scenario.ID != scenarioID {
		return errors.New("the scenario is not in user scenarios")
	}

	ldb.db.Model(&scenario).Updates(updatedScenario)

	return nil
}

// Methods

// To Start scenario
func (s *Scenario) Start() error {
	//log.Printf("Scenario root %s is Executing:", s.RootService.Name)
	//s.RootService.LoadOwner().Execute()
	return nil
}
