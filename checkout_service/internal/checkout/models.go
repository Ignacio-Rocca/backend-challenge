package checkout

import "github.com/google/uuid"

type Basket struct {
	ID          string         `json:"id"`
	Products    map[string]int `json:"products"`
	TotalAmount float64        `json:"total_amount"`
}

type Product struct {
	Code  string  `json:"code"`
	Name  string  `json:"name"`
	Price float64 `json:"price"`
}

func NewEmptyBasket() Basket {
	return Basket{
		ID:          uuid.NewString(),
		Products:    make(map[string]int),
		TotalAmount: 0,
	}
}
