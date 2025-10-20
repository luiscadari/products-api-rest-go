package controller

import (
	"html/template"
	"net/http"
	"strconv"

	"github.com/luiscadari/products-api-rest-go/db"
	"github.com/luiscadari/products-api-rest-go/models"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func GetProducts(w http.ResponseWriter, r *http.Request){

	products := models.GetProducts();
	templates.ExecuteTemplate(w, "Index", products)
}

func DeleteProduct(w http.ResponseWriter, r *http.Request){
	id := r.URL.Query().Get("id")
	models.DeleteProduct(id)
	http.Redirect(w, r, "/", 301)
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

func PutProduct(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if r.Method == "POST" {
		db := db.ConnectBD()
		newProduct, err := db.Prepare("UPDATE produtos set nome=$1, descricao=$2, preco=$3, quantidade=$4 WHERE id=$5")
		if err != nil {
			http.Error(w, "Unable to prepare statement", http.StatusInternalServerError)
			return
		}
		defer db.Close()
		_, err = newProduct.Exec(r.FormValue("nome"), r.FormValue("descricao"), r.FormValue("preco"), r.FormValue("quantidade"), r.FormValue("id"))
		if err != nil {
			http.Error(w, "Unable to execute statement", http.StatusInternalServerError)
			return
		}
		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}
	if r.Method == "GET" {
		product := models.GetProductById(id)
		templates.ExecuteTemplate(w, "PutProduct", product)
	}
}