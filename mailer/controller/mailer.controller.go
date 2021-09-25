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

	var mail interface{}
	var exists bool
	if mail, exists = context.Get("mail"); !exists {
		fmt.Println("E-mail not in context")
		context.JSON(http.StatusBadRequest, "Unable to parse e-mail from body")
		return
	}

	if err := m.mailer.SendEmail(mail.(*mailerM.Mail)); err != nil {
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