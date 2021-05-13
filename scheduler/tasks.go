package scheduler

import (
	"fmt"
	"github.com/vmihailenco/taskq/v3"
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

func startScenario(scenarioID int) error {
	log.Printf("scenario %d had been scheduled and is going to trigger.\n", scenarioID)

	// Make http request to do run this scenario
	schedulerUrl := fmt.Sprintf("http://192.168.1.20:8080/api/scenarios/%d/force", scenarioID)
	client := &http.Client{}
	req, err := http.NewRequest("POST", schedulerUrl, nil)
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
