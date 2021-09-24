package main

import (
	mailerC "GoMailer/mailer/controller"
	mailerS "GoMailer/mailer/service"
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

	dialer := mailerS.CreateMailer("smtp.gmail.com", os.Getenv("EMAIL"), os.Getenv("PASSWORD"))

	mailerC.Inject(&dialer)
	mailerC.Routes(router)
}