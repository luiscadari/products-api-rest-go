package main

import (
	"net/http"

	"github.com/luiscadari/products-api-rest-go/routes"
)

func main() {
	routes.GetRoutes()
	http.ListenAndServe(":8000", nil)
}