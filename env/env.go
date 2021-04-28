package env

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

const (
	DefaultTlsPort = "443"
	DefaultWebHost          = "http://localhost"

	// Postgres database
	DefaultDatabaseHost     = "database"
	DefaultDatabasePort     = "5432"
	DefaultDatabaseUsername = "legato"
	DefaultDatabasePassword = "legato"
	DefaultDatabaseName     = "legatodb"

	// App Connections URL
	SpotifyAuthenticateUrl  = "https://accounts.spotify.com/authorize?client_id=74049abbf6784599a1564060e7c9dc12&redirect_uri=http://localhost:3000/redirect/spotify/&response_type=code&scope=user-read-private&state=abc123"
	GoogleAuthenticateUrl   = "https://accounts.google.com/o/oauth2/v2/auth?client_id=906955768602-u0nu3ruckq6pcjvune1tulkq3n0kfvrl.apps.googleusercontent.com&response_type=code&scope=https://www.googleapis.com/auth/gmail.readonly&redirect_uri=http://localhost:3000/redirect/gmail/&access_type=offline"
	GitAuthenticateUrl      = "https://github.com/login/oauth/authorize?access_type=online&client_id=Iv1.9f22bc1a9e8e6822&response_type=code&scope=user%3Aemail+repo&state=thisshouldberandom&redirect_uri=http://localhost:3000/redirect/github/"
	DiscordAuthenticateUrl  = "https://discord.com/api/oauth2/authorize?access_type=online&client_id=830463353079988314&redirect_uri=http://localhost:3000/redirect/discord/&response_type=code&scope=identify+email&state=h8EecvhXJqHsG5EQ3K0gei4EUrWpaFj_HqH3WNZdrzrX1BX1COQRsTUv3-yGi3WmHQbw0EHJ58Rx1UOkvwip-Q%3D%3D"
)

type env struct {
	ServingPort      string
	DatabaseHost     string
	DatabasePort     string
	DatabaseUsername string
	DatabasePassword string
	DatabaseName     string
	WebHost          string
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

	envWebHost := os.Getenv("WEB_HOST")
	if envWebHost == "" {
		envWebHost = DefaultWebHost
	}

	ENV = env{
		ServingPort: envPort,
		// Postgres database
		DatabaseHost:     envDatabaseHost,
		DatabasePort:     envDatabasePort,
		DatabaseUsername: envDatabaseUsername,
		DatabasePassword: envDatabasePassword,
		DatabaseName:     envDatabaseName,
		WebHost:          envWebHost,
	}

	log.Printf("Environment Variables is Loaded: %+v\n", ENV)
}
