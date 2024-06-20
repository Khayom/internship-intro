package controller

import (
	"fmt"
	"intro/golang/structures"
	"time"

	"github.com/gin-gonic/gin"
)

func SendMessage(c *gin.Context)  {
	var shablon structures.Message
	c.ShouldBindJSON(&shablon)

	if shablon.Message == ""{
		c.JSON(400, "Error! Empty line")
	}else{
		// ! Generate Date
		Year, Month, Day, Hour, Minute, Second := time.Now().Year(), time.Now().Month(), time.Now().Day(), time.Now().Hour(), time.Now().Minute(), time.Now().Second()
		Date := fmt.Sprintf("%v.%v.%v - %v:%v:%v", Day, Month, Year, Hour, Minute, Second)
		shablon.Data = Date

		shablon.Status = "Unread"

		c.JSON(200, "Succes")

		structures.MessageSlice = append(structures.MessageSlice, shablon)
	}
}