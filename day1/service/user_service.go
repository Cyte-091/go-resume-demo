package service

import (
	"day1/model"
	"github.com/gin-gonic/gin"
	"strconv"
)

type UserService struct{}

func (s *UserService) GetUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	c.JSON(200, model.User{ID: uint(id), Name: "fake"})
}
