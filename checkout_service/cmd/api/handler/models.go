package handler

// RESPONSES

type ErrorResponse struct {
	Code    int    `json:"status_code"`
	Message string `json:"message"`
}

type BasketAmountResponse struct {
	BasketID    string  `json:"basket_id"`
	TotalAmount float64 `json:"total_amount"`
}

// REQUESTS

type AddProductRequest struct {
	BasketID    string `json:"basket_id"`
	ProductCode string `json:"product_code"`
}
