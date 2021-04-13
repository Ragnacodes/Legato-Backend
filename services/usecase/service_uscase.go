package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"
)

type ServiceUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewServiceUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.ServiceUseCase {
	return &ServiceUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (s *ServiceUseCase) GetServiceNodeById(u *api.UserInfo, scenarioId uint, serviceId uint) (api.ServiceNode, error) {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	serv, err := s.db.GetServiceById(&scenario, serviceId)
	if err != nil {
		return api.ServiceNode{}, nil
	}

	return converter.ServiceDbToServiceNode(*serv), nil
}

func (s *ServiceUseCase) DeleteServiceNodeById(u *api.UserInfo, scenarioId uint, serviceId uint) error {
	user := converter.UserInfoToUserDb(*u)
	scenario, err := s.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	err = s.db.DeleteServiceById(&scenario, serviceId)
	if err != nil {
		return err
	}

	return nil
}
