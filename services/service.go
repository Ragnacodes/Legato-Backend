package services

import (
	"fmt"
)

// Service contains details about provided service.
// Execute runs the related action in the main thread.
// Post runs the related actions in the background thread.
// Next runs the next node(s)
type Service interface {
	Execute(*OutputData)
	Post(*OutputData)
	Resume(data ...interface{})
	Next(*OutputData)
}

// used to transfer data between all services in a scenario
type OutputData struct {
	Data map[string]interface{}
}

func (dict *OutputData) AddData(serviceName string, serviceData interface{}) {
	dict.Data[serviceName] = serviceData
}

func (dict *OutputData) GetData(serviceName string) (interface{}, error) {
	result, found := dict.Data[serviceName]
	if !found {
		err := fmt.Errorf("no such service %s", serviceName)
		return nil, err
	}
	return result, nil
}
