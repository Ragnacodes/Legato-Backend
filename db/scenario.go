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

func (ldb *LegatoDB) AddScenario(scenario *Scenario) error {
	ldb.db.Create(scenario)

	return nil
}