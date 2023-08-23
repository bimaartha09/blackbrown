package method

import (
	"errors"
	"main/shorterURL/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RetrieveURLbyDestinationRequest struct {
	Destination string `form:"destination" binding:"required"`
}

func RetrieveURLbyDestination(ctx *gin.Context) {
	var req RetrieveURLbyDestinationRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	url, err := searchURLByDestination(req.Destination)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, url)
}

func searchURLByDestination(destination string) (*entity.Url, error) {
	for _, todo := range entity.Urls {
		if todo.Destination == destination {
			return &todo, nil
		}
	}

	return nil, errors.New("destination not found")
}
