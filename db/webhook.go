package legatoDb

import (
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"gorm.io/gorm"
	"log"
)

const webhookType string = "webhook"

type Webhook struct {
	gorm.Model
	WebhookID uuid.UUID
	IsEnable  bool    `gorm:"default:False"`
	Service   Service `gorm:"polymorphic:Owner;"`
}

func (w *Webhook) BeforeCreate(tx *gorm.DB) (err error) {
	w.WebhookID = uuid.NewV4()
	return nil
}

func (w *Webhook) String() string {
	return fmt.Sprintf("(@Webhooks: %+v)", *w)
}

func (ldb *LegatoDB) CreateWebhook(name string) *Webhook {
	wh := Webhook{Service: Service{Name: name}}
	ldb.db.Create(&wh)
	return &wh
}

func (ldb *LegatoDB) UpdateWebhook(uuid uuid.UUID, vals map[string]interface{}) error {
	var err error
	for key, value := range vals {
		if key == "name" {
			var wh Webhook
			err = ldb.db.Model(&Webhook{}).Where(&Webhook{WebhookID: uuid}).First(&wh).Error
			wh.Service.Name = value.(string)
			ldb.db.Save(&wh)
		}
		err = ldb.db.Model(&Webhook{}).Where(&Webhook{WebhookID: uuid}).Update(key, value).Error
	}
	if err != nil {
		return err
	}
	return nil
}

func (ldb *LegatoDB) GetWebhookByUUID(uuid uuid.UUID) (*Webhook, error) {
	webhook := Webhook{}
	ldb.db.Where(&Webhook{WebhookID: uuid}).First(&webhook)
	if webhook.WebhookID != uuid {
		return &Webhook{}, errors.New("webhook obj does not exist")
	}
	return &webhook, nil
}

// Service Interface for Webhook

func (w *Webhook) Execute(...interface{}) {
	err := legatoDb.db.Preload("Service").Find(&w).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing %s node: %s\n", "webhook", w.Service.Name)
	w.IsEnable = true
	legatoDb.db.Save(&w)

	w.Next()
}

func (w *Webhook) Post() {
	log.Printf("Executing %s node in background: %s\n", "webhook", w.Service.Name)
}

func (w *Webhook) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&w).Error
	if err != nil {
		panic(err)
	}

	for _, node := range w.Service.Children {
		node.LoadOwner().Execute()
	}
}