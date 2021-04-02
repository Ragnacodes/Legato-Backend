package usecase

import (
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/models"
	"time"

	uuid "github.com/satori/go.uuid"
)

type WebhookUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewWebhookUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.WebhookUseCase {
	return &WebhookUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (w *WebhookUseCase) Create(name string) (models.WebhookUrl, error){
	wh, err := w.db.CreateWebhook(name)
	if err!=nil{
		return models.WebhookUrl{}, err
	}
	return models.WebhookUrl{WebhookUrl: wh.String()}, nil 
}

func (w *WebhookUseCase) Exists(ids string) (legatoDb.Webhook, error){
	id, err := uuid.FromString(ids)
	if err!=nil{
		return legatoDb.Webhook{}, err
	}
	wh, err := w.db.GetWebhookByUUID(id)
	if err!=nil{
		return legatoDb.Webhook{}, err
	}
	return wh, nil
}

func (w *WebhookUseCase) Update(ids string, vals map[string]interface{}) error{
	id, err := uuid.FromString(ids)
	if err!=nil{
		return  err
	}
	w.db.UpdateWebhook(id, vals)
	return nil
}
