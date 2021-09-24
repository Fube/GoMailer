package mailer

import "github.com/gin-gonic/gin"

type Mailer interface {
	SendEmail(*Mail) error
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
