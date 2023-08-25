package main

import (
	"main/pizzaHub/method"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/menus", method.GetMenus)
	router.POST("/chefs", method.AddChef)
	router.POST("/orders", method.AddOrder)

	router.Run("localhost:9090")
}
