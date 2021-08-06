package legatoDb

import (
	"encoding/json"
	"errors"
	"fmt"
	"legato_server/cache"
	"legato_server/env"
	"legato_server/services"
	"log"
	uuid "github.com/satori/go.uuid"
	"gorm.io/gorm"
)

const webhookType string = "webhooks"

type Webhook struct {
	gorm.Model
	Token    uuid.UUID
	IsEnable bool    `gorm:"default:False"`
	Service  Service `gorm:"polymorphic:Owner;"`
	GetMethod bool	 
	GetHeaders bool  
}

func (w *Webhook) String() string {
	return fmt.Sprintf("(@Webhooks: %+v)", *w)
}

func (w *Webhook) BeforeCreate(tx *gorm.DB) (err error) {
	w.Token = uuid.NewV4()
	w.Service.Name = "webhook"
	return nil
}

func (w *Webhook) GetURL() string {
	return fmt.Sprintf("%s/api/services/webhook/%v", env.ENV.WebUrl, w.Token)
}

func (ldb *LegatoDB) CreateWebhookForScenario(s *Scenario, wh Webhook) (*Webhook, error) {
	wh.Service.UserID = s.UserID
	wh.Service.ScenarioID = &s.ID

	ldb.db.Create(&wh)
	ldb.db.Save(&wh)
	ldb.db.Preload("Service").Find(&wh)
	return &wh, nil
}

func (ldb *LegatoDB) CreateSeparateWebhook(u *User, wh Webhook) (*Webhook, error) {
	wh.Service.UserID = u.ID
	wh.Service.ScenarioID = nil

	ldb.db.Create(&wh)
	ldb.db.Save(&wh)

	return &wh, nil
}

func (ldb *LegatoDB) CreateWebhookInScenario(u *User, s *Scenario, parent *Service, name string, x int, y int) *Webhook {
	var wh Webhook
	if parent != nil {
		wh = Webhook{Service: Service{Name: name, UserID: u.ID, ScenarioID: &s.ID, ParentID: &parent.ID, PosX: x, PosY: y}}
	} else {
		wh = Webhook{Service: Service{Name: name, UserID: u.ID, ScenarioID: &s.ID, PosX: x, PosY: y}}
	}
	ldb.db.Create(&wh)
	ldb.db.Save(&wh)
	return &wh
}

func (ldb *LegatoDB) UpdateWebhook(s *Scenario, servId uint, nwh Webhook) error {
	var serv Service
	err := ldb.db.Where(&Service{ScenarioID: &s.ID}).Where("id = ?", servId).Find(&serv).Error
	if err != nil {
		return err
	}

	var wh Webhook
	err = ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&wh).Error
	if err != nil {
		return err
	}
	if wh.Service.ID != servId {
		return errors.New("the webhook service is not in this scenario")
	}

	ldb.db.Model(&serv).Updates(nwh.Service)
	ldb.db.Model(&wh).Updates(nwh)


	return nil
}

func (ldb *LegatoDB) UpdateSeparateWebhook(u *User, wid uint, nwh Webhook) error {
	var wh Webhook
	err := ldb.db.Where("id = ?", wid).Preload("Service").Find(&wh).Error
	if err != nil {
		return err
	}
	if wh.ID != wid {
		return errors.New("the webhook service is not existed")
	}
	if wh.Service.UserID != u.ID {
		return errors.New("the webhook service is not for this user")
	}
	serv := wh.Service
	ldb.db.Model(&serv).Updates(nwh.Service)
	ldb.db.Model(&wh).Updates(nwh)

	return nil
}

func (ldb *LegatoDB) GetWebhookByService(serv Service) (*Webhook, error) {
	var wh Webhook
	err := ldb.db.Where("id = ?", serv.OwnerID).Preload("Service").Find(&wh).Error
	if err != nil {
		return nil, err
	}
	if wh.ID != uint(serv.OwnerID) {
		return nil, errors.New("the webhook service is not in this scenario")
	}

	return &wh, nil
}

func (ldb *LegatoDB) GetScenarioRootServices(s Scenario) ([]Service, error) {
	var ss []Service
	err := ldb.db.Where("parent_id is NULL").
		Where("scenario_id = ?", s.ID).
		Find(&ss).Error
	if err != nil {
		return nil, err
	}

	return ss, nil
}

func (ldb *LegatoDB) GetWebhookByUUID(uuid uuid.UUID) (*Webhook, error) {
	webhook := Webhook{}
	ldb.db.Where(&Webhook{Token: uuid}).Preload("Service").First(&webhook)
	if webhook.Token != uuid {
		return &Webhook{}, errors.New("webhook obj does not exist")
	}
	return &webhook, nil
}

func (ldb *LegatoDB) GetUserWebhooks(u *User) ([]Webhook, error) {
	var s []int
	var webhooks []Webhook
	err := ldb.db.Model(&Service{}).Where("user_id", u.ID).Where("owner_type", webhookType).Pluck("owner_id", &s).Error
	if len(s) == 0{
		return webhooks, nil
	}
	err = ldb.db.Preload("Service").Find(&webhooks, s).Error
	
	if err != nil || len(webhooks) == 0{
		return nil, err
	}
	return webhooks, nil
}

