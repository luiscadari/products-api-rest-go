package main

import (
	"database/sql"
	"html/template"
	"net/http"

	_ "github.com/lib/pq"
)

func connectBD() *sql.DB {
	conexao := "user=postgres dbname=alura_loja password=123456 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil{
		panic(err.Error())
	}
	return db
}

type Product struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

var templates = template.Must(template.ParseGlob("templates/*html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request){
	db := connectBD()
	getProducts, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	} 
	product := Product{}
	products := []Product{}

	for getProducts.Next(){
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = getProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		product.Nome = nome
		product.Descricao = descricao
		product.Preco = preco
		product.Quantidade = quantidade
		products = append(products, product)
	}

	templates.ExecuteTemplate(w, "Index", products)
	defer db.Close()
}