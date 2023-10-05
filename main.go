package main

import (
	"TeamRegistrationSystem-Back/app/midwares"
	"TeamRegistrationSystem-Back/app/services/userService"
	"TeamRegistrationSystem-Back/config/database"
	"TeamRegistrationSystem-Back/config/router"
	"log"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	r.Use(cors.Default())
	r.Use(midwares.RateLimitMiddleware(time.Second,100,100))
	r.Use(midwares.ErrHandler())
	r.NoMethod(midwares.HandleNotFound)
	r.NoRoute(midwares.HandleNotFound)
	userService.CreateAdministrator()
	r.Static("/uploads", "./uploads")
	router.Init(r)
	err:=r.Run()
	if err !=nil{
		log.Fatal("Server start error:",err)
	}
}