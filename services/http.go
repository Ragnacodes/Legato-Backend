package services

import (
	"log"
	"time"
)

type Http struct {
	Model
}

func NewHttp(name string, children []Service) Service {
	var h Http
	h.Type = "http"
	h.Name = name
	h.Children = children
	return h
}

func (h Http) Execute(attrs ...interface{}) {
	log.Printf("Executing %s node: %s\n", "http", h.Name)
	time.Sleep(time.Second)

	h.Next(attrs)
}

func (h Http) Post() {
	log.Printf("Executing %s node in background: %s\n", "http", h.Name)
}

func (h Http) Next(attrs ...interface{}) {
	children := h.Children
	for _, node := range children {
		node.Execute(attrs...)
	}
}
