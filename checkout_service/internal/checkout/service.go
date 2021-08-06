package checkout

import (
	"errors"
	"log"
)

var (
	ErrBasketNotFound = errors.New("Basket not found")
	ErrInvalidBasketId = errors.New("Invalid basket_id")
	ErrInvalidProduct = errors.New("Invalid product")
)

type Service interface {
	GetBasketTotalAmount(basketId string) (float64, error)
	CreateBasket() (Basket, error)
	AddProduct(basketId, productCode string) error
	DeleteBasket(basketId string) error
}

type service struct {

}

func NewService() *service {
	return &service{}
}

func (s *service) GetBasketTotalAmount(basketId string) (float64, error) {
	basket, ok := checkoutDB[basketId]
	if !ok {
		log.Println(ErrBasketNotFound.Error())
		return 0, ErrBasketNotFound
	}

	totalAmount := 0.00

	for product, quantity := range basket.Products {
		switch product.Code {
		case _penCode:
			buy2Get1Free(product.Price, quantity)
		case _tshirtCode:
			if quantity >= 3 {
				totalAmount +=  product.Price * float64(quantity) * _tshirtDiscount
			} else {
				totalAmount +=  product.Price * float64(quantity)
			}
		case _mugCode:
			totalAmount += product.Price * float64(quantity)
		}
	}

	return totalAmount, nil
}

func (s *service) CreateBasket() (Basket, error) {
	return Basket{}, nil
}

func (s *service) AddProduct(basketId, productCode string) error {
	basket, ok := checkoutDB[basketId]
	if !ok {
		log.Println(ErrInvalidBasketId.Error())
		return ErrInvalidBasketId
	}

	product, err := productsDB.GetProductByCode(productCode)
	if err != nil {
		log.Println(err.Error())
		return err
	}

	basket.Products[product] += 1
	return nil
}

func (s *service) DeleteBasket(basketId string) error {
	if _, ok := checkoutDB[basketId]; !ok {
		return ErrInvalidBasketId
	}
	delete(checkoutDB, basketId)
	return nil
}

func buy2Get1Free(price float64, quantity int) float64 {
	total := 0.00
	if quantity % 2 == 0 {
		multiplier := quantity / 2
		total +=  price * float64(multiplier)
	} else {
		multiplier := (quantity / 2) + 1
		total +=  price * float64(multiplier)
	}
	return total
}