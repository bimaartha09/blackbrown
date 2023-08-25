package method

import (
	"errors"
	"main/pizzaHub/entity"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

type AddOrderRequest struct {
	Items []OrderItemRequest `json:"items"`
}

type OrderItemRequest struct {
	MenuID uint   `json:"menu_id" binding:"required"`
	Total  uint64 `json:"total" binding:"required"`
}

type AddOrderResponse struct {
	ID         uint64             `json:"id"`
	ChefID     uint64             `json:"chef_id"`
	State      string             `json:"state"`
	TotalPrice uint               `json:"total_price"`
	StartTime  time.Time          `json:"start_time"`
	EndTime    time.Time          `json:"end_time"`
	Items      []entity.OrderItem `json:"items"`
	CreatedAt  time.Time          `json:"created_at"`
	UpdatedAt  time.Time          `json:"updated_at"`
}

func AddOrder(ctx *gin.Context) {
	var request AddOrderRequest

	if err := ctx.ShouldBindJSON(&request); err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	totalTime, totalPrice, err := getOrderDetail(request)
	if err != nil {
		ctx.IndentedJSON(http.StatusBadRequest, gin.H{"errors": err.Error()})
		return
	}

	if len(entity.Chefs) == 0 {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"errors": "chef are not exist. add first."})
		return
	}

	newOrder := entity.Order{
		TotalPrice: uint(totalPrice),
		StartTime:  time.Now(),
		EndTime:    time.Now().Add(time.Second * time.Duration(totalTime)),
		State:      entity.OrderInProgress,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	var order entity.Order

	// know who chef was assigned.
	var closestOrder entity.Order
	for _, c := range entity.Chefs {
		lastOrder := entity.GetLastOrderbyChefID(c.ID)

		if (lastOrder == entity.Order{}) {
			newOrder.ChefID = c.ID
			order = entity.AddOrder(newOrder)
			closestOrder = entity.Order{}
			break
		}

		if time.Now().After(lastOrder.EndTime) {
			entity.UpdateStateOrder(lastOrder, entity.OrderDone)
			newOrder.ChefID = c.ID
			order = entity.AddOrder(newOrder)
			closestOrder = entity.Order{}
			break
		}

		if (closestOrder.ID == 0) || (lastOrder.EndTime).Before(closestOrder.EndTime) {
			closestOrder = lastOrder
		}
	}

	if closestOrder.ID > 0 {
		newOrder.ChefID = closestOrder.ChefID
		newOrder.StartTime = closestOrder.EndTime
		newOrder.EndTime = closestOrder.EndTime.Add(time.Second * time.Duration(totalTime))
		order = entity.AddOrder(newOrder)
	}

	if order.ID == 0 {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"errors": err.Error()})
		return
	}

	// add order item
	var responseOrderItems []entity.OrderItem

	for _, it := range request.Items {
		newOrderItem := entity.OrderItem{
			OrderID: order.ID,
			MenuID:  it.MenuID,
			Total:   it.Total,
		}

		newOrderItem = entity.AddOrderItem(newOrderItem)
		responseOrderItems = append(responseOrderItems, newOrderItem)
	}

	responseAddOrder := AddOrderResponse{
		ID:         order.ID,
		ChefID:     order.ChefID,
		TotalPrice: order.TotalPrice,
		StartTime:  order.StartTime,
		EndTime:    order.EndTime,
		State:      entity.OrderStatetoString(order.State),
		Items:      responseOrderItems,
		CreatedAt:  order.CreatedAt,
		UpdatedAt:  order.UpdatedAt,
	}

	ctx.IndentedJSON(http.StatusCreated, responseAddOrder)
}

func getOrderDetail(req AddOrderRequest) (int64, int64, error) {
	var totalTime, totalPrice int64

	for _, it := range req.Items {
		menu, err := entity.GetMenuByID(int64(it.MenuID))

		if err != nil {
			return 0, 0, errors.New(err.Error())
		}

		totalTime += int64(menu.Duration) * int64(it.Total)
		totalPrice += int64(menu.Price) * int64(it.Total)
	}

	return totalTime, totalPrice, nil
}
