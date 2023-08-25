package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"main/pizzaHub/entity"
	"main/pizzaHub/method"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestAddChefHandler(t *testing.T) {
	r := SetUpRouter()
	r.POST("/chefs", method.AddChef)
	chefReq := method.AddChefRequest{
		Name:        "Charles",
		NIP:         "PH0001",
		Description: "Cheese Pizza Chef",
	}
	jsonValue, _ := json.Marshal(chefReq)
	req, _ := http.NewRequest("POST", "/chefs", bytes.NewBuffer(jsonValue))

	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	assert.Equal(t, http.StatusCreated, w.Code)
}

func TestGetMenusHandler(t *testing.T) {
	r := SetUpRouter()
	r.GET("/menus", method.GetMenus)
	req, _ := http.NewRequest("GET", "/menus", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)

	var menus []entity.Menu
	json.Unmarshal(w.Body.Bytes(), &menus)

	fmt.Println(menus)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.NotEmpty(t, menus)
}
