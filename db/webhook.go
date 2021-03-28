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
	Service
	WebhookID uuid.UUID
}

func (w *Webhook) BeforeCreate(tx *gorm.DB) (err error){
	w.WebhookID = uuid.NewV4()
	w.Type = Type
	if err != nil {
		return err
	}
	return
}

func (w *Webhook) String() string {
	absolute_path := env.ENV.WebHost + ":" + env.ENV.ServingPort
	return fmt.Sprintf("%s/services/webhook/%v", absolute_path, w.WebhookID)
}


func (ldb *LegatoDB) CreateWebhook(name string) (Webhook, error) {
	wh:= Webhook{}
	wh.Name = name
	ldb.db.Create(&wh)
	return wh, nil
}


func (ldb *LegatoDB) GetWebhookByUUID(uuid uuid.UUID) (Webhook, error) {
	webhook := Webhook{}
	ldb.db.Where(&Webhook{WebhookID: uuid}).First(&webhook)
	if webhook.WebhookID != uuid {
		return Webhook{}, errors.New("Webhook obj does not exist")
	}

	return webhook, nil
}