package main

import (
	"database/sql"

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

}
