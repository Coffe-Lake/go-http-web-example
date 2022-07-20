package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

const file string = "store.db"

type product struct {
	id      int
	model   string
	company string
	price   int
}

func main() {
	db, err := sql.Open("sqlite3", file)
	if err != nil {
		fmt.Println(db)
		panic(err)
	}
	defer db.Close()

	result, err := db.Exec("insert into products (model, company, price) values ('iPhone X', $1, $2)",
		"Apple", 72000)
	if err != nil {
		panic(err)
	}
	fmt.Print("ID последнего добавленного продукта: ")
	fmt.Println(result.LastInsertId()) // id последнего добавленного объекта
	fmt.Print("Количество добавленных продуктов: ")
	fmt.Println(result.RowsAffected()) // количество добавленных строк

	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	products := []product{}

	for rows.Next() {
		p := product{}
		err := rows.Scan(&p.id, &p.model, &p.company, &p.price)
		if err != nil {
			fmt.Println(err)
			continue
		}
		products = append(products, p)
	}
	for _, p := range products {
		fmt.Println(p.id, p.model, p.company, p.price)
	}

}
