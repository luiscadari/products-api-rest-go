package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/luiscadari/products-api-rest-go/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func GetProducts(w http.ResponseWriter, r *http.Request){

	products := models.GetProducts();
	templates.ExecuteTemplate(w, "Index", products)
}

func NewProducts(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		var product models.Product
		err := r.ParseForm()
		if err != nil {
			http.Error(w, "Unable to parse form", http.StatusBadRequest)
			return
		}
		product.Nome = r.FormValue("nome")
		product.Descricao = r.FormValue("descricao")
		preco, _ := strconv.ParseFloat(r.FormValue("preco"), 64)
		quantidade, _ := strconv.Atoi(r.FormValue("quantidade"))
		product.Preco = preco
		product.Quantidade = quantidade
		models.CreateProduct(product)
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	templates.ExecuteTemplate(w, "CreateProduct", nil)
}