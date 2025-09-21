package router

import (
	"github.com/Cyte-091/go-resume-demo/api"
	"github.com/gin-gonic/gin"
)

func InitLoginRouter(r *gin.RouterGroup) {
	loginApi := api.LoginApi{}
	r.POST("/login", loginApi.Login)
}
