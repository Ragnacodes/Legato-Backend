package legatoDb

import (
	"errors"
	"fmt"
	"gorm.io/gorm"
	"log"
)

const discordType = "discords"

type Discord struct {
	gorm.Model
	Service Service `gorm:"polymorphic:Owner;"`
}

func (t *Discord) String() string {
	return fmt.Sprintf("(@Discord: %+v)", *t)
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
func (t Discord) Execute(...interface{}) {
	log.Println("*******Starting Discord Service*******")

	err := legatoDb.db.Preload("Service").Find(&t).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing type (%s) : %s\n", discordType, t.Service.Name)

	switch t.Service.SubType {
	default:
		break
	}

	t.Next()
}

func (t Discord) Post() {
	log.Printf("Executing type (%s) node in background : %s\n", discordType, t.Service.Name)
}

func (t Discord) Next(...interface{}) {
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
