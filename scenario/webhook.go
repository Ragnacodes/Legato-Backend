package scenario

import (
	"log"
	"time"
)

type WebhookHandler struct {
	name string
}

func NewWebhookHandler(name string) WebhookHandler {
	return WebhookHandler{
		name: name,
	}
}

func (wh *WebhookHandler) Observe(events []Event) {
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
	wh.Trigger(events)
}

func (wh *WebhookHandler) Trigger(events []Event) {
	log.Println("Trigger")

	// Do each action in tree
	for _, e := range events {
		log.Println(e.Name())
		log.Println(e.Type())
		e.Execute()
		e.Post()
	}
}

func (wh *WebhookHandler) Type() string {
	return "Webhook"
}

func (wh *WebhookHandler) Name() string {
	return wh.name
}
