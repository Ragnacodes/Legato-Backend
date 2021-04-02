package services

import "legato_server/db"
// Model is the base model for every services.
// Each service should have this Model struct as an embedded struct.
type Model struct {
	Name     string
	Type     string
	dbChildren []legatoDb.Service
	Children []Service
}

// Service contains details about provided service.
// Execute runs the related action in the main thread.
// Post runs the related actions in the background thread.
// Next runs the next node(s)
type Service interface {
	Execute(attrs ...interface{})
	Post()
	Next(atrrs ...interface{})
}


func entityToService(service legatoDb.Service) Service{
	switch service.Type{
	case "webhook":
		return NewWebhook(service.Name, service.Children)
	case "http":
		return NewHttp(service.Name, service.Children)
	default:
		return nil
	}
}