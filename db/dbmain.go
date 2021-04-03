package legatoDb

import (
	"fmt"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"legato_server/env"
	"log"
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

	db, err := gorm.Open(postgres.Open(config), &gorm.Config{})
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
	_ = db.AutoMigrate(User{})
	_ = db.AutoMigrate(Scenario{})
	_ = db.AutoMigrate(Service{})

	return nil
}
