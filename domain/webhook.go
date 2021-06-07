package domain

import (
	"legato_server/api"
)

type WebhookUseCase interface {
	AddWebhookToScenario(userInfo *api.UserInfo, scenarioId uint, newWebhook api.NewServiceNode) (api.ServiceNode, error)
	CreateSeparateWebhook(userInfo *api.UserInfo, newWebhook api.NewSeparateWebhook) (api.WebhookInfo, error)
	Exists(ids string) (*api.WebhookInfo, error)
	Update(u *api.UserInfo, scenarioId uint, nodeId uint, nw api.NewServiceNode) error
	UpdateSeparateWebhook(u *api.UserInfo, nodeId uint, nw api.NewSeparateWebhook) error
	GetUserWebhooks(userInfo *api.UserInfo) ([]api.WebhookInfo, error)
	GetUserWebhookById(u *api.UserInfo, wid uint) (api.WebhookInfo, error)
	DeleteUserWebhookById(u *api.UserInfo, wid uint) error
	TriggerWebhook(wid string, data map[string]interface{}) error
	GetUserWebhookHistoryById(u *api.UserInfo, webhookId uint)([]api.ServiceLogInfo , error)
}
