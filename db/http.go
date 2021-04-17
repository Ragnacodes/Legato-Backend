package legatoDb

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

const httpType string = "https"

type Http struct {
	gorm.Model
	Url     string
	Method  string
	Service Service `gorm:"polymorphic:Owner;"`
}

func (h *Http) String() string {
	return fmt.Sprintf("(@Http: %+v)", *h)
}

// Database methods
func (ldb *LegatoDB) CreateHttp(s *Scenario, h Http) (*Http, error) {
	h.Service.UserID = s.UserID
	h.Service.ScenarioID = &s.ID

	ldb.db.Create(&h)
	ldb.db.Save(&h)

	return &h, nil
}

func (ldb *LegatoDB) UpdateHttp(s *Scenario, servId uint, nh Http) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var h Http
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&h).Error
	if err != nil {
		return err
	}
	if h.Service.ID != servId {
		return errors.New("the http service is not in this scenario")
	}

	ldb.db.Model(&serv).Updates(nh.Service)
	ldb.db.Model(&h).Updates(nh)

	return nil
}

func (ldb *LegatoDB) GetHttpByService(serv Service) (*Http, error) {
	var h Http
	err := ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&h).Error
	if err != nil {
		return nil, err
	}
	if h.ID != uint(serv.OwnerID) {
		return nil, errors.New("the http service is not in this scenario")
	}

	return &h, nil
}

// Service Interface for Http
func (h Http) Execute(...interface{}) {
	err := legatoDb.db.Preload("Service").Find(&h).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing %s node: %s\n", "http", h.Service.Name)

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

	log.Printf("Executing %s node in background: %s\n", "http", h.Service.Name)
}

func (h Http) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&h).Error
	if err != nil {
		panic(err)
	}

	for _, node := range h.Service.Children {
		serv, err := node.Load()
		if err != nil {
			log.Println("error in loading services in Next()")
			return
		}
		serv.Execute()
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
