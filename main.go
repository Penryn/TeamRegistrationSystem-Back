package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	err:=r.Run()
	if err !=nil{
		log.Fatal("Server start error:",err)
	}
}