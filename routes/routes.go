package routes

import (
	"net/http"

	"github.com/luiscadari/products-api-rest-go/controller"
)

func GetRoutes(){
	http.HandleFunc("/", controller.GetProducts)
	http.HandleFunc("/products/create", controller.NewProducts)
	http.HandleFunc("/delete", controller.DeleteProduct)
}