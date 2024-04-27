package controller

import (
	"database/sql"
	"main/bayarXYZ/request"
	"main/bayarXYZ/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type PaymentController struct {
	DB *sql.DB
}

func NewPaymentController(db *sql.DB) PaymentController {
	return PaymentController{
		DB: db,
	}
}

func (c PaymentController) CreatePayment(ctx *gin.Context) {
	var req request.CreatePaymentRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	srv := service.NewPaymentService(c.DB)
	id, err := srv.CreatePayment(req)

	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	res := map[string]interface{}{
		"code":    200,
		"message": "OK",
		"data": map[string]interface{}{
			"id": id,
		},
	}

	ctx.IndentedJSON(http.StatusCreated, res)
}

func (c PaymentController) ListPaymentVA(ctx *gin.Context) {
	pService := service.NewPaymentService(c.DB)
	list, err := pService.ListPaymentVA(1)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	res := map[string]interface{}{
		"code":    200,
		"message": "OK",
		"data":    list,
	}

	ctx.IndentedJSON(http.StatusOK, res)
}
