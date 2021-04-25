package legatoDb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const httpType string = "https"

type Http struct {
	gorm.Model
	Service Service `gorm:"polymorphic:Owner;"`
}

type httpRequestData struct {
	Url    string
	Method string
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

	if nh.Service.ParentID == nil {
		legatoDb.db.Model(&serv).Select("parent_id").Update("parent_id", nil)
	}

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
	log.Println("*******Starting Http Service*******")

	err := legatoDb.db.Preload("Service").Find(&h).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing type (%s) : %s\n", httpType, h.Service.Name)

	// Http just has one kind of sub service so we do not need any switch-case statement.
	// Provide data for make request
	var data httpRequestData
	err = json.Unmarshal([]byte(h.Service.Data), &data)
	if err != nil {
		log.Fatalln(err)
	}

	_, err = makeHttpRequest(data.Url, data.Method, nil)
	if err != nil {
		log.Fatalln(err)
	}

	h.Next()
}

func (h Http) Post() {
	log.Printf("Executing type (%s) node in background : %s\n", httpType, h.Service.Name)
}

func (h Http) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&h).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing \"%s\" Children \n", h.Service.Name)

	for _, node := range h.Service.Children {
		serv, err := node.Load()
		if err != nil {
			log.Println("error in loading services in Next()")
			return
		}
		serv.Execute()
	}

	log.Printf("*******End of \"%s\"*******", h.Service.Name)
}

// Service interface helper functions
func makeHttpRequest(url string, method string, body []byte) (res *http.Response, err error) {
	log.Println("Make http request")

	switch method {
	case strings.ToLower(http.MethodGet):
		res, err = http.Get(url)
		break
	case strings.ToLower(http.MethodPost):
		if body != nil {
			log.Printf("\nurl: %s\nbody:\n%s\n", url, string(body))
			reqBody := bytes.NewBuffer(body)
			res, err = http.Post(url, "application/json", reqBody)
			break
		}
		res, err = http.Post(url, "application/json", nil)
		break
	}

	if err != nil {
		return nil, err
	}

	// Log the result
	bodyBytes, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}
	bodyString := string(bodyBytes)
	log.Printf("Response from http request is : \n%s\n", bodyString)

	return res, nil
}
