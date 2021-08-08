package main

import (
	"checkout_client/request"
	"checkout_client/response"
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	client := &http.Client{}

	setFirstExample(client)
	setSecondExample(client)
	setThirdExample(client)
	setFourthExample(client)
}

func setFirstExample(client *http.Client) {
	basket := createBasket(client)

	addProduct(client, basket.BasketID, "PEN", 1)
	addProduct(client, basket.BasketID, "TSHIRT", 1)
	addProduct(client, basket.BasketID, "MUG", 1)

	validateTotalAmount(client, basket.BasketID, 32.50)
}

func setSecondExample(client *http.Client) {
	basket := createBasket(client)

	addProduct(client, basket.BasketID, "PEN", 2)
	addProduct(client, basket.BasketID, "TSHIRT", 1)

	validateTotalAmount(client, basket.BasketID, 25.00)
}
func setThirdExample(client *http.Client) {
	basket := createBasket(client)

	addProduct(client, basket.BasketID, "PEN", 1)
	addProduct(client, basket.BasketID, "TSHIRT", 4)

	validateTotalAmount(client, basket.BasketID, 65.00)
}

func setFourthExample(client *http.Client) {
	basket := createBasket(client)

	addProduct(client, basket.BasketID, "PEN", 3)
	addProduct(client, basket.BasketID, "TSHIRT", 3)
	addProduct(client, basket.BasketID, "MUG", 1)

	validateTotalAmount(client, basket.BasketID, 62.5)
}

func createBasket(client *http.Client) response.CreateBasket {
	createCheckoutRequest, err := http.NewRequest(http.MethodPost, "http://localhost:8080/checkout/create", nil)
	isOkRequest(err)

	resp, err := client.Do(createCheckoutRequest)
	isOkResponse(resp, err)

	basket := response.CreateBasket{}
	parseResponse(resp, &basket)

	return basket
}

func addProduct(client *http.Client, basketID, productCode string, quantity int) {
	url := "http://localhost:8080/checkout/basket/add_product"
	body := request.AddProduct{
		BasketID:    basketID,
		ProductCode: productCode,
		Quantity:    quantity,
	}
	requestAddProduct, err := http.NewRequest(http.MethodPut, url, buildBody(body))
	isOkRequest(err)

	resp, err := client.Do(requestAddProduct)
	isOkResponse(resp, err)
}

func validateTotalAmount(client *http.Client, basketID string, amountExpected float64) {
	url := fmt.Sprintf("http://localhost:8080/checkout/basket/%s/amount", basketID)
	requestGetTotalAmount, err := http.NewRequest(http.MethodGet, url, nil)
	isOkRequest(err)

	resp, err := client.Do(requestGetTotalAmount)
	isOkResponse(resp, err)

	basketAmount := response.BasketAmount{}
	parseResponse(resp, &basketAmount)

	if basketAmount.TotalAmount != amountExpected {
		log.Println("First basket total amount is invalid")
		os.Exit(1)
	}
}
