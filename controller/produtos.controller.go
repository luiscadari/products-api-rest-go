package controller

import (
	"html/template"
	"net/http"

	"github.com/luiscadari/products-api-rest-go/models"
)

var templates = template.Must(template.ParseGlob("templates/*html"))

func GetProducts(w http.ResponseWriter, r *http.Request){

	products := models.GetProducts();
	templates.ExecuteTemplate(w, "Index", products)
}