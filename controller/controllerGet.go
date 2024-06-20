package controller

import (
	"intro/golang/structures"

	"github.com/gin-gonic/gin"
)

func GetMessage(c *gin.Context)  {
	c.JSON(200, structures.MessageSlice)
}