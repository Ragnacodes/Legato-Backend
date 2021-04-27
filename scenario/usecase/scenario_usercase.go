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

func (s scenarioUseCase) StartScenario(u *api.UserInfo, scenarioId uint) error {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	err = scenario.Start()
	if err != nil {
		return err
	}

	return nil
}
