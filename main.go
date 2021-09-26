package main

import (
	mailerM "GoMailer/mailer"
	mailerC "GoMailer/mailer/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func main() {

	godotenv.Load()

	router := gin.Default()
	defer router.Run()

	dialer := mailerM.CreateMailer("smtp.gmail.com", os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	mailerController := mailerC.MailerControllerImpl{}

	mailerController.Inject(&dialer)
	if err := mailerController.Routes(router); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}