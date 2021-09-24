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

	validate := validator.New()

	mail := c.Keys["mail"].(*mailerM.Mail)

	if mail == nil {
		c.JSON(http.StatusInternalServerError, "Unable to unmarshall email or none set")
	}

	if err := validate.Struct(mail); err != nil {
		fmt.Println(err)
		c.JSON(http.StatusBadRequest, err)
		return
	}
}
