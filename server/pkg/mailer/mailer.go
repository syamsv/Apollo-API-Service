package mailer

import (
	"github.com/go-mail/mail"
	"github.com/syamsv/apollo/config"
)

func SendActivactionMail(email, content string) error {
	m := mail.NewMessage()
	m.SetHeader("From", config.SMTP_USER)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Email Activation")
	m.SetBody("text/html", content)
	d := mail.NewDialer(config.SMTP_SERVER, config.SMTP_PORT, config.SMTP_USER, config.SMTP_PASSWORD)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
