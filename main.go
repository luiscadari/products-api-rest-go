package main

import (
	"html/template"
	"net/http"

	"github.com/luiscadari/products-api-rest-go/models"
)




var templates = template.Must(template.ParseGlob("templates/*html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){

	products := models.GetProducts();
	templates.ExecuteTemplate(w, "Index", products)
}