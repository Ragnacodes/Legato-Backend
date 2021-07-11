package legatoDb

import (
	"encoding/json"
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
	MessageType 	string
	Context			string 
	ServiceLogID	uint
}

func (l *LogMessage) String() string {
	return fmt.Sprintf("(@LogMessage: %+v)", *l)
}


func (ldb *LegatoDB) GetScenarioHistories(scid uint)(historyList []History, err error){
	err = ldb.db.Model(&History{}).Where("scenario_id = ?", scid).Find(&historyList).Error
	if err != nil {
		return nil, err
	}
	return historyList, nil
}

func (ldb *LegatoDB) GetScenarioHistoriesByScenarioIds(sids []uint)(historyList []History, err error){
	err = ldb.db.Model(&History{}).
		Where("scenario_id IN ?", sids).
		Order("created_at desc").
		Find(&historyList).Error
	if err != nil {
		return nil, err
	}

	return historyList, nil
}


func (ldb *LegatoDB) GetHistoryLogs(historyID uint)(logs []ServiceLog, err error){
	err = ldb.db.Where(&ServiceLog{HistoryID: uint(historyID)}).Preload("Service").Preload("Messages").Find(&logs).Error
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

func (ldb *LegatoDB) GetHistoryById(hid uint) (history History, err error){
	err = ldb.db.Find(&history, hid).Error
	if err!=nil{
		return History{}, err
	}
	return history, nil
}

func (ldb *LegatoDB) CreateLogMessage(data string, serviceId uint, scenarioId uint) error{
	var messageType string 

	if isJSON(data){
		messageType = "json"
	} else{
		messageType = "string"
	}

	var h History
	ldb.db.Last(&h, "scenario_id = ?", scenarioId)
	var slog ServiceLog
	err := ldb.db.FirstOrCreate(&slog, ServiceLog{HistoryID: h.ID, ServiceID: serviceId}).Error
	if err != nil{
		return err
	}
	err = ldb.db.Create(&LogMessage{ServiceLogID: slog.ID, Context: data, MessageType: messageType}).Error
	if err != nil{
		return err
	}
	return err
}

// Send sse meassge and save it if service ID parameter is not nill
func SendLogMessage(message string, scId uint, serviceId *uint){
	if serviceId != nil{
		legatoDb.CreateLogMessage(message, *serviceId, scId)
	}
	logging.SSE.SendEvent(message, scId)
}


func isJSON(s string) bool {
    var js map[string]interface{}
    return json.Unmarshal([]byte(s), &js) == nil

}
