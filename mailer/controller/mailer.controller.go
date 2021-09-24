package controller

import (
	mailerM "GoMailer/mailer"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

var mailer mailerM.Mailer

func Inject(m mailerM.Mailer) {
	mailer = m
}

func handleSendEmail(context *gin.Context) {

	mail := context.Keys["mail"].(*mailerM.Mail)

	if err := mailer.SendEmail(mail); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.Status(200)
}

// Routes injection for mailer
func Routes(route *gin.Engine) error {

	if mailer == nil {
		return errors.New("dialer dependency not injected or point to nil")
	}

	group := route.Group("/mail").Use(UnMarshallMail, ValidateEmail)
	group.POST("", handleSendEmail)
	return nil
}