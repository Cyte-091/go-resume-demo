package main

import (
	"github.com/gin-gonic/gin"
	"go-resume-demo/service"
)

func main() {
	r := gin.Default()
	us := &service.UserService{}
	r.GET("/user/:id", us.GetUser)
	r.Run(":8080")
}
