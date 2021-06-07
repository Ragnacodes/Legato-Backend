package api

type WebhookInfo struct {
	Id         uint   `json:"id"`
	GetMethod  bool	  `json:"getMethod"`
	GetHeaders bool   `json:"getHeaders"`
	WebhookUrl string `json:"url"`
	Name       string `json:"name"`
	IsEnable   bool   `json:"isEnable"`
}

type NewSeparateWebhook struct {
	Name       string `json:"name"`
	IsEnable   bool   `json:"isEnable"`
	GetMethod  *bool	  `json:"getMethod"`
	GetHeaders *bool   `json:"getHeaders"`
}
