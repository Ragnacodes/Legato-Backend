package api

type WebhookInfo struct {
	WebhookUrl string `json:"url"`
	Name       string `json:"name"`
	IsEnable   bool   `json:"isEnable"`
}

type NewSeparateWebhook struct {
	Name string `json:"name"`
	IsEnable   bool   `json:"isEnable"`
}
