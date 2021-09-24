package controller

import (
	mailerM "GoMailer/mailer"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
	"os"
)

var mailer mailerM.Mailer

func Inject(m mailerM.Mailer) {
	mailer = m
}

func handleSendEmail(context *gin.Context) {

	validate := validator.New()



	if err := validate.Struct(mail); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusBadRequest, err)
		return
	}

	if err := mailer.SendEmail(&mail); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.Status(200)
}

// Routes injection for mailer
func Routes(route *gin.Engine) {

	if mailer == nil {
		_, _ = fmt.Fprintln(os.Stderr, fmt.Errorf("dialer dependency not injected or point to nil"))
		os.Exit(1)
	}

	group := route.Group("/mail")
	group.POST("", handleSendEmail)
}