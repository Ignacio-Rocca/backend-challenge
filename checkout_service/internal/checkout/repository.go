package checkout

import "checkout_service/internal/product"

type Repository interface {
	GetProductByCode(code string) (product.Product, error)
	AddNewBasket() Basket
	DeleteBasket(basketID string) error
	GetBasketByID(basketID string) (*Basket, error)
	AddProductToBasket(basketId, productCode string, quantity int) (Basket, error)
}
