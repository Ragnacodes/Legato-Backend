package legatoDb

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Http struct {
	gorm.Model
	Url     string
	Method  string
	Service Service `gorm:"polymorphic:Owner;"`
}

func (h *Http) String() string {
	return fmt.Sprintf("(@Http: %+v)", *h)
}

// Service Interface for Http
func (h Http) Execute(...interface{}) {
	err := legatoDb.db.Preload("Service").Find(&h).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing %s node: %s\n", "webhook", h.Service.Name)

	_, err = makeHttpRequest(h.Url, h.Method)
	if err != nil {
		log.Fatalln(err)
	}

	h.Next()
}

func (h Http) Post() {
	err := legatoDb.db.Preload("Service").Find(&h).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing %s node in background: %s\n", "webhook", h.Service.Name)
}

func (h Http) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&h).Error
	if err != nil {
		panic(err)
	}

	for _, node := range h.Service.Children {
		node.LoadOwner().Execute()
	}
}

// Service interface helper functions
func makeHttpRequest(url string, method string) (res *http.Response, err error) {
	switch method {
	case http.MethodGet:
		res, err = http.Get(url)
		break
	case http.MethodPost:
		res, err = http.Post(url, "application/json", nil)
		break
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
