package main

import (
	"fmt"
	"net/http"

	"github.com/luiscadari/products-api-rest-go/routes"
)

func main() {
	routes.GetRoutes()
	fmt.Println("Server is listening on port 8000")
	http.ListenAndServe(":8000", nil)
}