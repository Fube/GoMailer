package main

import (
	mailerC "GoMailer/mailer/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
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

	mailerC.Routes(router)
}