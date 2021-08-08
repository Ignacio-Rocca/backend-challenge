package request

type AddProduct struct {
	BasketID    string `json:"basket_id"`
	ProductCode string `json:"product_code"`
	Quantity    int    `json:"quantity"`
}
