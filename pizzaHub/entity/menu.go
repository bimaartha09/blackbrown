package entity

import (
	"fmt"
)

type Menu struct {
	ID       uint64 `json:"id"`
	Name     string `json:"name"`
	Duration uint   `json:"duration"`
	Price    uint   `json:"price"`
}

var Menus = []Menu{
	{ID: 1, Name: "Chesse Pizza", Duration: 3, Price: 50000},
	{ID: 2, Name: "BBQ Pizza", Duration: 5, Price: 60000},
}

func List() []Menu {
	return Menus
}

func GetMenuByID(id int64) (Menu, error) {
	if id == 0 {
		return Menu{}, fmt.Errorf("Menu ID %v is not found", id)
	}

	return Menus[id-1], nil
}
