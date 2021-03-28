package domain

import (
	"legato_server/models"
)

type WebhookUseCase interface {
	CreateNewWebhook(name string) (models.WebhookUrl, error)
	WebhookExistOr404(ids string) (bool, error)
}
