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

func (s scenarioUseCase) GetUserScenarioGraphById(u *api.UserInfo, scenarioId string) (api.FullScenarioGraph, error) {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.FullScenarioGraph{}, err
	}

	// Load the whole graph
	scenario.RootService, _ = s.db.GetServicesGraph(scenario.RootService)

	fullScenario := converter.ScenarioDbToFullScenarioGraph(scenario)

	return fullScenario, nil
}

func (s scenarioUseCase) GetUserScenarioById(u *api.UserInfo, scenarioId string) (api.FullScenario, error) {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.FullScenario{}, err
	}

	fullScenario := converter.ScenarioDbToFullScenario(scenario)

	return fullScenario, nil
}

func (s scenarioUseCase) UpdateUserScenarioById(u *api.UserInfo, scenarioId string, us api.FullScenarioGraph) error {
	user := converter.UserInfoToUserDb(*u)

	updatedScenario := converter.FullScenarioGraphToScenarioDb(us, u.ID)

	err := s.db.UpdateUserScenarioById(&user, scenarioId, updatedScenario)
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
