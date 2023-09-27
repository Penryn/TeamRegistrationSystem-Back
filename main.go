package main

import (
	"TeamRegistrationSystem-Back/config/database"
	"TeamRegistrationSystem-Back/config/router"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	database.Init()
	r := gin.Default()
	router.Init(r)

	err:=r.Run()
	if err !=nil{
		log.Fatal("Server start error:",err)
	}
}