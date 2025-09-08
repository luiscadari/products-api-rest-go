package models

import "github.com/luiscadari/products-api-rest-go/db"

type Product struct {
	Nome string
	Descricao string
	Preco float64
	Quantidade int
}

func GetProducts()[]Product{
	db := db.ConnectBD()
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
	defer db.Close()
	return products
}