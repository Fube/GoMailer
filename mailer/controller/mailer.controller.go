package controller

import (
	"awesomeProject1/mailer"
	mailerS "awesomeProject1/mailer/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var dialer mailerS.Dialer

func setup() {
	dialer = mailerS.CreateMailer("smtp.gmail.com", os.Getenv("EMAIL"), os.Getenv("PASSWORD"))
}

func handleSendEmail(context *gin.Context) {

	var mail mailer.Mail

	if err := context.ShouldBindJSON(&mail); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	if err := dialer.SendEmail(&mail); err != nil {
		fmt.Println(err)
		context.JSON(http.StatusInternalServerError, err)
		return
	}

	context.Status(200)
}

// Routes injection for mailer
func Routes(route *gin.Engine) {

	setup()

	group := route.Group("/mail")
	group.POST("", handleSendEmail)
}