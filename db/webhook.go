package legatoDb

import (
	"errors"
	"fmt"
	"legato_server/env"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
	"log"
)

const Type string = "webhook"

type Webhook struct {
	gorm.Model
	WebhookID uuid.UUID
	Enable bool `gorm:"default:False"`
	Service Service `gorm:"polymorphic:Owner;"`
}

func (w *Webhook) BeforeCreate(tx *gorm.DB) (err error){
	w.WebhookID = uuid.NewV4()
	return nil
}

func (w *Webhook) String() string {
	return fmt.Sprintf("%s:%s/api/services/webhook/%v", env.ENV.WebHost, env.ENV.ServingPort, w.WebhookID)
}


func (ldb *LegatoDB) CreateWebhook(name string) *Webhook {
	wh := Webhook{Service:  Service{Name : name}}
	ldb.Db.Create(&wh)
	return &wh
}


func (ldb *LegatoDB) UpdateWebhook(uuid uuid.UUID, vals map[string]interface{}) error{
	var err error 
	for key, value := range vals{
		if key == "name" {
			var wh Webhook
			err = ldb.Db.Model(&Webhook{}).Where(&Webhook{WebhookID: uuid}).First(&wh).Error
			wh.Service.Name = value.(string)
			ldb.Db.Save(&wh)
		}
		err = ldb.Db.Model(&Webhook{}).Where(&Webhook{WebhookID: uuid}).Update(key, value).Error
	}
	if err!=nil{
		return err
	}
	return nil
} 

func (ldb *LegatoDB) GetWebhookByUUID(uuid uuid.UUID) (*Webhook, error) {
	webhook := Webhook{}
	ldb.Db.Where(&Webhook{WebhookID: uuid}).First(&webhook)
	if webhook.WebhookID != uuid {
		return &Webhook{}, errors.New("webhook obj does not exist")
	}
	return &webhook, nil
}

//implement Service Interface for Webhook

func (w *Webhook) Execute(attrs ...interface{}) {
	log.Printf("Executing %s node: %s\n", "webhook", w.Service.Name)
	w.Enable = true
	legatoDb.Db.Save(&w)
	w.Post()
}

func (w *Webhook) Post() {
	log.Printf("Executing %s node in background: %s\n", "webhook", w.Service.Name)
}

func (w *Webhook) Next(attrs ...interface{}) {
	///Next function is not working ???!!
	for _, node := range w.Service.Children {
		node.LoadOwner().Execute()
	}
}
