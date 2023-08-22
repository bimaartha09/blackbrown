package entity

type Todo struct {
	ID        string `json:"id"`
	Item      string `json:"item"`
	Completed bool   `json:"completed"`
}

var Todos = []Todo{
	{ID: "1", Item: "Bango", Completed: true},
	{ID: "2", Item: "Teh Pucuk", Completed: false},
}
