package method

import (
	"main/shorterURL/entity"
	"math"
	"net/http"

	"github.com/gin-gonic/gin"
)

type RetrieveURLRequest struct {
	Limit  int    `form:"limit" default:"20" binding:"min=1,max=20"`
	Offset int    `form:"offset" default:"0"`
	Sort   string `form:"sort" default:"ASC"`
}

func RetrieveURL(ctx *gin.Context) {
	var req RetrieveURLRequest

	if err := ctx.ShouldBindQuery(&req); err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	url, err := searchURL(req)

	if err != nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"errors": err.Error()})
		return
	}

	ctx.IndentedJSON(http.StatusOK, url)
}

func searchURL(req RetrieveURLRequest) ([]entity.Url, error) {
	arr := entity.Urls

	if req.Sort == "DESC" {
		arr = reverseArray(entity.Urls)
	}

	start := req.Offset
	end := math.Min(float64(req.Offset+req.Limit), float64(len(arr)))

	return arr[start:int(end)], nil
}

func reverseArray(s []entity.Url) []entity.Url {
	a := make([]entity.Url, len(s))
	copy(a, s)

	for i := len(a)/2 - 1; i >= 0; i-- {
		opp := len(a) - 1 - i
		a[i], a[opp] = a[opp], a[i]
	}

	return a
}
