package routes

import (
	"net/http"

	"github.com/luiscadari/products-api-rest-go/controller"
)

func GetRoutes(){
	http.HandleFunc("/", controller.GetProducts)
}