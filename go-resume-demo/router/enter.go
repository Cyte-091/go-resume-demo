package router

import (
	"github.com/Cyte-091/go-resume-demo/middleware"
	"github.com/gin-gonic/gin"
)

// Setup 汇总所有路由
func Setup(r *gin.Engine) {
	// 1. 登录（公开）
	InitLoginRouter(r.Group(""))

	// 2. 文章（需 JWT）
	article := r.Group("")
	article.Use(middleware.JWT()) // 先挂守卫
	InitArticleRouter(article)

	// 3. 示例 Hello 接口（可选）
	InitUserRouter(r.Group(""))
}
