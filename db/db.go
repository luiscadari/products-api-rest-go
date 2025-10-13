package db

import (
	"database/sql"

	_ "github.com/lib/pq"
)

func ConnectBD() *sql.DB {
	conexao := "user=postgres dbname=products password=1808 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil{
		panic(err.Error())
	}
	return db
}