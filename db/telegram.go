package legatoDb

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
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
	ChatId string   `json:"chat_id"`
	Text   string `json:"text"`
}

type getChatMemberData struct {
	ChatId string   `json:"chat_id"`
	UserId   string `json:"user_id"`
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

		_, _ = makeHttpRequest(fmt.Sprintf(sendMessageEndpoint, t.Key), "post", []byte(t.Service.Data))
		break
	case getChatMember:
		var data getChatMemberData
		err = json.Unmarshal([]byte(t.Service.Data), &data)
		if err != nil {
			log.Fatal(err)
		}

		_, _ = makeHttpRequest(fmt.Sprintf(getChatMemberEndpoint, t.Key), "post", []byte(t.Service.Data))
		break
	default:
		break
	}

	t.Next()
}

func (t Telegram) Post() {
	log.Printf("Executing type (%s) node in background : %s\n", httpType, t.Service.Name)
}

func (t Telegram) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&t).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing \"%s\" Children \n", t.Service.Name)

	for _, node := range t.Service.Children {
		serv, err := node.Load()
		if err != nil {
			log.Println("error in loading services in Next()")
			return
		}
		serv.Execute()
	}

	log.Printf("*******End of \"%s\"*******", t.Service.Name)
}
