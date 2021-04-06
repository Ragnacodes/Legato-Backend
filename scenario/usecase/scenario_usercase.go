package usecase

import (
	legatoDb "legato_server/db"
	"legato_server/domain"

	// "legato_server/services"
	"log"
	"time"
)

type scenarioUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewScenarioUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.ScenarioUseCase {
	return &scenarioUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (s scenarioUseCase) AddUserScenario() error {
	return nil
}

func (s scenarioUseCase)TestScenario(){
	time.Sleep(1500 * time.Millisecond)
	log.Println("---------------------------")
	log.Println("Testing Scenario mode")

	//Create some Webhooks
	child := legatoDb.Webhook{Service:legatoDb.Service{Name :"abc"}}
	s.db.Db.Save(&child)
	root := legatoDb.Webhook{Service: legatoDb.Service{Name: "fuck",
			Children: []legatoDb.Service{child.Service}}}
	s.db.Db.Create(&root)

	// Create scenario
	println(root.Service.Name)
	ns := legatoDb.Scenario{
		Name: "My first scenario",
		Root: root.Service,
	}
	log.Println("hi")
	sc := s.db.CreateScenario(ns)
	// Start the scenario
	log.Println("Going to start the scenario...")
	_ = sc.Start()
	log.Println("---------------------------")
}
