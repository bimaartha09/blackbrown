package main

import (
	"main/shorterURL/method"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/urlshorters", method.Create)
	router.GET("/urlsbydestination", method.RetrieveURLbyDestination)
	router.GET("/urlshorters", method.RetrieveURL)
	router.GET("/urlshorters/counter/:id", method.CountURL)

	router.Run("localhost:9090")
}
