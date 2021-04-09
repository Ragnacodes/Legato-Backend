package converter

import (
	"legato_server/db"
	"legato_server/api"
)

func WebhookDbToWebhookInfo(s legatoDb.Webhook) api.WebhookInfo {
	wh := api.WebhookInfo{}
	wh.WebhookUrl = s.GetURL()
	wh.Name = s.Service.Name
	wh.IsEnable = s.IsEnable

	return wh
}