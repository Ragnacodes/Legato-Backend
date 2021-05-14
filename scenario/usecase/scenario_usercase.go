package usecase

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/env"
	"legato_server/helper/converter"
	"log"
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

func (s scenarioUseCase) StartScenarioInstantly(u *api.UserInfo, scenarioId uint) error {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	err = scenario.Start(true)
	if err != nil {
		return err
	}

	return nil
}

// ForceStartScenario is not accessible for users.
// It is used for starting the scheduled scenarios by scheduler.
func (s scenarioUseCase) ForceStartScenario(scenarioId uint) error {
	scenario, err := s.db.GetScenarioById(scenarioId)
	if err != nil {
		return err
	}

	err = scenario.Start(false)
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
	schedulerUrl := fmt.Sprintf("%s/api/schedule/scenario/%d", env.ENV.SchedulerUrl, scenarioId)
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

func (s scenarioUseCase) SetInterval(u *api.UserInfo, scenarioId uint, interval *api.NewScenarioInterval) error {
	user := converter.UserInfoToUserDb(*u)
	_, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	// Update the interval for this scenario
	err = s.db.UpdateScenarioIntervalById(&user, scenarioId, interval.Interval)
	if err != nil {
		return err
	}

	// Get scenario to act based on isActive
	// if it is active, it should be scheduled for interval minutes later.
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	log.Println(scenario.String())
	if scenario.IsActive == nil {
		return errors.New("this scenario has null isActive field")
	}
	if *scenario.IsActive {
		minutes := time.Duration(scenario.Interval) * time.Minute
		schedule := &api.NewStartScenarioSchedule{
			ScheduledTime: time.Now().Add(minutes),
			SystemTime:    time.Now(),
		}
		// Make http request to enqueue this job
		schedulerUrl := fmt.Sprintf("%s/api/schedule/scenario/%d", env.ENV.SchedulerUrl, scenarioId)
		body, err := json.Marshal(schedule)
		if err != nil {
			return err
		}
		reqBody := bytes.NewBuffer(body)
		_, err = http.Post(schedulerUrl, "application/json", reqBody)
		if err != nil {
			return err
		}
	}

	return nil
}
