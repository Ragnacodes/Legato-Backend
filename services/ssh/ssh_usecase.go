package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"
)

type SshUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewSshUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.SshUseCase {
	return &SshUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

// func (ss *SshUseCase) AddSsh(username string, ssh *api.SshInfo) (api.SshInfo, error) {
// 	user, err := ss.db.GetUserByUsername(username)
// 	if err != nil {
// 		return api.SshInfo{}, err
// 	}
// 	if err != nil {
// 		return api.SshInfo{}, err
// 	}
// 	sshDb := converter.SshInfoToSshDb(ssh)

// 	err = ss.db.CreateSshForScenario(&user, sshDb)
// 	if err != nil {
// 		return api.SshInfo{}, err
// 	}

// 	return *ssh, nil
// }
func (ss *SshUseCase) GetSshs(username string) ([]api.SshInfo, error) {
	user, err := ss.db.GetUserByUsername(username)
	if err != nil {
		return nil, err
	}

	sshs, err := ss.db.GetUserSsh(&user)
	if err != nil {
		return nil, err
	}

	var sshInfos []api.SshInfo
	sshInfos = []api.SshInfo{}
	for _, s := range sshs {
		sshInfos = append(sshInfos, converter.SshDbToSshInfo(&s))
	}

	return sshInfos, nil
}

func (ss *SshUseCase) GetSshWithId(sid uint, username string) (api.SshInfo, error) {
	user, err := ss.db.GetUserByUsername(username)
	if err != nil {
		return api.SshInfo{}, err
	}

	ssh, err := ss.db.GetSshByID(sid, &user)
	if err != nil {
		return api.SshInfo{}, err
	}

	return converter.SshDbToSshInfo(&ssh), nil

}

func (sp *SshUseCase) AddToScenario(u *api.UserInfo, scenarioId uint, ns api.NewServiceNode) (api.ServiceNode, error) {
	user, err := sp.db.GetUserByUsername(u.Username)
	if err != nil {
		return api.ServiceNode{}, err
	}

	scenario, err := sp.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return api.ServiceNode{}, err
	}

	var ss legatoDb.Ssh
	ss.Service = converter.NewServiceNodeToServiceDb(ns)

	s, err := sp.db.CreateSshForScenario(&scenario, ss)
	if err != nil {
		return api.ServiceNode{}, err
	}
	return converter.ServiceDbToServiceNode(s.Service), nil
}

func (ss *SshUseCase) Update(u *api.UserInfo, scenarioId uint, serviceId uint, ns api.NewServiceNode) error {
	user, err := ss.db.GetUserByUsername(u.Username)
	if err != nil {
		return err
	}
	scenario, err := ss.db.GetUserScenarioById(&user, scenarioId)
	if err != nil {
		return err
	}

	var ssh legatoDb.Ssh

	ssh.Service = converter.NewServiceNodeToServiceDb(ns)
	err = ss.db.UpdateSsh(&scenario, serviceId, ssh)
	if err != nil {
		return err
	}

	return nil
}
