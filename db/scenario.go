package legatoDb

import (
	"gorm.io/gorm"
)

type Scenario struct {
	gorm.Model
	//Name   string `gorm:"unique"`
	Name   string
	UserID uint
}

func (edb *LegatoDB) AddScenario(scenario *Scenario) error {
	edb.db.Create(scenario)

	return nil
}