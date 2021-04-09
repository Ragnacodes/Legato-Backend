package legatoDb

import (
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

type Http struct {
	gorm.Model
	Url     string
	Method  string
	Service Service `gorm:"polymorphic:Owner;"`
}

func (h *Http) String() string {
	return fmt.Sprintf("(@Http: %+v)", *h)
}

// Database methods
func (ldb *LegatoDB) CreateHttp(u *User, name string, url string, method string) *Http {
	h := Http{
		Service: Service{Name: name, UserID: int(u.ID)},
		Url:     url,
		Method:  method,
	}

	ldb.db.Create(&h)
	u.Services = append(u.Services, h.Service)
	ldb.db.Save(&h)

	return &h
}

func (ldb *LegatoDB) UpdateHttp(id string, values map[string]interface{}) (err error) {
	//for key, value := range values {
	//	if key == "name" {
	//		var wh Webhook
	//		err = ldb.db.Model(&Http{}).Where(&Http{Token: uuid}).First(&wh).Error
	//		wh.Service.Name = value.(string)
	//		ldb.db.Save(&wh)
	//	}
	//	err = ldb.db.Model(&Http{}).Where(&Http{Token: uuid}).Update(key, value).Error
	//}
	//
	//if err != nil {
	//	return err
	//}

	return nil
}

func (ldb *LegatoDB) GetHttpByID(id string) (*Http, error) {
	return nil, nil
}

// Service Interface for Http
func (h Http) Execute(...interface{}) {
	err := legatoDb.db.Preload("Service").Find(&h).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing %s node: %s\n", "webhook", h.Service.Name)

	_, err = makeHttpRequest(h.Url, h.Method)
	if err != nil {
		log.Fatalln(err)
	}

	h.Next()
}

func (h Http) Post() {
	err := legatoDb.db.Preload("Service").Find(&h).Error
	if err != nil {
		panic(err)
	}

	log.Printf("Executing %s node in background: %s\n", "webhook", h.Service.Name)
}

func (h Http) Next(...interface{}) {
	err := legatoDb.db.Preload("Service.Children").Find(&h).Error
	if err != nil {
		panic(err)
	}

	for _, node := range h.Service.Children {
		node.LoadOwner().Execute()
	}
}

// Service interface helper functions
func makeHttpRequest(url string, method string) (res *http.Response, err error) {
	switch method {
	case http.MethodGet:
		res, err = http.Get(url)
		break
	case http.MethodPost:
		res, err = http.Post(url, "application/json", nil)
		break
	}

	if err != nil {
		return nil, err
	}

	return res, nil
}
