package main

import (
	"checkout_service/cmd/api/handler"
	"checkout_service/internal/checkout"
	"checkout_service/internal/repository/memorydb"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

func main() {
	router := chi.NewRouter()

	checkoutService := checkout.NewService(memorydb.NewDB())
	checkoutHandler := handler.NewCheckoutHandler(checkoutService)

	setRoutes(router, checkoutHandler)

	log.Println("Server listening on http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", router))
}
