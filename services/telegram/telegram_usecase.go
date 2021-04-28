package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"
)

type TelegramUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewTelegramUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.TelegramUseCase {
	return &TelegramUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (tu *TelegramUseCase) Test() {

}

func (tu *TelegramUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, nt api.NewServiceNode) (api.ServiceNode, error) {
	user, err := tu.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := tu.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	var telegram legatoDb.Telegram
	telegram.Service = converter.NewServiceNodeToServiceDb(nt)
	if key, ok := nt.Data.(map[string]interface{})["key"]; ok {
		telegram.Key = key.(string)
	}

	h, err := tu.db.CreateTelegram(&scenario, telegram)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.ServiceDbToServiceNode(h.Service), nil
}

func (tu *TelegramUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, nt api.NewServiceNode) error {
	user, err := tu.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	scenario, err := tu.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	var telegram legatoDb.Telegram
	telegram.Service = converter.NewServiceNodeToServiceDb(nt)
	if key, ok := nt.Data.(map[string]interface{})["key"]; ok {
		telegram.Key = key.(string)
	}

	err = tu.db.UpdateTelegram(&scenario, serviceId, telegram)
	if err != nil {
		return err
	}

	return nil
}