func (ldb *LegatoDB) GetUserWebhookById(u *User, wid uint) (Webhook, error) {
	var service Service
	err := ldb.db.Where(&Service{UserID: u.ID, OwnerID: int(wid)}).Find(&service).Error
	if err != nil {
		return Webhook{}, err
	}

	var webhooks Webhook
	err = ldb.db.Where("id = ?", wid).Preload("Service").Find(&webhooks).Error
	if err != nil {
		return Webhook{}, err
	}

	return webhooks, nil
}

func (ldb *LegatoDB) DeleteSeparateWebhookById(u *User, wid uint) error {
	var wh Webhook
	err := ldb.db.Where("id = ?", wid).Preload("Service").Find(&wh).Error
	if err != nil {
		return err
	}
	if wh.ID != wid {
		return errors.New("the webhook service is not existed")
	}
	// if webhook was not deleted in scenario
	if wh.Service.ID != 0{
		if wh.Service.UserID != u.ID {
			return errors.New("the webhook service is not for this user")
		}
	}


	ldb.db.Delete(&wh)
	ldb.db.Delete(&wh.Service)

	return nil
}

// Service Interface for Webhook
func (w Webhook) Execute(Odata *services.OutputData) {
	err := legatoDb.db.Preload("Service").Find(&w).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		w.Next(Odata)
		return
	}

	SendLogMessage("*******Starting Webhook Service*******", *w.Service.ScenarioID, nil)

	logData := fmt.Sprintf("Executing type (%s) : %s\n", webhookType, w.Service.Name)
	SendLogMessage(logData, *w.Service.ScenarioID, nil)
	
	w.IsEnable = true
	legatoDb.db.Save(&w)

	w.Post(Odata)

}

func (w Webhook) Post(Odata *services.OutputData) {
	err := legatoDb.db.Preload("Service").Find(&w).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		w.Next(Odata)
		return
	}
	logData := fmt.Sprintf("Executing type (%s) node in background : %s\n", webhookType, w.Service.Name)
	SendLogMessage(logData, *w.Service.ScenarioID, &w.Service.ID)

	// Saving scenario data into redis to load whenever webhook got triggered
	key := fmt.Sprintf("%s", w.Token)
	err = cache.Cache.Set(key, Odata)
	if err != nil{
		log.Println("can not set scenario data in redis", err)
	}
	
}

func (w Webhook) Resume(data ...interface{}) {
	err := legatoDb.db.Preload("Service").Find(&w).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		w.Next(&services.OutputData{})
		return
	}
	// Load data from redis and make Outputdata struct to continue scenario
	var Odata services.OutputData
	key := fmt.Sprintf("%s", w.Token)
	_, err = cache.Cache.Get(key, &Odata)
	if err!=nil{
		log.Println("can not get scenario data from redis", err)
		w.Next(&services.OutputData{})
	}
	webhookData := data[0]
	Odata.AddData(w.Service.Name, webhookData)

	w.Next(&Odata)
}

func (w Webhook) Next(Odata *services.OutputData) {
	err := legatoDb.db.Preload("Service").Preload("Service.Children").Find(&w).Error
	if err != nil {
		log.Println("!! CRITICAL ERROR !!", err)
		return
	}

	// disable webhook after receiving data
	w.IsEnable = false
	legatoDb.db.Save(&w)

	logData := fmt.Sprintf("webhook with id %v got payload:", w.Token)
	SendLogMessage(logData, *w.Service.ScenarioID, &w.Service.ID)

	dataInterface, err := Odata.GetData(w.Service.Name)
	if err == nil {
		webhookData, ok := dataInterface.(map[string]interface{})
		if ok {
			payloadJson, _ := json.Marshal(webhookData)
			SendLogMessage(string(payloadJson), *w.Service.ScenarioID, &w.Service.ID)
		}
	}else {
		fmt.Print(err.Error())
	}

	logData = fmt.Sprintf("Executing \"%s\" Children \n", w.Service.Name)
	SendLogMessage(logData, *w.Service.ScenarioID, &w.Service.ID)

	for _, node := range w.Service.Children {
		go func(n Service) {
			serv, err := n.Load()
			if err != nil {
				log.Println("error in loading services in Next()")
				return
			}

			serv.Execute(Odata)
		}(node)
	}

	logData = fmt.Sprintf("*******End of \"%s\"*******", w.Service.Name)
	SendLogMessage(logData, *w.Service.ScenarioID, &w.Service.ID)

}


func (ldb *LegatoDB) GetWebhookHistoryLogsById(u *User, wid uint) (logs []ServiceLog, err error) {
	var wdb Webhook
	err = ldb.db.Where("id = ?", wid).Preload("Service").Find(&wdb).Error
	if err != nil || wdb.ID == 0{
		return nil, errors.New("no webhook exists with given id")
	}
	err = ldb.db.Where(&ServiceLog{ServiceID: uint(wdb.Service.ID)}).Preload("Service").Preload("Messages", "message_type = ?", "json").Find(&logs).Error
	if err != nil {
		return nil, err
	}
	return logs, nil
}