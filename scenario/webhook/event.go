package webhook

import (
	"legato_server/scenario"
	"log"
)

type Event scenario.Event

func (e *Event) Execute() {
	log.Println("Run in main thread")
}

func (e *Event) Post() {
	log.Println("Run in background thread")
}
