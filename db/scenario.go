package legatoDb

import (
	"fmt"
	"gorm.io/gorm"
	"legato_server/services"
	"log"
	"strconv"
)

// Each Scenario describes a schema that includes Handler and Events.
// Name is the title of that Scenario.
// Root is the first Service of the schema that start the scenario.
type Scenario struct {
	gorm.Model
	UserID        uint
	Name          string
	IsActive      *bool
	RootServiceID *uint
	RootService   *Service         `gorm:"RootServiceID:"`
	Root          services.Service `gorm:"-"`
}

func (s *Scenario) String() string {
	return fmt.Sprintf("(@Scenario: %+v)", *s)
}

func (ldb *LegatoDB) AddScenario(u *User, s *Scenario) error {
	log.Println(s.String())
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


func (ldb *LegatoDB) UpdateUserScenarioById(u *User, scenarioID string, updatedScenario Scenario) error {
	sid, _ :=  strconv.Atoi(scenarioID)
	updatedScenario.ID = uint(sid)
	ldb.db.Session(&gorm.Session{FullSaveAssociations: true}).Updates(&updatedScenario)

	return nil
}


// Scenario methods

// To Start scenario
func (s *Scenario) Start() error {
	s.RootService.LoadOwner().Execute()
	return nil
}