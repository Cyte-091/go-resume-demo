package main

import (
	"day1/service"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	us := &service.UserService{}
	r.GET("/user/:id", us.GetUser)
}
