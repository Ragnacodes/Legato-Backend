package scheduler

import (
	"github.com/go-redis/redis/v8"
	"github.com/vmihailenco/taskq/v3"
	"github.com/vmihailenco/taskq/v3/redisq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var Redis *redis.Client

var QueueFactory taskq.Factory
var MainQueue taskq.Queue

type legatoTasks []*taskq.Task

func CreateQueue(redisAddress string) error {
	log.Println("Connecting to redis....")
	Redis = redis.NewClient(&redis.Options{
		Addr: redisAddress,
	})

	QueueFactory = redisq.NewFactory()
	MainQueue = QueueFactory.RegisterQueue(&taskq.QueueOptions{
		Name:  "api-worker",
		Redis: Redis,
	})
	
	return nil
}

func LogStats() {
	for range time.Tick(3 * time.Second) {
		log.Println("checking...")
	}
}

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
