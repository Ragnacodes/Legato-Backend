package services

import (
	"errors"
	"fmt"
	"log"
	"strings"
)

// Pipe is an object that be passed to the next service in the scenario.
// It used to transfer data between all services in a scenario
type Pipe struct {
	DataByName map[string]interface{}
}

// AddData is to add some data about a node to the Pipe
// So that later we can take that data from Pipe
func (p *Pipe) AddData(serviceName string, serviceData interface{}) {
	p.DataByName[serviceName] = serviceData
}

// GetData returns the data related to specific node by their name
func (p *Pipe) GetData(serviceName string) (interface{}, error) {
	result, found := p.DataByName[serviceName]
	if !found {
		err := fmt.Errorf("no such service %s", serviceName)
		return nil, err
	}

	return result, nil
}
