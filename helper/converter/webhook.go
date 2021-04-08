package converter

import (
	"legato_server/db"
	"legato_server/models"
)

func WebhookDbToWebhookInfo(s legatoDb.Webhook) models.WebhookInfo {
	wh := models.WebhookInfo{}
	wh.WebhookUrl = s.GetURL()
	wh.Name = s.Service.Name
	wh.IsEnable = s.IsEnable

	return wh
}