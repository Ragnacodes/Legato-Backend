package scheduler

import (
	"github.com/vmihailenco/taskq/v3"
	"log"
)

var Tasks = legatoTasks{
	taskq.RegisterTask(&taskq.TaskOptions{
		Name: "start_scenario",
		Handler: func(scenarioID string) error {
			log.Println("scenario should start here")
			return nil
		},
	}),
}