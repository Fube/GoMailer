package utils

import (
	mailerS "GoMailer/mailer/service"
	"crypto/tls"
	gomail "gopkg.in/mail.v2"
)

func CreateMailer(host, email, password string, port... int) mailerS.Dialer {

	truePort := 587

	// Ghetto optional param
	if len(port) != 0 {
		truePort = port[0]
	}

	embeddedDialer := gomail.NewDialer(host, truePort, email, password)

	// Set to false for prod
	embeddedDialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	return mailerS.Dialer{Dialer: embeddedDialer}
}

