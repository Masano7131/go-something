package main

import (
	"github.com/Masano7131/hello/src/controller"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", controller.IndexGET)
	router.Run(":8080")
}
