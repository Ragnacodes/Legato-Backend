package main

import (
	"legato_server/scheduler"
	"log"
)

func init() {
	log.Println("Starting the scheduler server.")
}

func main() {
	_ = scheduler.NewRouter().Run(":8090")
}
