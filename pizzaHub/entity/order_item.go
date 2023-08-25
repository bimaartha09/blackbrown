package entity

type OrderItem struct {
	ID      uint64 `json:"id"`
	OrderID uint64 `json:"order_id"`
	MenuID  uint   `json:"menu_id"`
	Total   uint64 `json:"total"`
}

var OrderItems = []OrderItem{}

func AddOrderItem(orderItem OrderItem) OrderItem {
	orderItem.ID = generateIDOrderItem()

	OrderItems = append(OrderItems, orderItem)

	return orderItem
}

func generateIDOrderItem() uint64 {
	if len(OrderItems) == 0 {
		return 1
	}

	return uint64(OrderItems[len(OrderItems)-1].ID) + 1
}
