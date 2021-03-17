package env

import (
	"github.com/joho/godotenv"
	"log"
	"os"
)

const DefaultTlsPort = "443"
const DefaultDatabaseHost = "database"
const DefaultDatabasePort = "5432"

type env struct {
	ServingPort  string
	DatabaseHost string
	DatabasePort string
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

	ENV = env{
		ServingPort:  envPort,
		DatabaseHost: envDatabaseHost,
		DatabasePort: envDatabasePort,
	}

	log.Printf("Environment Variables is Loaded: %+v\n", ENV)
}
