package models

import (
	"fmt"

	"github.com/luiscadari/products-api-rest-go/db"
)

type Product struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func GetProducts() []Product {
	db := db.ConnectBD()
	getProducts, err := db.Query("SELECT * FROM produtos")
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	products := []Product{}

	for getProducts.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64
		err = getProducts.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}
		product.Id = id
		product.Nome = nome
		product.Descricao = descricao
		product.Preco = preco
		product.Quantidade = quantidade
		products = append(products, product)
	}
	defer db.Close()
	return products
}

func GetProductById(idProduto string) Product {
	db := db.ConnectBD()
	getProduct, err := db.Query("select * from produtos where id = $1", idProduto)
	if err != nil {
		panic(err.Error())
	}
	product := Product{}
	for getProduct.Next() {
	var id, quantidade int
	var nome, descricao string
	var preco float64
	err = getProduct.Scan(&id, &nome, &descricao, &preco, &quantidade)
	if err != nil {
		panic(err.Error())
	}
	product.Id = id
	product.Nome = nome
	product.Descricao = descricao
	product.Preco = preco
	product.Quantidade = quantidade
	}

	defer db.Close()
	return product
}

func CreateProduct(newProduct Product) Product {
	db := db.ConnectBD()
	// Verificando se o produto já existe
	query := "SELECT " + "*" + " FROM produtos WHERE produtos.nome = '" + newProduct.Nome + "'"
	getProducts, err := db.Query(query)
	if err != nil {
		panic(err.Error())
	}
	var product Product
	var products []Product
	for getProducts.Next() {
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
	if len(products) > 0 {
		panic("Product already exists")
	}

	query = fmt.Sprintf("INSERT INTO produtos (nome, descricao, preco, quantidade) VALUES('%s', '%s', %f, %d)", newProduct.Nome, newProduct.Descricao, newProduct.Preco, newProduct.Quantidade)
	_, err = db.Exec(query)
	if err != nil {
		panic(err.Error())
	}
	defer db.Close()
	return newProduct
}

func DeleteProduct(id string) {
	db := db.ConnectBD()
	produto, err := db.Prepare("delete from produtos where id = $1")
	if err != nil {
		panic(err.Error())
	}
	produto.Exec(id)
	defer db.Close()
}
