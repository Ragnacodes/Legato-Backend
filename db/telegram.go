package legatoDb

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"legato_server/env"
	"log"
	"net/http"
	"net/url"
	"strings"

	"golang.org/x/net/proxy"
	"gorm.io/gorm"
)

const telegramType string = "telegrams"

// Sub services
const sendMessage string = "sendMessage"
const sendMessageEndpoint string = "https://api.telegram.org/bot%s/sendMessage"
const getChatMember string = "getChatMember"
const getChatMemberEndpoint string = "https://api.telegram.org/bot%s/getChatMember"

type Telegram struct {
	gorm.Model
	Key     string
	Service Service `gorm:"polymorphic:Owner;"`
}

type sendMessageData struct {
	ChatId string `json:"chat_id"`
	Text   string `json:"text"`
}

type getChatMemberData struct {
	ChatId string `json:"chat_id"`
	UserId string `json:"user_id"`
}

func (t *Telegram) String() string {
	return fmt.Sprintf("(@Telegram: %+v)", *t)
}

// Database methods
func (ldb *LegatoDB) CreateTelegram(s *Scenario, telegram Telegram) (*Telegram, error) {
	telegram.Service.UserID = s.UserID
	telegram.Service.ScenarioID = &s.ID

	ldb.db.Create(&telegram)
	ldb.db.Save(&telegram)

	return &telegram, nil
}

func (ldb *LegatoDB) UpdateTelegram(s *Scenario, servId uint, nt Telegram) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var t Telegram
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&t).Error
	if err != nil {
		return err
	}
	if t.Service.ID != servId {
		return errors.New("the telegram service is not in this scenario")
	}

	ldb.db.Model(&serv).Updates(nt.Service)
	ldb.db.Model(&t).Updates(nt)

	if nt.Service.ParentID == nil {
		legatoDb.db.Model(&serv).Select("parent_id").Update("parent_id", nil)
	}

	return nil
}

func (ldb *LegatoDB) GetTelegramByService(serv Service) (*Telegram, error) {
	var t Telegram
	err := ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&t).Error
	if err != nil {
		return nil, err
	}
	if t.ID != uint(serv.OwnerID) {
		return nil, errors.New("the telegram service is not in this scenario")
	}

	return &t, nil
}

// Service Interface for telegram
func (t Telegram) Execute(...interface{}) {
	log.Println("*******Starting Telegram Service*******")

	err := legatoDb.db.Preload("Service").Find(&t).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing type (%s) : %s\n", telegramType, t.Service.Name)

	switch t.Service.SubType {
	case sendMessage:
		var data sendMessageData
		err = json.Unmarshal([]byte(t.Service.Data), &data)
		if err != nil {
			log.Fatal(err)
		}

		if env.ENV.Mode == env.DEVELOPMENT{
			_, err = makeHttpRequest(fmt.Sprintf(sendMessageEndpoint, t.Key), "post", []byte(t.Service.Data), nil, t.Service.ScenarioID, &t.Service.ID)
		} else {
		_, err = makeTorifiedHttpRequest(fmt.Sprintf(sendMessageEndpoint, t.Key), "post", []byte(t.Service.Data), t.Service.ScenarioID, &t.Service.ID)
		}

		if err != nil {
			log.Fatal(err)
		}
		break
	case getChatMember:
		var data getChatMemberData
		err = json.Unmarshal([]byte(t.Service.Data), &data)
		if err != nil {
			log.Fatal(err)
		}

		if env.ENV.Mode == env.DEVELOPMENT{
			_, err = makeHttpRequest(fmt.Sprintf(getChatMemberEndpoint, t.Key), "post", []byte(t.Service.Data), nil, t.Service.ScenarioID, &t.Service.ID)
		} else {
			_, err = makeTorifiedHttpRequest(fmt.Sprintf(getChatMemberEndpoint, t.Key), "post", []byte(t.Service.Data), t.Service.ScenarioID, &t.Service.ID)
		}

		if err != nil {
			log.Fatal(err)
		}
		break
	default:
		break
	}

	t.Next()
}

func (t Telegram) Post() {
	log.Printf("Executing type (%s) node in background : %s\n", telegramType, t.Service.Name)
}

func (t Telegram) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&t).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing \"%s\" Children \n", t.Service.Name)

	for _, node := range t.Service.Children {
		go func(n Service) {
			serv, err := n.Load()
			if err != nil {
				log.Println("error in loading services in Next()")
				return
			}

			serv.Execute()
		}(node)
	}

	log.Printf("*******End of \"%s\"*******", t.Service.Name)
}

// Service interface helper functions
func makeTorifiedHttpRequest(inputUrl string, method string, body []byte, scenarioId *uint, hId *uint) (res *http.Response, err error) {
	logData := fmt.Sprintf("Make http request")
	SendLogMessage(logData, *scenarioId, hId)

	logData = fmt.Sprintf("\nurl: %s\nmethod: %s", inputUrl, method)
	SendLogMessage(logData, *scenarioId, hId)

	SendLogMessage(string(body), *scenarioId, hId)

	tbProxyURL, err := url.Parse("socks5://tor:9050")
	if err != nil {
			log.Printf("Failed to parse proxy URL: %v\n", err)
	}

	// Get a proxy Dialer that will create the connection on our
	// behalf via the SOCKS5 proxy.  Specify the authentication
	// and re-create the dialer/transport/client if tor's
	// IsolateSOCKSAuth is needed.
	tbDialer, err := proxy.FromURL(tbProxyURL, proxy.Direct)
	if err != nil {
			log.Printf("Failed to obtain proxy dialer: %v\n", err)
	}

	// Make a http.Transport that uses the proxy dialer, and a
	// http.Client that uses the transport.
	tbTransport := &http.Transport{Dial: tbDialer.Dial}
	client := &http.Client{Transport: tbTransport}

	switch method {
	case strings.ToLower(http.MethodGet):
			res, err = client.Get(inputUrl)
			break
	case strings.ToLower(http.MethodPost):
			if body != nil {
					log.Printf("\nurl: %s\nbody:\n%s\n", inputUrl, string(body))
					reqBody := bytes.NewBuffer(body)
					res, err = client.Post(inputUrl, "application/json", reqBody)
					break
			}
			res, err = client.Post(inputUrl, "application/json", nil)
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

	logData = fmt.Sprintf("Got Respose from http request")
	SendLogMessage(logData, *scenarioId, hId)

	SendLogMessage(bodyString, *scenarioId, hId)

	logData = fmt.Sprintf("service status: %s, %v, service response body: %v", res.Status, res.StatusCode, res.Body)
	SendLogMessage(logData, *scenarioId, hId)

	return res, nil

}