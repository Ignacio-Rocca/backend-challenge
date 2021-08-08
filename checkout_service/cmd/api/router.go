package main

import (
	"checkout_service/cmd/api/handler"
	"github.com/go-chi/chi"
)

func setRoutes(router *chi.Mux, checkoutHandler handler.CheckoutHandler) {
	// todo: define middlewares for auth and validation
	router.Get("/checkout/basket/{basket_id}/amount", checkoutHandler.GetBasketTotalAmount)
	router.Post("/checkout/create", checkoutHandler.CreateBasket)
	router.Put("/checkout/basket/add_product", checkoutHandler.AddProduct)
	router.Delete("/checkout/basket/{basket_id}", checkoutHandler.DeleteBasket)
}
