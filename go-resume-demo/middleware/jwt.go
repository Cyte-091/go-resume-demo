package middleware

import (
	"github.com/Cyte-091/go-resume-demo/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.GetHeader("Authorization")
		if token == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "need token"})
			c.Abort()
			return
		}
		if len(token) > 7 && token[:7] == "Bearer " {
			token = token[7:]
		}
		claims, err := utils.ParseToken(token)
		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "invalid token"})
			c.Abort()
			return
		}
		c.Set("username", claims.UserName) // 后续可取值
		c.Next()
	}
}
