package main

import (
	"github.com/Cyte-091/go-resume-demo/docs"
	"github.com/Cyte-091/go-resume-demo/initialize"
	"github.com/Cyte-091/go-resume-demo/middleware"
	"github.com/Cyte-091/go-resume-demo/router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @securityDefinitions.apikey Bearer
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	_ = docs.SwaggerInfo
	initialize.InitGorm()
	r := gin.New()
	r.Use(middleware.ErrorHandler()) // ← 全局异常
	r.Use(gin.Recovery(), gin.Logger())
	router.Setup(r) // ← 唯一入口
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
