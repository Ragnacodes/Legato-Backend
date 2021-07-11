package usecase

import (
	"legato_server/api"
	legatoDb "legato_server/db"
	"legato_server/domain"
	"legato_server/helper/converter"
	"time"
)

type loggerUseCase struct {
	db             *legatoDb.LegatoDB
	contextTimeout time.Duration
}

func NewLoggerUseCase(db *legatoDb.LegatoDB, timeout time.Duration) domain.LoggerUseCase {
	return &loggerUseCase{
		db:             db,
		contextTimeout: timeout,
	}
}

func (s loggerUseCase) SaveMessage(data string, servicelogId uint, historyID uint) error {

	err := s.db.CreateLogMessage(data, servicelogId, historyID)
	if err != nil {
		return err
	}
	return nil
}

func (s loggerUseCase) GetRecentUserLogsWithScenario(u *api.UserInfo) (historyList []api.HistoryInfo, err error) {
	user := converter.UserInfoToUserDb(*u)
	scenarios, err := s.db.GetUserScenarios(&user)
	if err != nil {
		return nil, err
	}

	// The ids of the scenarios
	var sids []uint
	sids = []uint{}
	for _, scenario := range scenarios {
		sids = append(sids, scenario.ID)
	}

	histories, err := s.db.GetScenarioHistoriesByScenarioIds(sids)
	if err != nil {
		return nil, err
	}

	for _, h := range histories {
		briefHistory := converter.HistoryDbToHistoryInfos(h)
		historyList = append(historyList, briefHistory)
	}

	return historyList, nil
}

func (s loggerUseCase) GetScenarioHistoriesById(scid uint) (historyList []api.HistoryInfo, err error) {

	histories, err := s.db.GetScenarioHistories(scid)
	if err != nil {
		return nil, err
	}

	for _, h := range histories {
		briefHistory := converter.HistoryDbToHistoryInfos(h)
		historyList = append(historyList, briefHistory)
	}

	return historyList, nil
}

func (s loggerUseCase) GetHistoryLogsById(historyId uint) (serviceLogs []api.ServiceLogInfo, err error) {

	logs, err := s.db.GetHistoryLogs(historyId)
	if err != nil {
		return nil, err
	}

	for _, l := range logs {
		log := converter.ServiceLogDbToServiceLogInfos(l)
		serviceLogs = append(serviceLogs, log)
	}

	return serviceLogs, nil
}

func (s loggerUseCase) GetHistoryById(hid uint) (history api.HistoryInfo, err error) {

	historydb, err := s.db.GetHistoryById(hid)
	if err != nil {
		return api.HistoryInfo{}, err
	}
	history = converter.HistoryDbToHistoryInfos(historydb)

	return history, nil
}
