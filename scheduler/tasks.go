package scheduler

import (
	"github.com/vmihailenco/taskq/v3"
	"log"
)

const StartScenarioTask = "StartScenarioTask"

var Tasks = legatoTasks{
	StartScenarioTask: taskq.RegisterTask(&taskq.TaskOptions{
		Name:    StartScenarioTask,
		Handler: startScenario,
	}),
}

func startScenario(scenarioID int) error {
	log.Printf("scenario %d should start here\n", scenarioID)
	return nil
}
