package handler

import (
	"checkout_service/cmd/api/request"
	"checkout_service/cmd/api/response"
	"checkout_service/internal/checkout"
	"encoding/json"
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

const (
	_basketIdParam = "basket_id"
)

type CheckoutHandler struct {
	checkoutService checkout.Service
}

func NewCheckoutHandler(service checkout.Service) CheckoutHandler {
	return CheckoutHandler{
		checkoutService: service,
	}
}

func (s *CheckoutHandler) GetBasketTotalAmount(w http.ResponseWriter, r *http.Request) {
	basketId := chi.URLParam(r, _basketIdParam)
	if len(basketId) == 0 {
		setResponse(w, http.StatusBadRequest, response.Error{Code: http.StatusBadRequest, Message: "Url param basket_id is missing"})
		return
	}

	totalAmount, err := s.checkoutService.GetBasketTotalAmount(basketId)
	if err != nil {
		setResponse(w, http.StatusBadRequest, response.Error{Code: http.StatusNotFound, Message: "Basket not found"})
		return
	}

	setResponse(w, http.StatusOK, response.GetBasketAmount{BasketID: basketId, TotalAmount: totalAmount})
}

func (s *CheckoutHandler) CreateBasket(w http.ResponseWriter, r *http.Request) {
	basket := s.checkoutService.CreateBasket()
	setResponse(w, http.StatusCreated, response.CreateBasket{BasketID: basket.ID})
	return
}

func (s *CheckoutHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	addProductRequest := request.AddProduct{}
	err := json.NewDecoder(r.Body).Decode(&addProductRequest)
	if err != nil {
		log.Println(err.Error())
		setResponse(w, http.StatusBadRequest, response.Error{Code: http.StatusBadRequest, Message: "Invalid request body"})
		return
	}

	basketUpdated, err := s.checkoutService.AddProduct(addProductRequest.BasketID, addProductRequest.ProductCode, addProductRequest.Quantity)
	if err != nil {
		setResponse(w, http.StatusInternalServerError, response.Error{Code: http.StatusInternalServerError, Message: "Problems to add new product"})
		return
	}

	setResponse(w, http.StatusOK, response.AddProduct{BasketID: basketUpdated.ID, Products: basketUpdated.Products})
}

func (s *CheckoutHandler) DeleteBasket(w http.ResponseWriter, r *http.Request) {
	basketId := chi.URLParam(r, _basketIdParam)
	if len(basketId) == 0 {
		setResponse(w, http.StatusBadRequest, response.Error{Code: http.StatusBadRequest, Message: "Url param basket_id is missing"})
		return
	}

	err := s.checkoutService.DeleteBasket(basketId)
	if err != nil {
		log.Println(err.Error())
		setResponse(w, http.StatusBadRequest, response.Error{Code: http.StatusBadRequest, Message: "basket_id not found"})
		return
	}

	setResponse(w, http.StatusNoContent, nil)
}
