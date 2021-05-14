package legatoDb

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"legato_server/api"
	"legato_server/env"
	"legato_server/services"
	"log"
	"net/http"
	"time"

	"gorm.io/gorm"
)

// Each Scenario describes a schema that includes Handler and Events.
// Name is the title of that Scenario.
// Root is the first Service of the schema that start the scenario.
type Scenario struct {
	gorm.Model
	UserID            uint
	Name              string
	IsActive          *bool
	Interval          int32
	RootServices      []services.Service `gorm:"-"`
	Services          []Service
	ScheduleToken     []byte
	LastScheduledTime time.Time
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
		Find(&sc).Error
	if err != nil {
		return Scenario{}, err
	}

	return sc, nil
}

func (ldb *LegatoDB) GetScenarioById(scenarioId uint) (Scenario, error) {
	var sc Scenario
	err := ldb.db.
		Where("id = ?", scenarioId).
		Preload("Services").
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

func (ldb *LegatoDB) DeleteUserScenarioById(u *User, scenarioID uint) error {
	var scenario Scenario
	ldb.db.Where(&Scenario{UserID: u.ID}).Where("id = ?", scenarioID).Find(&scenario)
	if scenario.ID != scenarioID {
		return errors.New("the scenario is not in user scenarios")
	}

	// Note: webhook and http records should be deleted here, too
	ldb.db.Where("scenario_id = ?", scenario.ID).Delete(&Service{})
	ldb.db.Delete(&scenario)
	return nil
}

func (ldb *LegatoDB) UpdateScenarioScheduleInfoById(
	u *User, scenarioID uint, lastScheduledTime time.Time, interval int32,
) error {
	var scenario Scenario
	ldb.db.Where(&Scenario{UserID: u.ID}).Where("id = ?", scenarioID).Find(&scenario)
	if scenario.ID != scenarioID {
		return errors.New("the scenario is not in user scenarios")
	}

	ldb.db.Model(&scenario).Updates(&Scenario{Interval: interval, LastScheduledTime: lastScheduledTime})

	return nil
}

func generateRandomKey() []byte {
	key := make([]byte, 8)
	if _, err := rand.Read(key); err != nil {
		panic(err)
	}
	fmt.Printf("Key: %b \n", key)

	return key
}

func (ldb *LegatoDB) SetNewScheduleToken(u *User, scenarioID uint) ([]byte, error) {
	var scenario Scenario
	ldb.db.Where(&Scenario{UserID: u.ID}).Where("id = ?", scenarioID).Find(&scenario)
	if scenario.ID != scenarioID {
		return []byte{}, errors.New("the scenario is not in user scenarios")
	}

	// Generate new token
	token := generateRandomKey()

	ldb.db.Model(&scenario).Updates(&Scenario{ScheduleToken: token})

	return token, nil
}

// Service management methods

// Start
// To Start scenario
// isInstantMode specify whether the scenario started instantly or not.
// when the isInstantMode is ture the scenario just executed once.
func (s *Scenario) Start() error {

	log.Println("Preparing scenario to start")
	err := s.Prepare()
	if err != nil {
		return err
	}

	log.Println("Executing root services of this scenario")
	go func() {
		for _, serv := range s.RootServices {
			serv.Execute()
		}
	}()
	log.Println("Executing finished")

	return nil
}

func (s Scenario) Schedule(scheduleToken []byte) error {
	if s.Interval != 0 {
		log.Printf("Scheduling the scenario for %d minutes later\n", s.Interval)
		minutes := time.Duration(s.Interval) * time.Minute
		schedule := &api.NewStartScenarioSchedule{
			ScheduledTime: time.Now().Add(minutes),
			SystemTime:    time.Now(),
			Token:         scheduleToken,
		}
		// Make http request to enqueue this job
		schedulerUrl := fmt.Sprintf("%s/api/schedule/scenario/%d", env.ENV.SchedulerUrl, s.ID)
		body, err := json.Marshal(schedule)
		if err != nil {
			return err
		}
		reqBody := bytes.NewBuffer(body)
		_, err = http.Post(schedulerUrl, "application/json", reqBody)
		if err != nil {
			return err
		}
		log.Println("Scenario Scheduled successfully")
	}

	return nil
}

// Prepare
// To Prepare scenario
func (s *Scenario) Prepare() error {
	err := s.LoadRootService()
	if err != nil {
		return err
	}

	return nil
}

// LoadRootService
// To Load Root Service of scenario
func (s *Scenario) LoadRootService() error {
	servicesEntities, err := legatoDb.GetScenarioRootServices(*s)
	if err != nil {
		return err
	}

	var ss []services.Service
	ss = []services.Service{}
	for _, serv := range servicesEntities {
		loadedServ, err := serv.Load()
		if err != nil {
			return nil
		}

		ss = append(ss, loadedServ)
	}
	s.RootServices = ss

	return nil
}

type OwnerType struct {
	OwnerType string
}

func (ldb *LegatoDB) GetScenarioNodeTypes(scenario *Scenario) (t []OwnerType, err error) {
	err = ldb.db.Distinct("OwnerType").Model(&Service{}).
		Where(&Service{ScenarioID: &scenario.ID}).
		Find(&t).Error

	if err != nil {
		return []OwnerType{}, err
	}

	return t, nil
}
