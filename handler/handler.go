package handler

import (
	"intro/golang/controller"

	"github.com/gin-gonic/gin"
)

func Handler()  {
	r := gin.Default()

	r.POST("/post", controller.Post)
	r.GET("/get", controller.Get)

	r.Run(":1212")
}