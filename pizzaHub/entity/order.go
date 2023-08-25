package entity

import "time"

type Order struct {
	ID         uint64     `json:"id"`
	ChefID     uint64     `json:"chef_id"`
	State      OrderState `json:"state"`
	TotalPrice uint       `json:"total_price"`
	StartTime  time.Time  `json:"start_time"`
	EndTime    time.Time  `json:"end_time"`
	CreatedAt  time.Time  `json:"created_at"`
	UpdatedAt  time.Time  `json:"updated_at"`
}

var Orders = []Order{}

func GetLastOrderbyChefID(chefID uint64) Order {
	for i := len(Orders) - 1; i >= 0; i-- {
		if Orders[i].ChefID == chefID {
			return Orders[i]
		}
	}

	return Order{}
}

func AddOrder(order Order) Order {
	order.ID = generateIDOrder()

	Orders = append(Orders, order)

	return order
}

func UpdateStateOrder(order Order, state OrderState) {
	Orders[order.ID-1].State = state
}

func generateIDOrder() uint64 {
	if len(Orders) == 0 {
		return 1
	}

	return uint64(Orders[len(Orders)-1].ID) + 1
}

type OrderState uint8

const (
	OrderInProgress OrderState = iota + 1
	OrderDone
)

func OrderStatetoString(order OrderState) string {
	mapping := map[OrderState]string{
		OrderInProgress: "In Progress",
		OrderDone:       "Done",
	}

	return mapping[order]
}
