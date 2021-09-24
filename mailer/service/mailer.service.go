package service

import (
	"awesomeProject1/mailer"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
)


type Dialer struct {
	Dialer *gomail.Dialer
}

func (d Dialer) SendEmail(mail *mailer.Mail) error {

	m := gomail.NewMessage()

	// Set E-Mail sender
	m.SetHeader("From", d.Dialer.Username)

	// Set E-Mail receivers
	m.SetHeader("To", mail.To)

	// Set E-Mail subject
	m.SetHeader("Subject", "Gomail test subject")

	// Set E-Mail body. You can set plain text or html with text/html
	m.SetBody("text/plain", mail.Message)

	// Now send E-Mail
	return d.Dialer.DialAndSend(m)
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
