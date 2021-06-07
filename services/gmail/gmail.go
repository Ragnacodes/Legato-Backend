package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"
)

type GmailUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewGmailUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.GmailUseCase {
	return &GmailUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}
func (gu *GmailUseCase) GetGmailWithId(gid uint, username string) (api.GmailInfo, error) {
	user, err := gu.db.GetUserByUsername(username)
	if err != nil {
		return api.GmailInfo{}, err
	}

	git, err := gu.db.GetGmailByID(gid, &user)
	if err != nil {
		return api.GmailInfo{}, err
	}

	return converter.GmailDbToGitInfo(&git), nil
}

func (gu *GmailUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, ns api.NewServiceNode) (api.ServiceNode, error) {

	user, err := gu.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := gu.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	var gg legatoDb.Gmail
	gg.Service = converter.NewServiceNodeToServiceDb(ns)
	g, err := gu.db.CreateGmailForScenario(&scenario, gg)
	if err != nil {
		return api.ServiceNode{}, err
	}
	return converter.ServiceDbToServiceNode(g.Service), nil
}

func (gu *GmailUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, ns api.NewServiceNode) error {
	user, err := gu.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}
	scenario, err := gu.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	var g legatoDb.Gmail
	g.Service = converter.NewServiceNodeToServiceDb(ns)
	err = gu.db.UpdateGmail(&scenario, serviceId, g)
	if err != nil {
		return err
	}

	return nil
}
