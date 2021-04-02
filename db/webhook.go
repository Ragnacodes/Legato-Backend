package legatoDb

import (
	"errors"
	"fmt"
	"legato_server/env"
	"github.com/jinzhu/gorm"
	"github.com/satori/go.uuid"
)

const Type string = "webhook"

type Webhook struct {
	Service Service `gorm:"foreignKey:ID"`
	WebhookID uuid.UUID
	Enable bool `gorm:"default:False"`
}

func (w *Webhook) BeforeCreate(tx *gorm.DB) (err error){
	w.WebhookID = uuid.NewV4()
	return nil
}

func (w *Webhook) String() string {
	return fmt.Sprintf("%s:%s/api/services/webhook/%v", env.ENV.WebHost, env.ENV.ServingPort, w.WebhookID)
}


func (ldb *LegatoDB) CreateWebhook(name string) (Webhook, error) {
	sr := Service{Name : name, Type : Type}
	wh := Webhook{Service: sr}
	ldb.db.Create(&wh)
	return wh, nil
}


func (ldb *LegatoDB) UpdateWebhook(uuid uuid.UUID, vals map[string]interface{}) {
	for key, value := range vals{
		if key == "name" {
			var wh Webhook
			ldb.db.Model(&Webhook{}).Where("WebhookID = ?", uuid).First(&wh)
			wh.Service.Name = value.(string)
			ldb.db.Save(&wh)
		}
		ldb.db.Model(&Webhook{}).Where("WebhookID = ?", uuid).Update(key, value)
	}
} 

func (ldb *LegatoDB) GetWebhookByUUID(uuid uuid.UUID) (Webhook, error) {
	webhook := Webhook{}
	ldb.db.Where(&Webhook{WebhookID: uuid}).First(&webhook)
	if webhook.WebhookID != uuid {
		return Webhook{}, errors.New("webhook obj does not exist")
	}

	return webhook, nil
}
