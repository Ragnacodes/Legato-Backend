package main

import (
	"legato_server/scheduler"
	"log"
	"time"
)

const redisAddress = "192.168.1.20:6379"

func init() {
	err := scheduler.CreateQueue(redisAddress)
	if err != nil {
		panic(err)
	}

	time.Sleep(3 * time.Second)

	log.Println("Start log stats")
	go scheduler.LogStats()

	log.Println("Start Consuming")
	go scheduler.Listen()
}

func main() {
	log.Println("Starting the scheduler server.")
	_ = scheduler.NewRouter().Run(":8090")
}
