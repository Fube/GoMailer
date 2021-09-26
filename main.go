package main

import (
	mailerM "GoMailer/mailer"
	mailerC "GoMailer/mailer/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"os"
)

func LoadEnv() {

	if err := godotenv.Load(); err != nil {
		fmt.Println(err)
		panic("Unable to load .env file")
	}
}


func main() {

	LoadEnv()

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