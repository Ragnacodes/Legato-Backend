package legato_email

import (
	"fmt"
	"github.com/jordan-wright/email"
	"net/smtp"
)

func InitGmail(c Config) (*Email, error) {
	// Create the instance
	emailInstance = &Email{
		config: c,
	}

	// TODO: Some other actions might be happened here.

	return emailInstance, nil
}

func (se *Email) SendNewGmail(ne *email.Email) error {
	// Sending info
	ne.From = fmt.Sprintf("SPCU <%s>", emailInstance.config.Username)

	// Authenticate Gmail
	cred := smtp.PlainAuth(
		se.config.Identity,
		se.config.Username,
		se.config.Password,
		se.config.Host,
	)

	// Send a new message
	address := fmt.Sprintf("%s:587", se.config.Host)

	return ne.Send(address, cred)
}
