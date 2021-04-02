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

func (s scenarioUseCase) TestScenario() {

	log.Println("---------------------------")
	log.Println("Testing Scenario mode")

	// Define events
	// root :=
	// 	services.NewWebhook("First Webhook", []services.Service{
	// 		services.NewHttp("Http Event 1", []services.Service{
	// 			services.NewHttp("Http Event 2", []services.Service{
	// 				services.NewHttp("Http Event 3", []services.Service{}),
	// 				services.NewHttp("Http Event 4", []services.Service{}),
	// 				services.NewHttp("Http Event 5", []services.Service{}),
	// 			}),
	// 		}),
	// 	})

	// Create scenario
	ns := legatoDb.Scenario{
		Name: "My first scenario",
	}

	// Start the scenario
	log.Println("Going to start the scenario...")
	_ = ns.Start()
	log.Println("---------------------------")
}
