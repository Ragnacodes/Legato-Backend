package legato_email

import (
	"log"
)

type Config struct {
	Identity string
	Username string `validate:"required"`
	Password string `validate:"required"`
	Host     string `validate:"required"`
}

// Email is a wrapper for the SMTP provider
type Email struct {
	config Config
}

// Single instance of Email
var emailInstance *Email

// GetEmailInstance returns the single instance of SMS.
func GetEmailInstance() *Email {
	if emailInstance == nil {
		log.Println("instance is null. you should initialize first")
	}

	return emailInstance
}
