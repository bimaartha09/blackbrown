package method

import (
	"main/shorterURL/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type CreateUrlRequest struct {
	Name        string `json:"name" binding:"required"`
	Destination string `json:"destination" binding:"required"`
	ExpireDay   int    `json:"expire_day"`
}

func Create(ctx *gin.Context) {
	var request CreateUrlRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	for _, url := range entity.Urls {
		if url.Destination == request.Destination || url.Name == request.Name {
			ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": "URL already exist."})
			return
		}
	}

	var expireTime time.Time

	if request.ExpireDay > 0 {
		expireTime = time.Now().AddDate(0, 0, request.ExpireDay)
	}

	url := entity.Url{
		ID:          uint64(len(entity.Urls) + 1),
		Name:        request.Name,
		Destination: request.Destination,
		Counter:     0,
		ExpireTime:  expireTime,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	entity.Urls = append(entity.Urls, url)
	ctx.IndentedJSON(http.StatusCreated, url)
}
