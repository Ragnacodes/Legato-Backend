package scheduler

import (
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
	"log"
	"os"
	"os/signal"
	"syscall"
)

var Redis = redis.NewClient(&redis.Options{
	Addr: "192.168.1.20:6379",
})

var (
	QueueFactory = redisq.NewFactory()
	MainQueue    = QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  "api-worker",
		Redis: Redis,
	})
	StartScenarioTask = taskq.RegisterTask(&taskq.TaskOptions{
		Name: "start_scenario",
		Handler: func(scenarioID string) error {
			log.Println("scenario should start here")
			return nil
		},
	})
)

func WaitSignal() os.Signal {
	ch := make(chan os.Signal, 2)
	signal.Notify(
		ch,
		syscall.SIGINT,
		syscall.SIGQUIT,
		syscall.SIGTERM,
	)
	for {
		sig := <-ch
		switch sig {
		case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
			return sig
		}
	}
}
