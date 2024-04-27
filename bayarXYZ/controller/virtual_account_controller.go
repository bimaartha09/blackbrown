package controller

import (
	"database/sql"
	"main/bayarXYZ/request"
	"main/bayarXYZ/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

type VirtualAccountController struct {
	DB *sql.DB
}

func NewVirtualAccountController(db *sql.DB) VirtualAccountController {
	return VirtualAccountController{
		DB: db,
	}
}

type CreateUrlRequest struct {
	Name        string `json:"name" binding:"required"`
	Destination string `json:"destination" binding:"required"`
	ExpireDay   int    `json:"expire_day"`
}

func (a VirtualAccountController) GenerateVA(ctx *gin.Context) {

	var req request.GenerateVARequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	srv := service.NewVirtualAccountService(a.DB)
	id, err := srv.CreateVA(req)

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
