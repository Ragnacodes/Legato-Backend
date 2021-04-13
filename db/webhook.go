package legatoDb

import (
	"errors"
	"fmt"
	"legato_server/env"
	"log"

	"github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const webhookType string = "webhook"

type Webhook struct {
	gorm.Model
	Token    uuid.UUID
	IsEnable bool    `gorm:"default:False"`
	Service  Service `gorm:"polymorphic:Owner;"`
}

func (w *Webhook) String() string {
	return fmt.Sprintf("(@Webhooks: %+v)", *w)
}

func (w *Webhook) BeforeCreate(tx *gorm.DB) (err error) {
	w.Token = uuid.NewV4()
	return nil
}

func (w *Webhook) GetURL() string {
	return fmt.Sprintf("%s:%s/api/services/webhook/%v", env.ENV.WebHost, env.ENV.ServingPort, w.Token)
}

func (ldb *LegatoDB) CreateWebhookForScenario(s *Scenario, wh Webhook) (*Webhook, error) {
	wh.Service.UserID = s.UserID
	wh.Service.ScenarioID = &s.ID

	ldb.db.Create(&wh)
	ldb.db.Save(&wh)

	return &wh, nil
}

func (ldb *LegatoDB) CreateSeparateWebhook(u *User, wh Webhook) (*Webhook, error) {
	wh.Service.UserID = u.ID
	wh.Service.ScenarioID = nil

	ldb.db.Create(&wh)
	ldb.db.Save(&wh)

	return &wh, nil
}

func (ldb *LegatoDB) CreateWebhookInScenario(u *User, s *Scenario, parent *Service, name string, x int, y int) *Webhook {
	var wh Webhook
	if parent != nil {
		wh = Webhook{Service: Service{Name: name, UserID: u.ID, ScenarioID: &s.ID, ParentID: &parent.ID, PosX: x, PosY: y}}
	} else {
		wh = Webhook{Service: Service{Name: name, UserID: u.ID, ScenarioID: &s.ID, PosX: x, PosY: y}}
	}
	ldb.db.Create(&wh)
	ldb.db.Save(&wh)
	return &wh
}

func (ldb *LegatoDB) UpdateWebhook(s *Scenario, servId uint, nwh Webhook) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var wh Webhook
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&wh).Error
	if err != nil {
		return err
	}
	if wh.Service.ID != servId {
		return errors.New("the webhook service is not in this scenario")
	}

	ldb.db.Model(&serv).Updates(nwh.Service)
	ldb.db.Model(&wh).Updates(nwh)

	return nil
}

func (ldb *LegatoDB) GetWebhookByUUID(uuid uuid.UUID) (*Webhook, error) {
	webhook := Webhook{}
	ldb.db.Where(&Webhook{Token: uuid}).Preload("Service").First(&webhook)
	if webhook.Token != uuid {
		return &Webhook{}, errors.New("webhook obj does not exist")
	}
	return &webhook, nil
}

func (ldb *LegatoDB) GetUserWebhooks(u *User) ([]Webhook, error) {
	var services []Service
	err := ldb.db.Select("id").Where(&Service{UserID: u.ID}).Find(&services).Error
	if err != nil {
		return nil, err
	}

	// Collect webhook service id
	var serviceIds []uint
	serviceIds = []uint{}
	for _, srv := range services {
		serviceIds = append(serviceIds, srv.ID)
	}

	var webhooks []Webhook
	err = ldb.db.Where(serviceIds).Preload("Service").Find(&webhooks).Error
	if err != nil {
		return nil, err
	}

	return webhooks, nil
}

// Service Interface for Webhook

func (w Webhook) Execute(...interface{}) {
	err := legatoDb.db.Preload("Service").Find(&w).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing %s node: %s\n", "webhook", w.Service.Name)

	w.IsEnable = true
	legatoDb.db.Save(&w)

	w.Next()
}

func (w Webhook) Post() {
	log.Printf("Executing %s node in background: %s\n", "webhook", w.Service.Name)
}

func (w Webhook) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&w).Error
	if err != nil {
		panic(err)
	}

	for _, node := range w.Service.Children {
		node.LoadOwner().Execute()
	}
}
