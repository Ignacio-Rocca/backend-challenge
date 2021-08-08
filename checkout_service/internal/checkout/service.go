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
	CreateBasket() Basket
	AddProduct(basketId, productCode string, quantity int) (Basket, error)
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

	for productCode, quantity := range basket.Products {
		switch productCode {
		case _penCode:
			totalAmount += calculatePenTotalAmount(quantity)
		case _tshirtCode:
			totalAmount += calculateTshirtTotalAmount(quantity)
		case _mugCode:
			totalAmount += calculateMugTotalAmount(quantity)
		}
	}

	return totalAmount, nil
}

func (s *service) CreateBasket() Basket {
	basket := NewEmptyBasket()
	_, exist := checkoutDB[basket.ID]
	for exist {
		basket = NewEmptyBasket()
		_, exist = checkoutDB[basket.ID]
	}
	checkoutDB[basket.ID] = basket
	return basket
}

func (s *service) AddProduct(basketId, productCode string, quantity int) (Basket, error) {
	basket, ok := checkoutDB[basketId]
	if !ok {
		log.Println(ErrInvalidBasketId.Error())
		return Basket{}, ErrInvalidBasketId
	}

	product, err := productsDB.GetProductByCode(productCode)
	if err != nil {
		log.Println(err.Error())
		return Basket{}, err
	}

	basket.Products[product.Code] += quantity
	return basket, nil
}

func (s *service) DeleteBasket(basketId string) error {
	if _, ok := checkoutDB[basketId]; !ok {
		return ErrInvalidBasketId
	}
	delete(checkoutDB, basketId)
	return nil
}


func calculatePenTotalAmount(quantity int) float64 {
	pen, err := productsDB.GetProductByCode(_penCode)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return buy2Get1Free(pen.Price, quantity)
}

func calculateTshirtTotalAmount(quantity int) float64 {
	tshirt, err := productsDB.GetProductByCode(_tshirtCode)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	if quantity >= 3 {
		return  tshirt.Price * float64(quantity) * _tshirtDiscount
	}

	return  tshirt.Price * float64(quantity)
}

func calculateMugTotalAmount(quantity int) float64 {
	mug, err := productsDB.GetProductByCode(_mugCode)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return mug.Price * float64(quantity)
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