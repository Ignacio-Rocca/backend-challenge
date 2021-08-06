package checkout

type Basket struct {
	ID          string          `json:"id"`
	Products    map[Product]int `json:"products"`
	TotalAmount float64         `json:"total_amount"`
}

type Product struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}
