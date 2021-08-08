package checkout

var (
	checkoutDB = initEmptyCheckoutDB()
	productsDB = initProductsDB()
)

const (
	// list of product codes
	_penCode    = "PEN"
	_tshirtCode = "TSHIRT"
	_mugCode    = "MUG"

	// discounts
	_tshirtDiscount = 0.75
)

type CheckoutDB map[string]Basket
type ProductsDB []Product

// initEmptyCheckout returns an empty repository mock
func initEmptyCheckoutDB() CheckoutDB {
	return make(map[string]Basket)
}

// initProductsDB returns an empty repository mock
func initProductsDB() ProductsDB {
	return []Product{
		{Code: _penCode, Name: "Lana Pen", Price: 5.00},
		{Code: _tshirtCode, Name: "Lana T-Shirt", Price: 20.00},
		{Code: _mugCode, Name: "Lana Coffee Mug", Price: 7.50},
	}
}

func (p ProductsDB) GetProductByCode(code string) (Product, error) {
	for _, product := range p {
		if product.Code == code {
			return product, nil
		}
	}
	return Product{}, ErrInvalidProduct
}
