package domain

import (
	"legato_server/models"
	"legato_server/db"
)

type WebhookUseCase interface {
	Create(userInfo *models.UserInfo, name string) (models.WebhookInfo)
	Exists(ids string) (*legatoDb.Webhook, error)
	Update(ids string, vals map[string]interface{}) error
	List(userInfo *models.UserInfo) ([]models.WebhookInfo, error)
}
