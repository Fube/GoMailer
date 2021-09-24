package controller

import (
	mailerM "GoMailer/mailer"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func UnMarshallMail(c *gin.Context) {

	var mail mailerM.Mail

	if err := c.ShouldBindJSON(&mail); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.Set("mail", mail)
	c.Next()
}
