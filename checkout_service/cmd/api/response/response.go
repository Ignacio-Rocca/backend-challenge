package response

type Error struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

type GetBasketAmount struct {
	BasketID    string  `json:"basket_id"`
	TotalAmount float64 `json:"total_amount"`
}

type CreateBasket struct {
	BasketID    string  `json:"basket_id"`
}

type AddProduct struct {
	BasketID string         `json:"basket_id"`
	Products map[string]int `json:"products"`
}
