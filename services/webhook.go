package services

import (
	"log"
	"github.com/satori/go.uuid"
	db "legato_server/db"

)

type Webhook struct{
	Model
	WebhookID uuid.UUID
	Enable bool
}


func NewWebhook(name string,children []db.Service) Service {
	var w Webhook
	w.Type = "webhook"
	w.Name = name
	w.dbChildren = children
	return w
}

func (w Webhook) Execute(attrs ...interface{}) {
	log.Printf("Executing %s node: %s\n", "webhook", w.Name)
	// Listen to trigger
	w.Post();
}

func (w Webhook) Post() {
	log.Printf("Executing %s node in background: %s\n", "webhook", w.Name)
}

func (w Webhook) Next(attrs ...interface{}) {
	for _, node := range w.dbChildren {
		child := entityToService(node)
		child.Execute(attrs...)
	} 
}
