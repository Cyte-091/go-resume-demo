package main

import (
	"github.com/Cyte-091/go-resume-demo/docs"
	"github.com/Cyte-091/go-resume-demo/initialize"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"log"
)

func main() {
	_ = docs.SwaggerInfo
	gin.SetMode(gin.DebugMode) // 确保 debug 模式
	initialize.InitGorm()
	r := gin.New()
	r.Use(gin.Recovery(), gin.Logger()) // 就这俩中间件
	InitRouter(r)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	log.Fatal(r.Run(":8080"))
}
