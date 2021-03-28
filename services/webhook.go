package services

import (
	"log"
	"time"
	"github.com/satori/go.uuid"
)

type WebhookService struct{
	Model
	WebhookID uuid.UUID
}

func NewWebhook(name string, children []Service) Service {
	var w WebhookService
	w.Type = "webhook"
	w.Name = name
	w.Children = children

	return w
}

func (w WebhookService) Execute() {
	log.Printf("Executing %s node: %s\n", "webhook", w.Name)
	// Listen to trigger
	i := 0
	for i < 5 {
		log.Printf("second number %d \n", i+1)
		time.Sleep(time.Second * 1)
		i += 1
	}

	w.Next()
}

func (w WebhookService) Post() {
	log.Printf("Executing %s node in background: %s\n", "webhook", w.Name)
}

func (w WebhookService) Next() {
	children := w.Children
	for _, node := range children {

		node.Execute()
	}
}
