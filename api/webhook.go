package api

type WebhookInfo struct {
	WebhookUrl  string `json:"url"`
	Name 		string `json:"name"`
	IsEnable 	bool `json:"active"`
}

type NewWebhook struct {
}