package models

type WebhookUrl struct {
	WebhookUrl  string `json:"url"`
}

type NewWebhook struct {
	Name  string `json:"name"`
}