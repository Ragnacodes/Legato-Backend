package legatoDb

import (
	"encoding/json"
	"errors"
	"fmt"
	"gorm.io/gorm"
	"legato_server/services"
	"log"
	"time"
)

const toolBoxType = "tool_boxes"

// Sub services
const toolBoxSleep string = "sleep"
const toolBoxRepeater string = "repeater"

type ToolBox struct {
	gorm.Model
	Service Service `gorm:"polymorphic:Owner;"`
}

type toolBoxSleepData struct {
	Time int32 `json:"time"`
}

type toolBoxRepeaterData struct {
	Count int `json:"count"`
}

func (t *ToolBox) String() string {
	return fmt.Sprintf("(@Toolbox: %+v)", *t)
}

// Database methods
func (ldb *LegatoDB) CreateToolBox(s *Scenario, toolBox ToolBox) (*ToolBox, error) {
	toolBox.Service.UserID = s.UserID
	toolBox.Service.ScenarioID = &s.ID

	ldb.db.Create(&toolBox)
	ldb.db.Save(&toolBox)

	return &toolBox, nil
}

func (ldb *LegatoDB) UpdateToolBox(s *Scenario, servId uint, nt ToolBox) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var t ToolBox
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&t).Error
	if err != nil {
		return err
	}
	if t.Service.ID != servId {
		return errors.New("the toolbox service is not in this scenario")
	}

	ldb.db.Model(&serv).Updates(nt.Service)
	ldb.db.Model(&t).Updates(nt)

	if nt.Service.ParentID == nil {
		legatoDb.db.Model(&serv).Select("parent_id").Update("parent_id", nil)
	}

	return nil
}

func (ldb *LegatoDB) GetToolBoxByService(serv Service) (*ToolBox, error) {
	var t ToolBox
	err := ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&t).Error
	if err != nil {
		return nil, err
	}
	if t.ID != uint(serv.OwnerID) {
		return nil, errors.New("the toolbox service is not in this scenario")
	}

	return &t, nil
}

// Service Interface for toolbox
func (t ToolBox) Execute(Odata *services.Pipe) {
	err := legatoDb.db.Preload("Service").Find(&t).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		t.Next(Odata)
		return
	}

	SendLogMessage("*******Starting Toolbox Service*******", *t.Service.ScenarioID, nil)
	
	logData := fmt.Sprintf("Executing type (%s) : %s\n", toolBoxType, t.Service.Name)
	SendLogMessage(logData, *t.Service.ScenarioID, &t.Service.ID)

	
	switch t.Service.SubType {
	case toolBoxSleep:
		var data toolBoxSleepData
		err = json.Unmarshal([]byte(t.Service.Data), &data)
		if err != nil {
			log.Println(err)
		}

		// Goes to sleep for data.Time seconds
		logData = fmt.Sprintf("Sleeping for %d seconds \n", data.Time)
		SendLogMessage(logData, *t.Service.ScenarioID, &t.Service.ID)

		time.Sleep(time.Duration(data.Time) * time.Second)

		break
	case toolBoxRepeater:
		var data toolBoxRepeaterData
		err = json.Unmarshal([]byte(t.Service.Data), &data)
		if err != nil {
			log.Println(err)
		}

		// Repeats the tail for data.Count
		// There is a t.Next() at end so we specify data.Count - 1
		logData = fmt.Sprintf("Repeating  %d times \n", data.Count)
		SendLogMessage(logData, *t.Service.ScenarioID, &t.Service.ID)

		for i := 1; i < data.Count; i++ {
			t.Next(Odata)
		}

		break
	default:
		break
	}

	t.Next(Odata)
}

func (t ToolBox) Post(Odata *services.Pipe) {
	log.Printf("Executing type (%s) node in background : %s\n", toolBoxType, t.Service.Name)
}

func (t ToolBox) Resume(data ...interface{}) {

}

func (t ToolBox) Next(Odata *services.Pipe) {
	err := legatoDb.db.Preload("Service.Children").Find(&t).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		return
	}

	log.Printf("Executing \"%s\" Children \n", t.Service.Name)

	for _, node := range t.Service.Children {
		go func(n Service) {
			serv, err := n.Load()
			if err != nil {
				log.Println("error in loading services in Next()")
				return
			}

			serv.Execute(Odata)
		}(node)
	}

	log.Printf("*******End of \"%s\"*******", t.Service.Name)
}
