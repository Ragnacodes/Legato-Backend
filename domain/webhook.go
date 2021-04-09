package domain

import (
	"legato_server/api"
	"legato_server/db"
)

type WebhookUseCase interface {
	Create(userInfo *api.UserInfo, name string) api.WebhookInfo
	Exists(ids string) (*legatoDb.Webhook, error)
	Update(ids string, values map[string]interface{}) error
	List(userInfo *api.UserInfo) ([]api.WebhookInfo, error)
}
