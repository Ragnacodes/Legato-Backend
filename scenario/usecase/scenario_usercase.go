package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
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

func (s scenarioUseCase) AddScenario(u *api.UserInfo, ns *api.NewScenario) (api.BriefScenario, error) {
	user, _ := s.db.GetUserByUsername(u.Username)
	scenario := converter.NewScenarioToScenarioDb(*ns)

	err := s.db.AddScenario(&user, &scenario)
	if err != nil {
		return api.BriefScenario{}, err
	}

	return converter.ScenarioDbToBriefScenario(scenario), nil
}

func (s scenarioUseCase) GetUserScenarios(u *api.UserInfo) ([]api.BriefScenario, error) {
	user := converter.UserInfoToUserDb(*u)
	scenarios, err := s.db.GetUserScenarios(&user)
	if err != nil {
		return nil, err
	}

	var briefScenarios []api.BriefScenario
	briefScenarios = []api.BriefScenario{}
	for _, scenario := range scenarios {
		briefScenarios = append(briefScenarios, converter.ScenarioDbToBriefScenario(scenario))
	}

	return briefScenarios, nil
}

// Deprecated
func (s scenarioUseCase) GetUserScenarioGraphById(u *api.UserInfo, scenarioId uint) (api.FullScenarioGraph, error) {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.FullScenarioGraph{}, err
	}

	// Load the whole graph
	//scenario.RootService, _ = s.db.GetServicesGraph(scenario.RootService)

	fullScenario := converter.ScenarioDbToFullScenarioGraph(scenario)

	return fullScenario, nil
}

func (s scenarioUseCase) GetUserScenarioById(u *api.UserInfo, scenarioId uint) (api.FullScenario, error) {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.FullScenario{}, err
	}

	fullScenario := converter.ScenarioDbToFullScenario(scenario)

	return fullScenario, nil
}

func (s scenarioUseCase) UpdateUserScenarioById(u *api.UserInfo, scenarioId uint, ns api.NewScenario) error {
	user := converter.UserInfoToUserDb(*u)

	updatedScenario := converter.NewScenarioToScenarioDb(ns)

	err := s.db.UpdateUserScenarioById(&user, scenarioId, updatedScenario)
	if err != nil {
		return err
	}

	return nil
}

func (s scenarioUseCase) DeleteUserScenarioById(u *api.UserInfo, scenarioId uint) error {
	user := converter.UserInfoToUserDb(*u)

	err := s.db.DeleteUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	return nil
}

func (s scenarioUseCase) TestScenario() {
	//time.Sleep(1500 * time.Millisecond)
	//log.Println("---------------------------")
	//log.Println("Testing Scenario mode")
	//
	////Create some Webhooks
	//child := legatoDb.Webhook{Service:legatoDb.Service{Name :"abc"}}
	//s.db.Db.Save(&child)
	//root := legatoDb.Webhook{Service: legatoDb.Service{Name: "fuck",
	//	Children: []legatoDb.Service{child.Service}}}
	//s.db.Db.Create(&root)
	//
	//// Create scenario
	//println(root.Service.Name)
	//ns := legatoDb.Scenario{
	//	Name: "My first scenario",
	//	RootService: &root.Service,
	//}
	//log.Println("hi")
	//sc := s.db.CreateScenario(ns)
	//// Start the scenario
	//log.Println("Going to start the scenario...")
	//_ = sc.Start()
	//log.Println("---------------------------")
}

// This is for testing purposes
// It puts a default scenario for legato user in the database.
func (s *scenarioUseCase) CreateDefaultScenario() error {
	// Default user
	user := &api.UserInfo{
		Username: "legato",
	}

	// Default scenario
	isActive := false
	newScenario := &api.NewScenario{
		Name: "my_very_first_scenario",
		IsActive: &isActive,
	}

	// Create the scenario
	ns, err := s.AddScenario(user, newScenario)
	if err != nil {
		return err
	}

	// Create root
	u, err := s.db.GetUserByUsername(user.Username)
	scenario, err := s.db.GetUserScenarioById(&u, ns.ID)

	// Create some services
	// Root service
	webhookRoot := s.db.CreateWebhookInScenario(&u, &scenario, nil, "My starter webhook")
	// Second level children
	firstHttp := s.db.CreateWebhookInScenario(&u, &scenario, &webhookRoot.Service, "My first http")
	_ = s.db.CreateWebhookInScenario(&u, &scenario, &webhookRoot.Service, "another http")
	// Third level children
	_ = s.db.CreateWebhookInScenario(&u, &scenario, &firstHttp.Service, "My second http")

	return nil
}
