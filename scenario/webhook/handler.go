package webhook

import (
	"legato_server/scenario"
	"log"
	"time"
)

type Handler struct {
	scenario.Handler
	Events []Event
}

func (h *Handler) Observe() {
	log.Println("Observe")

	// Listen to trigger
	i := 0
	for i < 5 {
		log.Printf("second number %d \n", i+1)
		time.Sleep(time.Second * 1)
		i += 1
	}

	// Trigger
	log.Println("Going to trigger...")
	h.Trigger()
}

func (h *Handler) Trigger() {
	log.Println("Trigger")

	// Do each action in tree
	for _, event := range h.Events {
		log.Println(event.Name)
		event.Execute()
		event.Post()
	}
}
