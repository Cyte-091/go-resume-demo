package main

import (
	"github.com/Cyte-091/go-resume-demo/docs"
	"github.com/Cyte-091/go-resume-demo/initialize"
	"github.com/Cyte-091/go-resume-demo/middleware"
	"github.com/Cyte-091/go-resume-demo/router"
	"github.com/gin-contrib/cors"
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
	// 1. 跨域拦截（必须放最前）
	// 1. 显式 CORS 配置（必须放最前）
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:8081"}, // 前端地址
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"}, // 允许带 Authorization
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))
	r.Use(middleware.ErrorHandler()) // ← 全局异常
	r.Use(gin.Recovery(), gin.Logger())
	router.Setup(r) // ← 唯一入口
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	r.Run(":8080")
}
