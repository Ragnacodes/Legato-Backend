package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const (
	DefaultTlsPort = "443"
	// Postgres database
	DefaultDatabaseHost     = "database"
	DefaultDatabasePort     = "5432"
	DefaultDatabaseUsername = "legato"
	DefaultDatabasePassword = "legato"
	DefaultDatabaseName     = "legatodb"
)

type env struct {
	ServingPort      string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
}

var ENV env

func LoadEnv() {
	_ = godotenv.Load("env/.env")

	envPort := os.Getenv("PORT")
	if envPort == "" {
		envPort = "8080"
		// Later it should be DefaultTlsPort
		// envPort = DefaultTlsPort
	}

	envDatabaseHost := os.Getenv("DATABASE_HOST")
	if envDatabaseHost == "" {
		envDatabaseHost = DefaultDatabaseHost
	}

	envDatabasePort := os.Getenv("DATABASE_PORT")
	if envDatabasePort == "" {
		envDatabasePort = DefaultDatabasePort
	}

	envDatabaseUsername := os.Getenv("DATABASE_USERNAME")
	if envDatabaseUsername == "" {
		envDatabaseUsername = DefaultDatabaseUsername
	}

	envDatabasePassword := os.Getenv("DATABASE_PASSWORD")
	if envDatabasePassword == "" {
		envDatabasePassword = DefaultDatabasePassword
	}

	envDatabaseName := os.Getenv("DATABASE_NAME")
	if envDatabaseName == "" {
		envDatabaseName = DefaultDatabaseName
	}

	ENV = env{
		ServingPort: envPort,
		// Postgres database
		DatabaseHost:     envDatabaseHost,
		DatabasePort:     envDatabasePort,
		DatabaseUsername: envDatabaseUsername,
		DatabasePassword: envDatabasePassword,
		DatabaseName:     envDatabaseName,
	}

	log.Printf("Environment Variables is Loaded: %+v\n", ENV)
}
