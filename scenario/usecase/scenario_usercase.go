package usecase

import (
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"legato_server/models"
	"legato_server/services"
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

func (s scenarioUseCase) AddScenario(u *models.UserInfo, ns *models.NewScenario) error {
	user, _ := s.db.GetUserByUsername(u.Username)
	scenario := converter.NewScenarioToScenarioDb(*ns)

	err := s.db.AddScenario(&user, &scenario)
	if err != nil {
		return err
	}

	return nil
}

func (s scenarioUseCase) TestScenario() {

	log.Println("---------------------------")
	log.Println("Testing Scenario mode")

	// Define events
	root :=
		services.NewWebhook("First Webhook", []services.Service{
			services.NewHttp("Http Event 1", []services.Service{
				services.NewHttp("Http Event 2", []services.Service{
					services.NewHttp("Http Event 3", []services.Service{}),
					services.NewHttp("Http Event 4", []services.Service{}),
					services.NewHttp("Http Event 5", []services.Service{}),
				}),
			}),
		})

	// Create scenario
	ns := legatoDb.Scenario{
		Name: "My first scenario",
		Root: root,
	}

	// Start the scenario
	log.Println("Going to start the scenario...")
	_ = ns.Start()
	log.Println("---------------------------")
}

func (s scenarioUseCase) GetUserScenarios(u *models.UserInfo) ([]models.BriefScenario, error) {
	user := converter.UserInfoToUserDb(*u)
	scenarios, err := s.db.GetUserScenarios(&user)
	if err != nil {
		return nil, err
	}

	var briefScenarios []models.BriefScenario
	for _, scenario := range scenarios {
		briefScenarios = append(briefScenarios, converter.ScenarioDbToBriefScenario(scenario))
	}

	return briefScenarios, nil
}

func (s scenarioUseCase) GetUserScenarioById(u *models.UserInfo, scenarioId string) (models.FullScenario, error) {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return models.FullScenario{}, err
	}

	// Load the whole graph
	scenario.RootService, _ = s.db.GetServicesGraph(scenario.RootService)

	fullScenario := converter.ScenarioDbToFullScenario(scenario)

	return fullScenario, nil
}

func (s scenarioUseCase) UpdateUserScenarioById(u *models.UserInfo, scenarioId string, us models.BriefScenario) error {
	user := converter.UserInfoToUserDb(*u)

	updatedScenario := converter.BriefServiceToServiceDb(us)

	err := s.db.UpdateUserScenarioById(&user, scenarioId, updatedScenario)
	if err != nil {
		return err
	}

	return nil
}
