package method

import (
	"main/pizzaHub/entity"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetMenus(ctx *gin.Context) {
	menus := entity.List()

	ctx.IndentedJSON(http.StatusOK, menus)
}
