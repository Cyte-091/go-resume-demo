package api

import "github.com/gin-gonic/gin"

type R struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func OK(c *gin.Context, data interface{}) {
	c.JSON(200, R{Code: 0, Msg: "success", Data: data})
}
func Fail(c *gin.Context, code int, msg string) {
	c.JSON(200, R{Code: code, Msg: msg, Data: nil})
}
