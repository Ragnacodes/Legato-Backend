package legatoDb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"

	"gorm.io/gorm"
)

const httpType string = "https"

type Http struct {
	gorm.Model
	Service Service `gorm:"polymorphic:Owner;"`
}

type httpRequestData struct {
	Url    string
	Method string
	Body   map[string]interface{}
}

type httpGetRequestData struct {
	Url    string
	Method string
	Body   string
}

func (w *httpRequestData) UnmarshalJSON(data []byte) error {
	var getData httpGetRequestData
	if err := json.Unmarshal(data, &httpGetRequestData{}); err == nil {
	  w.Url = getData.Url
	  w.Method = getData.Method
	  w.Body = make(map[string]interface{})
	  return nil
	}
	return nil
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
	err := legatoDb.db.Preload("Service").Find(&h).Error
	if err != nil {
		panic(err)
	}
	SendLogMessage("*******Starting Http Service*******", *h.Service.ScenarioID, nil)

	logData := fmt.Sprintf("Executing type (%s) : %s\n", httpType, h.Service.Name)
	SendLogMessage(logData, *h.Service.ScenarioID, nil)
	// Http just has one kind of sub service so we do not need any switch-case statement.
	// Provide data for make request
	var data httpRequestData
	err = json.Unmarshal([]byte(h.Service.Data), &data)
	if err != nil {
		log.Fatalln(err)
	}

	requestBody, err := json.Marshal(data.Body)
	if err != nil {
		log.Fatalln(err)
	}
	_, err = makeHttpRequest(data.Url, data.Method, requestBody, nil ,h.Service.ScenarioID, &h.Service.ID)
	if err != nil {
		log.Fatalln(err)
	}

	h.Next()
}

func (h Http) Post() {
	data := fmt.Sprintf("Executing type (%s) node in background : %s\n", httpType, h.Service.Name)
	SendLogMessage(data, *h.Service.ScenarioID, nil) 
}

func (h Http) Next(...interface{}) {
	err := legatoDb.db.Preload("Service").Preload("Service.Children").Find(&h).Error
	if err != nil {
		panic(err)
	}

	logData := fmt.Sprintf("Executing \"%s\" Children \n", h.Service.Name)
	SendLogMessage(logData, *h.Service.ScenarioID, nil)

	for _, node := range h.Service.Children {
		serv, err := node.Load()
		if err != nil {
			log.Println("error in loading services in Next()")
			return
		}
		serv.Execute()
	}

	logData = fmt.Sprintf("*******End of \"%s\"*******", h.Service.Name)
	SendLogMessage(logData, *h.Service.ScenarioID, nil)
}

// Service interface helper functions
func makeHttpRequest(url string, method string, body []byte, authorization *string, scenarioId *uint, hId *uint) (res *http.Response, err error) {
	logData := fmt.Sprintf("Make http request")
	SendLogMessage(logData, *scenarioId, hId)

	logData = fmt.Sprintf("\nurl: %s\nmethod: %s", url, method)
	SendLogMessage(logData, *scenarioId, hId)

	SendLogMessage(string(body), *scenarioId, hId)

	switch method {
	case strings.ToLower(http.MethodGet):
		res, err = http.Get(url)
		break
	case strings.ToLower(http.MethodPost):
		if body != nil {
			client := &http.Client{}
			req, err := http.NewRequest(http.MethodPost, url, bytes.NewBuffer(body))
			if err != nil {
				return nil, err
			}
			if authorization != nil {
				req.Header.Set("Authorization", *authorization)
			}
			req.Header.Set("Content-Type", "application/json")
			res, err = client.Do(req)
			if err != nil {
				return nil, err
			}
			break
		}
		res, err = http.Post(url, "application/json", nil)
		break
	case strings.ToLower(http.MethodPut):
		if body != nil {
			client := &http.Client{}
			req, err := http.NewRequest(http.MethodPut, url, bytes.NewBuffer(body))
			if err != nil {
				return nil, err
			}
			if authorization != nil {
				req.Header.Set("Authorization", *authorization)
			}
			req.Header.Set("Content-Type", "application/json")
			res, err = client.Do(req)
			if err != nil {
				return nil, err
			}
		} else {
			log.Println("body in put request is empty")
			client := &http.Client{}
			req, err := http.NewRequest(http.MethodPut, url, nil)
			if err != nil {
				return nil, err
			}
			if authorization != nil {
				req.Header.Set("Authorization", *authorization)
			}
			req.Header.Set("Content-Type", "application/json")
			res, err = client.Do(req)
			if err != nil {
				return nil, err
			}
		}
		break
	default:
		break
	}

	if err != nil {
		return nil, err
	}

	// Log the result
	bodyString := ""
	if res != nil && res.Body != nil {
		bodyBytes, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return nil, err
		}
		bodyString = string(bodyBytes)
	}

	logData = fmt.Sprintf("Got Respose from http request")
	SendLogMessage(logData, *scenarioId, hId)

	SendLogMessage(bodyString, *scenarioId, hId)

	return res, nil
}
