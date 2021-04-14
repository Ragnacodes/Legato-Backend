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

func (w *WebhookUseCase) AddWebhookToScenario(u *api.UserInfo, scenarioId uint, nw api.NewServiceNode) (api.ServiceNode, error) {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := w.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	webhook := converter.DataToWebhook(nw.Data)
	webhook.Service = converter.NewServiceNodeToServiceDb(nw)

	wh, err := w.db.CreateWebhookForScenario(&scenario, webhook)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.WebhookDbToServiceNode(*wh), nil
}

func (w *WebhookUseCase) CreateSeparateWebhook(u *api.UserInfo, nw api.NewSeparateWebhook) (api.WebhookInfo, error) {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.WebhookInfo{}, err
	}

	webhook := converter.NewSeparateWebhookToWebhook(nw)

	wh, err := w.db.CreateSeparateWebhook(&user, webhook)
	if err != nil {
		return api.WebhookInfo{}, err
	}

	return converter.WebhookDbToWebhookInfo(*wh), nil
}

func (w *WebhookUseCase) Exists(ids string) (*legatoDb.Webhook, error) {
	id, err := uuid.FromString(ids)
	if err != nil {
		return &legatoDb.Webhook{}, err
	}
	wh, err := w.db.GetWebhookByUUID(id)
	if err != nil {
		return &legatoDb.Webhook{}, err
	}
	return wh, nil
}

func (w *WebhookUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, nw api.NewServiceNode) error {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	scenario, err := w.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	webhook := converter.DataToWebhook(nw.Data)
	webhook.Service = converter.NewServiceNodeToServiceDb(nw)

	err = w.db.UpdateWebhook(&scenario, serviceId, webhook)
	if err != nil {
		return err
	}

	return nil
}

func (w *WebhookUseCase) UpdateSeparateWebhook(u *api.UserInfo, wid uint, nw api.NewSeparateWebhook) error {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	nwh := converter.NewSeparateWebhookToWebhook(nw)

	err = w.db.UpdateSeparateWebhook(&user, wid, nwh)
	if err != nil {
		return err
	}

	return nil
}

func (w *WebhookUseCase) GetUserWebhooks(u *api.UserInfo) ([]api.WebhookInfo, error) {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return nil, err
	}

	webhooks, err := w.db.GetUserWebhooks(&user)
	if err != nil {
		return nil, err
	}

	var whInfos []api.WebhookInfo
	whInfos = []api.WebhookInfo{}
	for _, w := range webhooks {
		whInfos = append(whInfos, converter.WebhookDbToWebhookInfo(w))
	}

	return whInfos, nil
}

func (w *WebhookUseCase) GetUserWebhookById(u *api.UserInfo, wid uint) (api.WebhookInfo, error) {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.WebhookInfo{}, err
	}

	wh, err := w.db.GetUserWebhookById(&user, wid)
	if err != nil {
		return api.WebhookInfo{}, err
	}

	return converter.WebhookDbToWebhookInfo(wh), nil
}

func (w *WebhookUseCase) DeleteUserWebhookById(u *api.UserInfo, wid uint) error {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	err = w.db.DeleteSeparateWebhookById(&user, wid)
	if err != nil {
		return  err
	}

	return nil
}
