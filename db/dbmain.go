package legatoDb

import (
	"fmt"
	_ "github.com/lib/pq"
	"legato_server/env"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type LegatoDB struct {
	db *gorm.DB
}

func Connect() (*LegatoDB, error) {
	config := fmt.Sprintf("host=%s port=%s user=%s dbname=%s password=%s sslmode=disable",
		env.ENV.DatabaseHost,
		env.ENV.DatabasePort,
		env.ENV.DatabaseUsername,
		env.ENV.DatabaseName,
		env.ENV.DatabasePassword,
	)

	db, err := gorm.Open("postgres", config)
	if err != nil {
		log.Println("Error in connecting to the postgres database")
		log.Fatal(err)
	}

	// Create LegatoDB instance
	//defer db.Close() // TODO: what should happen to this?
	dbObj := LegatoDB{}
	dbObj.db = db

	// Call createSchema to create all of our tables
	err = createSchema(dbObj.db)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the database and created the tables")

	return &dbObj, nil
}

// createSchema creates database schema (tables and ...)
// for all of our models.
func createSchema(db *gorm.DB) error {
	db.AutoMigrate(User{})
	db.AutoMigrate(Scenario{})
	db.AutoMigrate(Service{})
	db.AutoMigrate(Webhook{})
	
	return nil
}
