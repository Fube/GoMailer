package service

import (
	mailerM "GoMailer/mailer"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
)


type Dialer struct {
	Dialer *gomail.Dialer
}

func (d Dialer) SendEmail(mail *mailerM.Mail) error {

	m := gomail.NewMessage()

	m.SetHeader("From", d.Dialer.Username)
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Message)

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
