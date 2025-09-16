package main

import (
	"github.com/Cyte-091/go-resume-demo/router"
	"github.com/Cyte-091/go-resume-demo/service"
	"github.com/gin-gonic/gin"
)

func RegisterRouter(r *gin.Engine) {
	us := &service.UserService{}
	r.GET("/user/:id", us.GetUser)

}

func InitRouter(r *gin.Engine) {
	router.InitArticleRouter(r.Group(""))
}
