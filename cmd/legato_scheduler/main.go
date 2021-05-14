package main

import (
	"legato_server/env"
	"legato_server/scheduler"
	"log"
	"time"
)

func init() {
	// Load environment variables
	env.LoadEnv()

	err := scheduler.CreateQueue(env.ENV.RedisHost)
	if err != nil {
		panic(err)
	}

	time.Sleep(1 * time.Second)

	log.Println("Start log stats")
	go scheduler.LogStats()

	log.Println("Start Consuming")
	go scheduler.Listen()
}

func main() {
	log.Println("Starting the scheduler server.")
	_ = scheduler.NewRouter().Run(":8090")
}
