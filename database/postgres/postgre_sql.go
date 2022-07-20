package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type client struct {
	id      int
	phone   string
	name    string
	age     int
	address string
}

func main() {
	connStr := "user=postgres password=postgres dbname=postgres_test sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	// result, err := db.Exec(
	// 	"INSERT INTO clients(phone, name, age, address) VALUES ('89690392263', 'Test', '20', 'Moskva, street Bibirevskaya')",
	// )

	// if err != nil {
	// 	fmt.Println(err)
	// }

	// fmt.Println(result.RowsAffected()) // количество добавленных строк

	response, err := db.Query("SELECT * FROM clients")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Close()

	clients := []client{}

	for response.Next() {
		c := client{}
		err := response.Scan(&c.id, &c.age, &c.phone, &c.name, &c.address)

		if err != nil {
			fmt.Println(err)
			continue
		}
		clients = append(clients, c)
	}
	for _, c := range clients {
		fmt.Println(c.id, c.age, c.phone, c.name, c.address)
	}
}
