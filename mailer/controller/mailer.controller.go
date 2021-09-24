package controller

import (
	"GoMailer/mailer"
	mailerS "GoMailer/mailer/service"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

var dialer mailerS.Dialer

func Inject(d *mailerS.Dialer) {
	dialer = *d
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

	if dialer.Dialer == nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("dialer dependency not injected or point to nil"))
		os.Exit(1)
	}

	group := route.Group("/mail")
	group.POST("", handleSendEmail)
}