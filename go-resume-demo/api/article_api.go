package api

import (
	"errors"
	"github.com/Cyte-091/go-resume-demo/dao"
	"github.com/Cyte-091/go-resume-demo/model"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ArticleApi struct{}

// CreateArticle
// @Summary 新增文章
// @Accept json
// @Produce json
// @Param article body model.Article true "article"
// @Success 200 {object} model.Article
// @Security Bearer
// @Router /article [post]
func (a *ArticleApi) CreateArticle(c *gin.Context) {
	var art model.Article
	if err := c.ShouldBindJSON(&art); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err := dao.CreateArticle(&art); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, art)
}

// DeleteArticle
// @Summary 删除文章
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} model.Article
// @Security Bearer
// @Router /article/{id} [delete]
func (a *ArticleApi) DeleteArticle(c *gin.Context) {
	// 1. 把路由参数转成 uint
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid id"})
		return
	}
	id := uint(id64)

	var art model.Article
	if err := c.ShouldBindJSON(&art); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	if err := dao.DeleteArticle(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, art)
}

// UpdateArticle
// @Summary 修改文章
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Param article body model.Article true "article"
// @Success 200 {object} model.Article
// @Security Bearer
// @Router /article/{id} [put]
func (a *ArticleApi) UpdateArticle(c *gin.Context) {
	var art model.Article
	if err := c.ShouldBindJSON(&art); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	id, _ := strconv.Atoi(c.Param("id"))
	if err := dao.UpdateArticle(uint(id), art.Title, art.Body); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	c.JSON(http.StatusOK, art)
}

// GetArticle
// @Summary 获取文章
// @Accept json
// @Produce json
// @Param id path int true "id"
// @Success 200 {object} model.Article
// @Security Bearer
// @Router /article/{id} [get]
func (a *ArticleApi) GetArticle(c *gin.Context) {
	// 1. 解析路由参数 id
	idStr := c.Param("id")
	id64, err := strconv.ParseUint(idStr, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": "invalid id"})
		return
	}
	id := uint(id64)

	// 2. 查数据库
	art, err := dao.GetArticle(id)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"msg": "article not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}

	// 3. 返回结果
	c.JSON(http.StatusOK, art)
}

// ListArticle
// @Summary 获取文章列表
// @Accept json
// @Produce json
// @Success 200 {object} model.Article
// @Security Bearer
// @Router /article [get]
func (a *ArticleApi) ListArticle(c *gin.Context) {
	// 1. 直接查数据库
	list, err := dao.ListArticle()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"msg": err.Error()})
		return
	}
	// 2. 返回列表
	c.JSON(http.StatusOK, list)
}
