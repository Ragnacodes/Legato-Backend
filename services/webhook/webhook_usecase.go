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

func (w *WebhookUseCase) CreateNewWebhook(name string) (models.WebhookUrl, error){
	wh, err := w.db.CreateWebhook(name)
	if err!=nil{
		return models.WebhookUrl{WebhookUrl: ""}, err
	}
	return models.WebhookUrl{WebhookUrl: wh.String()}, nil 
}

func (w *WebhookUseCase) WebhookExistOr404(ids string) (bool, error){
	id, err := uuid.FromString(ids)
	if err!=nil{
		return false, err
	}
	_, err = w.db.GetWebhookByUUID(id)
	if err!=nil{
		return false, err
	}
	return true, nil
}