package method

import (
	"main/pizzaHub/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type AddChefRequest struct {
	NIP         string `json:"nip" binding:"required"`
	Name        string `json:"name" binding:"required"`
	Description string `json:"description"`
}

func AddChef(ctx *gin.Context) {
	var request AddChefRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	existingChef := entity.GetChefByNIP(request.NIP)
	if existingChef.ID > 0 {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": "chef already exist"})
		return
	}

	newChef := entity.Chef{
		NIP:         request.NIP,
		Name:        request.Name,
		Description: request.Description,
	}

	newChef = entity.AddChef(newChef)

	ctx.IndentedJSON(http.StatusCreated, newChef)
}
