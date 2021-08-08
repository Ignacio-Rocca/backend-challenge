package checkout

import (
	"checkout_service/internal/product"
	"log"
)

const _tshirtDiscount = 0.75

type Service interface {
	GetBasketTotalAmount(basketId string) (float64, error)
	CreateBasket() Basket
	AddProduct(basketId, productCode string, quantity int) (Basket, error)
	DeleteBasket(basketId string) error
}

type service struct {
	repo Repository
}

func NewService(repo Repository) *service {
	return &service{repo: repo}
}

func (s *service) GetBasketTotalAmount(basketId string) (float64, error) {
	basket, err := s.repo.GetBasketByID(basketId)
	if err != nil {
		log.Println(err.Error())
		return 0, err
	}

	totalAmount := 0.00

	for productCode, quantity := range basket.Products {
		switch productCode {
		case product.PenCode:
			totalAmount += s.calculatePenTotalAmount(quantity)
		case product.TshirtCode:
			totalAmount += s.calculateTshirtTotalAmount(quantity)
		case product.MugCode:
			totalAmount += s.calculateMugTotalAmount(quantity)
		}
	}

	return totalAmount, nil
}

func (s *service) CreateBasket() Basket {
	return s.repo.AddNewBasket()
}

func (s *service) AddProduct(basketId, productCode string, quantity int) (Basket, error) {
	return s.repo.AddProductToBasket(basketId, productCode, quantity)
}

func (s *service) DeleteBasket(basketId string) error {
	return s.repo.DeleteBasket(basketId)
}

func (s *service) calculatePenTotalAmount(quantity int) float64 {
	pen, err := s.repo.GetProductByCode(product.PenCode)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return buy2Get1Free(pen.Price, quantity)
}

func (s *service) calculateTshirtTotalAmount(quantity int) float64 {
	tshirt, err := s.repo.GetProductByCode(product.TshirtCode)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	if quantity >= 3 {
		return tshirt.Price * float64(quantity) * _tshirtDiscount
	}

	return tshirt.Price * float64(quantity)
}

func (s *service) calculateMugTotalAmount(quantity int) float64 {
	mug, err := s.repo.GetProductByCode(product.MugCode)
	if err != nil {
		log.Println(err.Error())
		return 0
	}

	return mug.Price * float64(quantity)
}

func buy2Get1Free(price float64, quantity int) float64 {
	total := 0.00
	if quantity%2 == 0 {
		multiplier := quantity / 2
		total += price * float64(multiplier)
	} else {
		multiplier := (quantity / 2) + 1
		total += price * float64(multiplier)
	}
	return total
}
