package domain

import (
	"legato_server/api"
	"legato_server/db"
)

type WebhookUseCase interface {
	Create(userInfo *api.UserInfo, scenarioId uint, newWebhook api.NewServiceNode) (api.ServiceNode, error)
	Exists(ids string) (*legatoDb.Webhook, error)
	Update(u *api.UserInfo, scenarioId uint, nodeId uint, nw api.NewServiceNode) error
	List(userInfo *api.UserInfo) ([]api.WebhookInfo, error)
}
