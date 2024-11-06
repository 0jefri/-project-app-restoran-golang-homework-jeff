package model

type Order struct {
	ID       int     `json:"id"`
	UserID   int     `json:"user_id"`
	MenuID   int     `json:"menu_id"`
	Status   string  `json:"status"`
	Discount float64 `json:"discount"`
	Rating   *int    `json:"rating"`
}
