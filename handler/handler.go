package handler

import (
	"intro/golang/controller"

	"github.com/gin-gonic/gin"
)

func Handler()  {
	r := gin.Default()

	r.POST("/message/send", controller.SendMessage)
	r.GET("/message/get", controller.GetMessage)

	r.Run(":1212")
}