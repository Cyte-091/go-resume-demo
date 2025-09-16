package router

import (
	"github.com/Cyte-091/go-resume-demo/api"
	"github.com/gin-gonic/gin"
)

func InitArticleRouter(r *gin.RouterGroup) {
	artApi := api.ArticleApi{}
	r.POST("/article", artApi.CreateArticle)
	r.DELETE("/article/:id", artApi.DeleteArticle)
	r.PUT("/article/:id", artApi.UpdateArticle)
	r.GET("/article/:id", artApi.GetArticle)
	r.GET("/articles", artApi.ListArticle)
}
