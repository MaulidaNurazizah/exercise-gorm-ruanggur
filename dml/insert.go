package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Connect() (*sql.DB, error) {
	dns := "host=localhost user=postgres password=Kmzway87a@ dbname=test_db_camp port=5432 sslmode=disable"

	db, err := sql.Open("postgres", dns)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func main() {
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	_, err = db.Exec(`INSERT INTO employees
	VALUES (1, 'Rizki', 25, 'Jl. Kebon Jeruk', 200000),
	(2, 'Andi', 27, 'Jl. Kebon Sirih', 300000),
	(3, 'Budi', 30, 'Jl. Kebon Melati', 400000 ),
	(4, 'Caca', 32, 'Jl. Kebon Anggrek', 500000),
	(5, 'Deni', 35, 'Jl. Kebon Mawar', 600000)`)
	if err != nil {
		panic(err)
	}
	fmt.Println("Insert data succes")
}
