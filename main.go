package main

import (
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/config/database"
	"TeamRegistrationSystem-Back/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	userService.CreateAdministrator()
	r.Static("/uploads", "./uploads")
	router.Init(r)
	err:=r.Run()
	if err !=nil{
		log.Fatal("Server start error:",err)
	}
}