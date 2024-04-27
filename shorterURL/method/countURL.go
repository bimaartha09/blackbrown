package method

import (
	"errors"
	"main/shorterURL/entity"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CountURL(ctx *gin.Context) {
	id := ctx.Param("id")

	intID, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	url, err := updateURL(intID)
	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, url)
}

func updateURL(id uint64) (*entity.Url, error) {
	for idx, url := range entity.Urls {
		if url.ID == id {
			url.Counter += 1
			entity.Urls[idx] = url
			return &url, nil
		}
	}

	return nil, errors.New("ID not found")
}
