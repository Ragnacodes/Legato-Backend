package domain

import (
	"legato_server/models"
	"legato_server/db"
)

type WebhookUseCase interface {
	Create(name string) (models.WebhookUrl, error)
	Exists(ids string) (legatoDb.Webhook, error)
	Update(ids string, vals map[string]interface{}) error
}
