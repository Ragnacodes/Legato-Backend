package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"

	uuid "github.com/satori/go.uuid"
)

type WebhookUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewWebhookUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.WebhookUseCase {
	return &WebhookUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (w *WebhookUseCase) Create(u *api.UserInfo, scenarioId uint, nw api.NewServiceNode) (api.ServiceNode, error) {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	webhook := converter.DataToWebhook(nw.Data)
	webhook.Service = converter.NewServiceNodeToServiceDb(nw)

	wh, err := w.db.CreateWebhook(&user, scenarioId, webhook)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.WebhookDbToServiceNode(*wh), nil
}

func (w *WebhookUseCase) Exists(ids string) (*legatoDb.Webhook, error){
	id, err := uuid.FromString(ids)
	if err!=nil{
		return &legatoDb.Webhook{}, err
	}
	wh, err := w.db.GetWebhookByUUID(id)
	if err!=nil{
		return &legatoDb.Webhook{}, err
	}
	return wh, nil
}

func (w *WebhookUseCase) Update(ids string, vals map[string]interface{}) error{
	id, err := uuid.FromString(ids)
	if err!=nil{
		return  err
	}
	err = w.db.UpdateWebhook(id, vals)
	if err!=nil{
		return  err
	}
	return nil
}

func (w *WebhookUseCase) List(u *api.UserInfo) ([]api.WebhookInfo, error){
	user := converter.UserInfoToUserDb(*u)
	webhooks, err := w.db.GetUserWebhooks(&user)
	if err != nil {
		return nil, err
	}

	var WebhookInfos []api.WebhookInfo
	for _, w := range webhooks {
		WebhookInfos = append(WebhookInfos, converter.WebhookDbToWebhookInfo(w))
	}

	return WebhookInfos, nil
}