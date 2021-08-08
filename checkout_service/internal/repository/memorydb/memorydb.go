package memorydb

import (
	"checkout_service/internal/checkout"
	"checkout_service/internal/product"
	"errors"
	"log"
)

type CheckoutDB map[string]checkout.Basket
type ProductsDB []product.Product

type MemoryDB struct {
	checkoutDB CheckoutDB
	productsDB ProductsDB
}

func NewDB() *MemoryDB {
	checkoutDB := initEmptyCheckoutDB()
	productsDB := initProductsDB()

	return &MemoryDB{
		checkoutDB: checkoutDB,
		productsDB: productsDB,
	}
}

var (
	ErrBasketNotFound  = errors.New("Basket not found")
	ErrInvalidBasketId = errors.New("Invalid basket_id")
	ErrInvalidProduct  = errors.New("Invalid product")
)

// initEmptyCheckout returns an empty repository mock
func initEmptyCheckoutDB() CheckoutDB {
	return make(map[string]checkout.Basket)
}

// initProductsDB returns an empty repository mock
func initProductsDB() ProductsDB {
	return []product.Product{
		{Code: product.PenCode, Name: "Lana Pen", Price: 5.00},
		{Code: product.TshirtCode, Name: "Lana T-Shirt", Price: 20.00},
		{Code: product.MugCode, Name: "Lana Coffee Mug", Price: 7.50},
	}
}

func (db *MemoryDB) GetProductByCode(code string) (product.Product, error) {
	for _, product := range db.productsDB {
		if product.Code == code {
			return product, nil
		}
	}
	return product.Product{}, ErrInvalidProduct
}

func (db *MemoryDB) AddNewBasket() checkout.Basket {
	basket := checkout.NewEmptyBasket()
	_, exist := db.checkoutDB[basket.ID]
	for exist {
		basket = checkout.NewEmptyBasket()
		_, exist = db.checkoutDB[basket.ID]
	}

	db.checkoutDB[basket.ID] = basket
	return basket
}

func (db *MemoryDB) GetBasketByID(basketID string) (*checkout.Basket, error) {
	basket, ok := db.checkoutDB[basketID]
	if !ok {
		log.Println(ErrBasketNotFound.Error())
		return nil, ErrBasketNotFound
	}
	return &basket, nil
}

func (db *MemoryDB) AddProductToBasket(basketID, productCode string, quantity int) (checkout.Basket, error) {
	basket, ok := db.checkoutDB[basketID]
	if !ok {
		log.Println(ErrBasketNotFound.Error())
		return checkout.Basket{}, ErrBasketNotFound
	}

	prod, err := db.GetProductByCode(productCode)
	if err != nil {
		log.Println(err.Error())
		return checkout.Basket{}, err
	}

	basket.Products[prod.Code] += quantity
	return basket, nil
}

func (db *MemoryDB) DeleteBasket(basketId string) error {
	if _, err := db.GetBasketByID(basketId); err != nil {
		return err
	}

	delete(db.checkoutDB, basketId)
	return nil
}
