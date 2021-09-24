package main

import (
	mailerC "awesomeProject1/mailer/controller"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func LoadEnv() {

	fmt.Println("I got called")
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