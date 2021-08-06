package handler

import (
	"checkout_service/internal/checkout"
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
		setResponse(w, http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "Url param basket_id is missing"})
		return
	}

	totalAmount, err := s.checkoutService.GetBasketTotalAmount(basketId)
	if err != nil {
		setResponse(w, http.StatusBadRequest, ErrorResponse{Code: http.StatusNotFound, Message: "Basket not found"})
		return
	}

	setResponse(w, http.StatusOK, BasketAmountResponse{BasketID: basketId, TotalAmount: totalAmount})
}

func (s *CheckoutHandler) CreateBasket(w http.ResponseWriter, r *http.Request) {
	s.checkoutService.CreateBasket()
	w.WriteHeader(http.StatusOK)
}

func (s *CheckoutHandler) AddProduct(w http.ResponseWriter, r *http.Request) {
	s.checkoutService.AddProduct()
}

func (s *CheckoutHandler) DeleteBasket(w http.ResponseWriter, r *http.Request) {
	basketId := chi.URLParam(r, _basketIdParam)
	if len(basketId) == 0 {
		setResponse(w, http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "Url param basket_id is missing"})
		return
	}

	err := s.checkoutService.DeleteBasket(basketId)
	if err != nil {
		log.Println(err.Error())
		setResponse(w, http.StatusBadRequest, ErrorResponse{Code: http.StatusBadRequest, Message: "basket_id not found"})
		return
	}

	setResponse(w, http.StatusNoContent, nil)
}
