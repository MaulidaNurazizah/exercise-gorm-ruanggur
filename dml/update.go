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

	_, err = db.Exec(`UPDATE employess SET salary = 100000 WHERE id =1 `)
	if err != nil {
		panic(err)
	}

	fmt.Println("Update data success")

	_, err = db.Exec(`DELETE FROM employess WHERE id = 5 `)
	if err != nil {
		panic(err)
	}
}
