package converter

import (
	"legato_server/api"
	legatoDb "legato_server/db"
)

func WebhookDbToWebhookInfo(s legatoDb.Webhook) api.WebhookInfo {
	wh := api.WebhookInfo{}
	wh.Id = s.ID
	wh.WebhookUrl = s.GetURL()
	wh.Name = s.Service.Name
	wh.IsEnable = s.IsEnable

	return wh
}

func WebhookToWebhookDb(s api.WebhookInfo) legatoDb.Webhook {
	var wh legatoDb.Webhook
	wh.Service.Name = s.Name
	wh.IsEnable = s.IsEnable

	return wh
}

func DataToWebhookDb(data interface{}) legatoDb.Webhook {
	var w legatoDb.Webhook
	if data != nil {
		_ = data.(map[string]interface{})
	}

	return w
}

func WebhookDbToServiceNode(wh legatoDb.Webhook) api.ServiceNode {
	var sn api.ServiceNode
	sn = ServiceDbToServiceNode(wh.Service)
	// Webhook data
	sn.Data = WebhookDbToWebhookInfo(wh)

	return sn
}

func NewSeparateWebhookToWebhook(s api.NewSeparateWebhook) legatoDb.Webhook {
	var wh legatoDb.Webhook
	wh.Service.Name = s.Name
	wh.GetHeaders = s.GetHeaders
	wh.GetMethod = s.GetMethod
	return wh
}
