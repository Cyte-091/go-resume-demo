package middleware

import (
	"github.com/Cyte-091/go-resume-demo/api"
	"github.com/Cyte-091/go-resume-demo/errno"
	"github.com/gin-gonic/gin"
	"net/http"
)

func ErrorHandler() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				// panic → 500
				api.Fail(c, http.StatusInternalServerError, "系统异常")
			}
		}()
		c.Next()
		// 404 统一返回
		if c.Writer.Status() == http.StatusNotFound {
			api.Fail(c, errno.PageNotFound, "页面不存在")
		}
		// 401 统一返回
		if c.Writer.Status() == http.StatusUnauthorized {
			api.Fail(c, errno.UserNotLogin, "未登录")
		}
		// 400 统一返回
		if c.Writer.Status() == http.StatusBadRequest {
			api.Fail(c, errno.ParamErr, "参数错误")
		}

	}
}
