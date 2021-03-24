package usecase

import (
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/scenario"
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

func (s scenarioUseCase) TestScenario() {
	log.Println("---------------------------")
	log.Println("Testing Scenario mode")

	// Define events
	event1 := scenario.NewHttpEvent("Event number 1")
	event2 := scenario.NewHttpEvent("Event number 2")
	event3 := scenario.NewHttpEvent("Event number 3")
	events := []scenario.Event{
		&event1,
		&event2,
		&event3,
	}

	// Create handler
	handler := scenario.NewWebhookHandler("My Webhook")

	// Create scenario
	ns := scenario.Scenario{
		Name:    "My first scenario",
		Handler: &handler,
		Events:  events,
	}

	// Start the scenario
	ns.Start()
}
