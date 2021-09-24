package service

import (
	mailerM "GoMailer/mailer"
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
