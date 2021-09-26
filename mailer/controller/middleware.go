package controller

import (
	mailerM "GoMailer/mailer"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

func UnMarshallMail(c *gin.Context) {

	var mail mailerM.Mail

	if err := c.ShouldBindJSON(&mail); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}

	c.Set("mail", &mail)
	c.Next()
}

func ValidateEmail(c *gin.Context) {

	var mail *mailerM.Mail

	validate := validator.New()

	//mail := c.Keys["mail"].(*mailerM.Mail)
	if m, e := c.Get("mail"); !e {
		s := "ERROR! e-mail not found"
		fmt.Println(s)
		c.JSON(http.StatusBadRequest, s)
		return
	} else {
		mail = m.(*mailerM.Mail)
	}

	if err := validate.Struct(mail); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
}
