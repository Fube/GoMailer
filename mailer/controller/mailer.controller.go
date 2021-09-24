package controller

import (
	mailerM "GoMailer/mailer"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

type MailerControllerImpl struct {
	mailer mailerM.Mailer
}


func (m *MailerControllerImpl) Inject(mailer mailerM.Mailer) {
	m.mailer = mailer
}

func (m MailerControllerImpl) handleSendEmail(context *gin.Context) {

	mail := context.Keys["mail"].(*mailerM.Mail)

	if err := m.mailer.SendEmail(mail); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.Status(200)
}

// Routes injection for mailer
func (m MailerControllerImpl) Routes(route *gin.Engine) error {

	if m.mailer == nil {
		return errors.New("dialer dependency not injected or point to nil")
	}

	group := route.Group("/mail").Use(UnMarshallMail, ValidateEmail)
	group.POST("", m.handleSendEmail)
	return nil
}