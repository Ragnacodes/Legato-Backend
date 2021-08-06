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

func (p *Pipe) GetValueByNodeVariable(nr NodeVariable) (string, error) {
	// Get the whole data from node
	data, err := p.GetData(nr.NodeName)
	if err != nil {
		return "", err
	}

	// Search for the value
	// Note: for now all of the data types are json
	// try to cast to map[string]interface{}
	_, canCast := data.(map[string]interface{})
	if !canCast {
		return "", errors.New("data is not json, it can't be cast to map[string]interface{}")
	}

	// Search for the desire attribute
	// It is going to search recursively until it find the value
	json := data
	for _, attr := range nr.AttributeChain {
		// Check being null
		if json == nil {
			return "", errors.New("there is not such attribute in this data")
		}

		// Converting and casting to map[string]interface{}
		// TODO: Here we may have array and we have no idea about it :)
		converted, canCast := json.(map[string]interface{})
		if !canCast {
			return "", errors.New("there is not such attribute in this data")
		}

		// Grab the attribute
		json = converted[attr]
	}

	// Convert json to the string
	// TODO: we should about this. Is everything end to a string? yes (!)
	converted, canCast := json.(string)
	if !canCast {
		return "", errors.New("can not cast the result to string value")
	}

	return converted, nil
}
