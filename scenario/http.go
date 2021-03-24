package scenario

import (
	"log"
)

type HttpEvent struct {
	name string
}

func NewHttpEvent(name string) HttpEvent {
	return HttpEvent{
		name: name,
	}
}

func (e *HttpEvent) Execute() {
	log.Println("Run in main thread")
}

func (e *HttpEvent) Post() {
	log.Println("Run in background thread")
}

func (e *HttpEvent) Type() string {
	return "HTTP"
}

func (e *HttpEvent) Name() string {
	return e.name
}
