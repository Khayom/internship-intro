package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	fmt.Println("[console]")
	
	r := gin.Default()
	r.Run(":1212")
}