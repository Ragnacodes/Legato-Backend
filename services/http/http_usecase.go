package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"
)

type HttpUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewHttpUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.HttpUseCase {
	return &HttpUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (w *HttpUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, nh api.NewServiceNode) (api.ServiceNode, error) {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := w.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	http := converter.DataToHttp(nh.Data)
	http.Service = converter.NewServiceNodeToServiceDb(nh)

	h, err := w.db.CreateHttp(&scenario, http)
	if err != nil {
		return api.ServiceNode{}, err
	}

	return converter.HttpDbToServiceNode(*h), nil
}

func (w *HttpUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, nh api.NewServiceNode) error {
	user, err := w.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}

	scenario, err := w.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	http := converter.DataToHttp(nh.Data)
	http.Service = converter.NewServiceNodeToServiceDb(nh)

	err = w.db.UpdateHttp(&scenario, serviceId, http)
	if err != nil {
		return err
	}

	return nil
}
