package legatoDb

import (
	"errors"
	"fmt"
	"legato_server/logging"
	"gorm.io/gorm"
)


type History struct{
	gorm.Model
	LogMessage []*ServiceLog
	ScenarioID uint
}

type ServiceLog struct {
	gorm.Model
	Status	   int
	Messages   []*LogMessage
	HistoryID  uint
	Service	   Service
	ServiceID  uint
}


type LogMessage struct{
	gorm.Model
	Context			string 
	ServiceLogID	uint
}

func (l *LogMessage) String() string {
	return fmt.Sprintf("(@LogMessage: %+v)", *l)
}


func (ldb *LegatoDB) GetScenarioHistories(scid uint)(historyList []History, er error){
	err := ldb.db.Model(&Scenario{}).Where("id = ?", scid).Association("Histories").Find(&historyList).Error()
	if err != "" {
		return nil, errors.New(err)
	}
	return historyList, nil

}


func (ldb *LegatoDB) GetHistoryLogs(historyID uint)(logs []ServiceLog, err error){
	err = ldb.db.Where(&ServiceLog{HistoryID: uint(historyID)}).Preload("Messages").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}

func (ldb *LegatoDB) CreateHistory(scenarioId uint) error{
	err := ldb.db.Create(&History{ScenarioID: scenarioId}).Error
	if err!=nil{
		return err
	}
	return nil
}

func (ldb *LegatoDB) CreateLogMessage(data string, servicelogId uint, historyID uint) error{
	var slog ServiceLog
	err := ldb.db.FirstOrCreate(&slog, ServiceLog{HistoryID: historyID, ServiceID:  servicelogId}).Error
	if err != nil{
		return err
	}
	err = ldb.db.Create(&LogMessage{ServiceLogID: slog.ID,Context: data}).Error
	if err != nil{
		return err
	}
	return err
}

// Send sse meassge and save it if service ID parameter is not nill
func SendLogMessage(message string, scId uint, serviceId *uint){
	if serviceId!=nil{
		legatoDb.CreateLogMessage(message, scId, *serviceId)
	}
	logging.SSE.SendEvent(message, scId)
}