package domain

import (
	"legato_server/api"
	"legato_server/db"
)

type WebhookUseCase interface {
	AddWebhookToScenario(userInfo *api.UserInfo, scenarioId uint, newWebhook api.NewServiceNode) (api.ServiceNode, error)
	CreateSeparateWebhook(userInfo *api.UserInfo, newWebhook api.NewSeparateWebhook) (api.WebhookInfo, error)
	Exists(ids string) (*legatoDb.Webhook, error)
	Update(u *api.UserInfo, scenarioId uint, nodeId uint, nw api.NewServiceNode) error
	UpdateSeparateWebhook(u *api.UserInfo, nodeId uint, nw api.NewSeparateWebhook) error
	GetUserWebhooks(userInfo *api.UserInfo) ([]api.WebhookInfo, error)
	GetUserWebhookById(u *api.UserInfo, wid uint) (api.WebhookInfo, error)
	DeleteUserWebhookById(u *api.UserInfo, wid uint) error
}
