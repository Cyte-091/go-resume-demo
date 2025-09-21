package router

import (
	"github.com/Cyte-091/go-resume-demo/service"
	"github.com/gin-gonic/gin"
)

func InitUserRouter(r *gin.RouterGroup) {
	us := &service.UserService{}
	r.GET("/user/:id", us.GetUser)
}
