package converter

import (
	"encoding/json"
	"legato_server/api"
	legatoDb "legato_server/db"
)

func WebhookDbToWebhookInfo(s legatoDb.Webhook) api.WebhookInfo {
	wh := api.WebhookInfo{}
	wh.Id = s.ID
	wh.WebhookUrl = s.GetURL()
	wh.Name = s.Service.Name
	wh.IsEnable = s.IsEnable
	wh.GetHeaders = s.GetHeaders
	wh.GetMethod = s.GetMethod

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
	data := &map[string]interface{}{
		"webhook": &map[string]interface{}{
			"url":      wh.GetURL(),
			"isEnable": wh.IsEnable,
			"id":       wh.ID,
			"getMethod": wh.GetMethod,
			"getHeaders": wh.GetHeaders,
			"name": wh.Service.Name,
		},
	}
	jsonString, _ := json.Marshal(data)
	_ = json.Unmarshal(jsonString, sn.Data)
	
	return sn
}

func NewSeparateWebhookToWebhook(s api.NewSeparateWebhook) legatoDb.Webhook {
	var wh legatoDb.Webhook
	wh.Service.Name = s.Name
	wh.GetHeaders = s.GetHeaders
	wh.GetMethod = s.GetMethod
	return wh
}
