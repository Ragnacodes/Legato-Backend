package services

// Model is the base model for every services.
// Each service should have this Model struct as an embedded struct.
type Model struct {
	Name     string
	Type     string
	Children []Service
}

// Service contains details about provided service.
// Execute runs the related action in the main thread.
// Post runs the related actions in the background thread.
// Next runs the next node(s)
type Service interface {
	Execute()
	Post()
	Next()
}

//basic implemetation of service interface for services struct

func Execution(s Service) {
	s.Execute()
}

func Postpone(s Service) {
	s.Post()
}

func NextService(s Service) {
	s.Next()
}

