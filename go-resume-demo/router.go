package main

import (
	"github.com/gin-gonic/gin"
	"go-resume-demo/service"
)

func RegisterRouter(r *gin.Engine) {
	us := &service.UserService{}
	r.GET("/user/:id", us.GetUser)
}
