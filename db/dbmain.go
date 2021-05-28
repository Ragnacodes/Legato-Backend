package legatoDb

import (
	"fmt"
	"legato_server/env"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type LegatoDB struct {
	db *gorm.DB
}

var legatoDb LegatoDB

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

	// AddWebhookToScenario LegatoDB instance
	//defer db.Close() // TODO: what should happen to this?
	legatoDb.db = db

	// Call createSchema to create all of our tables
	err = createSchema(legatoDb.db)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the database and created the tables")

	return &legatoDb, nil
}

// createSchema creates database schema (tables and ...)
// for all of our models.
func createSchema(db *gorm.DB) error {
	_ = db.AutoMigrate(User{})
	_ = db.AutoMigrate(Connection{})
	_ = db.AutoMigrate(Scenario{})
	_ = db.AutoMigrate(Service{})
	_ = db.AutoMigrate(Webhook{})
	_ = db.AutoMigrate(Http{})
	_ = db.AutoMigrate(Telegram{})
	_ = db.AutoMigrate(Spotify{})
	_ = db.AutoMigrate(Token{})
	_ = db.AutoMigrate(Ssh{})
	_ = db.AutoMigrate(History{})
	_ = db.AutoMigrate(ServiceLog{})
	_ = db.AutoMigrate(LogMessage{})

	return nil
}
