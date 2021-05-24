package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"
)

type DiscordUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewDiscordUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.DiscordUseCase {
	return &DiscordUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (du DiscordUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, nh api.NewServiceNode) (api.ServiceNode, error) {
	user, err := du.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := du.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	var discord legatoDb.Discord
	discord.Service = converter.NewServiceNodeToServiceDb(nh)

	h, err := du.db.CreateDiscord(&scenario, discord)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.ServiceDbToServiceNode(h.Service), nil
}

func (du DiscordUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, nd api.NewServiceNode) error {
	user, err := du.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	scenario, err := du.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	var discord legatoDb.Discord
	discord.Service = converter.NewServiceNodeToServiceDb(nd)

	err = du.db.UpdateDiscord(&scenario, serviceId, discord)
	if err != nil {
		return err
	}

	return nil
}
