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
	db, err := gorm.Open("postgres", fmt.Sprintf(
		"host=%s port=%s user=legato dbname=legatodb password=legato sslmode=disable",
		env.ENV.DatabaseHost, env.ENV.DatabasePort,
	))

	if err != nil {
		log.Println("Error in connecting to the postgres database")
		log.Fatal(err)
	}

	dbObj := LegatoDB{}

	//defer db.Close() // TODO: what should happen to this?
	dbObj.db = db

	err = createSchema(dbObj.db)
	if err != nil {
		return nil, err
	}
	log.Println("Connected to the database and created the tables")

	return &dbObj, nil
}

// createSchema creates database schema for Printer and ReciptRecord models.
func createSchema(db *gorm.DB) error {
	db.AutoMigrate(User{})

	return nil
}
