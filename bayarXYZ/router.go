package main

import (
	"database/sql"
	"main/bayarXYZ/controller"

	"github.com/gin-gonic/gin"
)

func callRouter(db *sql.DB) {
	router := gin.Default()

	paymentController := controller.NewPaymentController(db)
	vaController := controller.NewVirtualAccountController(db)

	router.GET("/payment/va", paymentController.ListPaymentVA)
	router.POST("/payment", paymentController.CreatePayment)
	router.POST("/va", vaController.GenerateVA)

	router.Run("localhost:9090")
}
