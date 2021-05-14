package scheduler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/vmihailenco/taskq/v3"
	"legato_server/api"
	"legato_server/env"
	"log"
	"net/http"
)

const StartScenarioTask = "StartScenarioTask"

var Tasks = legatoTasks{
	StartScenarioTask: taskq.RegisterTask(&taskq.TaskOptions{
		Name:    StartScenarioTask,
		Handler: startScenario,
	}),
}

func startScenario(scenarioID int, token []byte) error {
	log.Printf("scenario %d had been scheduled and is going to trigger.\n", scenarioID)

	data := &api.NewStartScenarioSchedule{
		Token: token,
	}
	body, err := json.Marshal(data)
	if err != nil {
		log.Fatal(err)
	}

	// Make http request to do run this scenario
	schedulerUrl := fmt.Sprintf("%s/api/scenarios/%d/force", env.DefaultLegatoUrl, scenarioID)
	client := &http.Client{}
	req, err := http.NewRequest("POST", schedulerUrl, bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Authorization", AccessToken)
	_, err = client.Do(req)
	if err != nil {
		return err
	}

	return nil
}
