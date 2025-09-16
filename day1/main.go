package main

import (
	"day1/service"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	us := &service.UserService{}
	r.GET("/user/:id", us.GetUser)
	r.Run(":8080")
}
