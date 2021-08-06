package legatoDb

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"legato_server/env"
	"log"
	"legato_server/services"
)

const discordType = "discords"

type Discord struct {
	gorm.Model
	Service Service `gorm:"polymorphic:Owner;"`
}

// Sub services
const discordSendMessage string = "sendMessage"
const discordSendMessageUrl string = "https://discord.com/api/channels/%s/messages"

type discordSendMessageData struct {
	Content string `json:"content"`
	Channel string `json:"channelId"`
}

const discordPinMessage string = "pinMessage"
const discordPinMessageUrl string = "https://discord.com/api/channels/%s/pins/%s"

type discordPinMessageData struct {
	Message string `json:"messageId"`
	Channel string `json:"channelId"`
}

const discordReactMessage string = "reactMessage"
const discordReactMessageUrl string = "https://discord.com/api/channels/%s/messages/%s/reactions/%s/@me"

type discordReactMessageData struct {
	Message string `json:"messageId"`
	Channel string `json:"channelId"`
	React   string `json:"react"`
}

func (d *Discord) String() string {
	return fmt.Sprintf("(@Discord: %+v)", *d)
}

// Database methods
func (ldb *LegatoDB) CreateDiscord(s *Scenario, discord Discord) (*Discord, error) {
	discord.Service.UserID = s.UserID
	discord.Service.ScenarioID = &s.ID

	ldb.db.Create(&discord)
	ldb.db.Save(&discord)

	return &discord, nil
}

func (ldb *LegatoDB) UpdateDiscord(s *Scenario, servId uint, nt Discord) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var t Discord
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&t).Error
	if err != nil {
		return err
	}
	if t.Service.ID != servId {
		return errors.New("the discord service is not in this scenario")
	}

	ldb.db.Model(&serv).Updates(nt.Service)
	ldb.db.Model(&t).Updates(nt)

	if nt.Service.ParentID == nil {
		legatoDb.db.Model(&serv).Select("parent_id").Update("parent_id", nil)
	}

	return nil
}

func (ldb *LegatoDB) GetDiscordByService(serv Service) (*Discord, error) {
	var t Discord
	err := ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&t).Error
	if err != nil {
		return nil, err
	}
	if t.ID != uint(serv.OwnerID) {
		return nil, errors.New("the discord service is not in this scenario")
	}

	return &t, nil
}

// Service Interface for discord
func (d Discord) Execute(Odata *services.OutputData) {
	log.Println("*******Starting Discord Service*******")

	err := legatoDb.db.Preload("Service").Find(&d).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		return
	}


	log.Printf("Executing type (%s) : %s\n", discordType, d.Service.Name)

	token := env.ENV.DiscordBotToken
	switch d.Service.SubType {
	case discordSendMessage:
		var data discordSendMessageData
		err = json.Unmarshal([]byte(d.Service.Data), &data)
		if err != nil {
			log.Println(err)
		}

		_, err = makeHttpRequest(fmt.Sprintf(discordSendMessageUrl, data.Channel), "post", []byte(d.Service.Data), &token, d.Service.ScenarioID, &d.Service.ID)
		if err != nil {
			log.Println(err)
		}
		break
	case discordPinMessage:
		var data discordPinMessageData
		err = json.Unmarshal([]byte(d.Service.Data), &data)
		if err != nil {
			log.Println(err)
		}

		_, err = makeHttpRequest(fmt.Sprintf(discordPinMessageUrl, data.Channel, data.Message), "put", nil, &token, d.Service.ScenarioID, &d.Service.ID)
		if err != nil {
			log.Println(err)
		}
		break
	case discordReactMessage:
		var data discordReactMessageData
		err = json.Unmarshal([]byte(d.Service.Data), &data)
		if err != nil {
			log.Println(err)
		}

		_, err = makeHttpRequest(fmt.Sprintf(discordReactMessageUrl, data.Channel, data.Message, data.React), "put", nil, &token, d.Service.ScenarioID, &d.Service.ID)
		if err != nil {
			log.Println(err)
		}
		break
	default:
		break
	}

	d.Next(Odata)
}

func (d Discord) Post(Odata *services.OutputData) {
	log.Printf("Executing type (%s) node in background : %s\n", discordType, d.Service.Name)
}

func (d Discord) Resume(data ...interface{}){

}

func (d Discord) Next(Odata *services.OutputData) {
	err := legatoDb.db.Preload("Service.Children").Find(&d).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		return
	}

	log.Printf("Executing \"%s\" Children \n", d.Service.Name)

	for _, node := range d.Service.Children {
		go func(n Service) {
			serv, err := n.Load()
			if err != nil {
				log.Println("error in loading services in Next()")
				return
			}

			serv.Execute(Odata)
		}(node)
	}

	log.Printf("*******End of \"%s\"*******", d.Service.Name)
}
