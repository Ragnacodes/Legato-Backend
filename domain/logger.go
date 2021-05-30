package domain

import (
	"legato_server/api"
)

type LoggerUseCase interface {
	GetScenarioHistoriesById(scid uint) (historyList []api.HistoryInfo, err error)
	GetHistoryById(hid uint) (history api.HistoryInfo, err error)
	GetHistoryLogsById(historyId uint) (serviceLogList []api.ServiceLogInfo, err error)
	SaveMessage(data string, servicelogId uint, historyID uint) (error)
}
