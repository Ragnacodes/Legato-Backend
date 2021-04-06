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

	// Create LegatoDB instance
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
	_ = db.AutoMigrate(Scenario{})
	_ = db.AutoMigrate(Service{})
	_ = db.AutoMigrate(Position{})
  _ = db.AutoMigrate(Webhook{})
	return nil
}


func (ldb *LegatoDB)DeleteCreatedEntities() func() {
    type entity struct {
        table   string
        keyname string
        key     interface{}
    }
    var entries []entity
    hookName := "cleanupHook"

    // Setup the onCreate Hook
    ldb.db.Callback().Create().After("gorm:create").Register(hookName, func(scope *gorm.Scope) {
        fmt.Printf("Inserted entities of %s with %s=%v\n", scope.TableName(), scope.PrimaryKey(), scope.PrimaryKeyValue())
        entries = append(entries, entity{table: scope.TableName(), keyname: scope.PrimaryKey(), key: scope.PrimaryKeyValue()})
    })
    return func() {
        // Remove the hook once we're done
        defer ldb.db.Callback().Create().Remove(hookName)
        // Find out if the current db object is already a transaction
        _, inTransaction := ldb.db.CommonDB().(*sql.Tx)
        tx := ldb.db
        if !inTransaction {
            tx = ldb.db.Begin()
        }
        // Loop from the end. It is important that we delete the entries in the
        // reverse order of their insertion
        for i := len(entries) - 1; i >= 0; i-- {
            entry := entries[i]
            fmt.Printf("Deleting entities from '%s' table with key %v\n", entry.table, entry.key)
            tx.Table(entry.table).Where(entry.keyname+" = ?", entry.key).Delete("")
        }

        if !inTransaction {
            tx.Commit()
        }
    }
}