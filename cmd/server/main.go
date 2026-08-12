package main

import (
	"log"
	"net/http"

	"github.com/diffpal/demo/internal/httpapi"
	"github.com/diffpal/demo/internal/orders"
)

func main() {
	handler := httpapi.New(orders.NewService(orders.NewMemoryRepository()))
	log.Println("Tiny Orders listening on :8080")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
