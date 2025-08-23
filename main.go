package main

import (
	"html/template"
	"net/http"
)

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
	produtos := []Product{
		{
			Nome: "Camiseta",
			Descricao: "Camiseta Amarela",
			Preco: 80.50,
			Quantidade: 10,
		},
		{"Tenis", "Tenis Adidas branco", 34.00, 6},
		{"Calça de Alfaiataria", "Calça tipo de alfaiataria preta.", 100.00, 10},
	}
	templates.ExecuteTemplate(w, "Index", produtos)
}