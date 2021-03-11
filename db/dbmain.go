package legatoDb

import (
	"legato_server/env"
	"fmt"
	_ "github.com/lib/pq"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

type LegatoDB struct {
	db *gorm.DB
}

func Connect() (*LegatoDB, error) {
	args := fmt.Sprintf("host=%s port=%s user=admin dbname=legato_db password=admin sslmode=disable",
		env.ENV.DatabaseHost, env.ENV.DatabasePort)
	db, err := gorm.Open("postgres", args)
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

type ReciptRecord struct {
	UUID          int64
	PrinterSerial string
	// TODO: Add the keys
}

func (s *ReciptRecord) String() string {
	return fmt.Sprintf("ReciptRecord<%v>", s)
}

// createSchema creates database schema for Printer and ReciptRecord models.
func createSchema(db *gorm.DB) error {
	db.AutoMigrate(User{})

	return nil
}
