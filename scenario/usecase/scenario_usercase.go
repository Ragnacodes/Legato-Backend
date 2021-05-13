package usecase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"net/http"
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
	scenario, err := converter.NewScenarioToScenarioDb(*ns)
	if err != nil {
		return api.BriefScenario{}, err
	}

	err = s.db.AddScenario(&user, &scenario)
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
		briefScenario := converter.ScenarioDbToBriefScenario(scenario)
		nodes, err := s.db.GetScenarioNodeTypes(&scenario)
		if err == nil {
			for _, node := range nodes {
				briefScenario.DigestNodes = append(briefScenario.DigestNodes, node.OwnerType)
			}
		}
		briefScenarios = append(briefScenarios, briefScenario)
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

	updatedScenario, err := converter.NewScenarioToScenarioDb(ns)
	if err != nil {
		return err
	}

	err = s.db.UpdateUserScenarioById(&user, scenarioId, updatedScenario)
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

func (s scenarioUseCase) ForceStartScenario(scenarioId uint) error {
	scenario, err := s.db.GetScenarioById(scenarioId)
	if err != nil {
		return err
	}

	err = scenario.Start()
	if err != nil {
		return err
	}

	return nil
}


func (s scenarioUseCase) Schedule(u *api.UserInfo, scenarioId uint, schedule *api.NewStartScenarioSchedule) error {
	user := converter.UserInfoToUserDb(*u)
	_, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	// Make http request to enqueue this job
	schedulerUrl  := fmt.Sprintf("http://192.168.1.20:8090/api/schedule/scenario/%d", scenarioId)
	body, err := json.Marshal(schedule)
	if err != nil {
		return err
	}
	reqBody := bytes.NewBuffer(body)
	_, err = http.Post(schedulerUrl, "application/json", reqBody)
	if err != nil {
		return err
	}

	return nil
}
