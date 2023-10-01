package main

import (
	"TeamRegistrationSystem-Back/app/midwares"
	"TeamRegistrationSystem-Back/config/database"
	"TeamRegistrationSystem-Back/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.Use(midwares.JWTAuthMiddleware())
	router.Init(r)
	err:=r.Run()
	if err !=nil{
		log.Fatal("Server start error:",err)
	}
}