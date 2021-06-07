package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"
)

type ToolBoxUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewToolBoxUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.ToolBoxUseCase {
	return &ToolBoxUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (tu ToolBoxUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, nt api.NewServiceNode) (api.ServiceNode, error) {
	user, err := tu.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := tu.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	var toolBox legatoDb.ToolBox
	toolBox.Service = converter.NewServiceNodeToServiceDb(nt)

	h, err := tu.db.CreateToolBox(&scenario, toolBox)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.ServiceDbToServiceNode(h.Service), nil
}

func (tu ToolBoxUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, nt api.NewServiceNode) error {
	user, err := tu.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	scenario, err := tu.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	var toolBox legatoDb.ToolBox
	toolBox.Service = converter.NewServiceNodeToServiceDb(nt)

	err = tu.db.UpdateToolBox(&scenario, serviceId, toolBox)
	if err != nil {
		return err
	}

	return nil
}
