package mailer

import (
	"crypto/tls"
	"github.com/gin-gonic/gin"
	gomail "gopkg.in/mail.v2"
)

type Mailer interface {
	SendEmail(*Mail) error
}

type Dialer struct {
	Dialer *gomail.Dialer
}

type Mail struct {
	To string `json:"to" validate:"required,email"`
	Message string `json:"message" validate:"required"`
	Subject string `json:"subject"`
}

type MailerController interface {
	Inject(mailer Mailer)
	Routes(route *gin.Engine) error
}

func CreateMailer(host, email, password string, port... int) Dialer {

	truePort := 587

	// Ghetto optional param
	if len(port) != 0 {
		truePort = port[0]
	}

	embeddedDialer := gomail.NewDialer(host, truePort, email, password)

	// Set to false for prod
	embeddedDialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return Dialer{Dialer: embeddedDialer}
}

func (d Dialer) SendEmail(mail *Mail) error {

	m := gomail.NewMessage()

	m.SetHeader("From", d.Dialer.Username)
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Message)

	return d.Dialer.DialAndSend(m)
}

