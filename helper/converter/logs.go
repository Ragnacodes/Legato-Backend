package converter

import (
	"legato_server/api"
	legatoDb "legato_server/db"
)


func HistoryDbToHistoryInfos(hdb legatoDb.History) (hInfo api.HistoryInfo){
	hInfo.CreatedAt = hdb.CreatedAt.Format("2006-01-02T15:04:05-0700")
	hInfo.ID = hdb.ID
	return hInfo
}

func ServiceLogDbToServiceLogInfos(dbServiceLog legatoDb.ServiceLog) (logInfo api.ServiceLogInfo){
	logInfo.Messages = MessageDbToMessageInfo(dbServiceLog.Messages) 
	logInfo.Id = int(dbServiceLog.ID)
	logInfo.Service = ServiceDbToServiceNode(dbServiceLog.Service)
	return logInfo
}

func MessageDbToMessageInfo(dbMesaages []*legatoDb.LogMessage) (messageInfos []api.MessageInfo){
	for _, m := range dbMesaages{
		var message api.MessageInfo
		message.Data = m.Context
		messageInfos = append(messageInfos, message)
	}
	return messageInfos
}