package api

import (
	"github.com/Cyte-091/go-resume-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

type LoginApi struct{}
type LoginReq struct {
	UserName string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// Login
// @Summary 登录
// @Accept json
// @Produce json
// @Param req body LoginReq true "req"
// @Success 200 {object} map[string]interface{}
// @Router /login [post]
func (l *LoginApi) Login(c *gin.Context) {
	var req LoginReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"msg": err.Error()})
		return
	}
	// 简易硬编码校验
	if req.UserName != "admin" || req.Password != "123456" {
		c.JSON(http.StatusUnauthorized, gin.H{"msg": "user or password errno"})
		return
	}
	token, _ := utils.GenerateToken(req.UserName)
	c.JSON(http.StatusOK, gin.H{"token": token})
}
