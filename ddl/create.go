package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "maulida1205"
	dbname   = "test_db_camp"
)

func ConnectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	return db, nil
}
func main() {
	db, err := ConnectDB()
	if err != nil {
		panic(err)
	}

	defer db.Close()

	_, err = db.Exec(`CREATE TABLE employee (
		id INT,
		name VARCHAR(255) NOT NULL,
		age INT NOT NULL,
		address VARCHAR(255),
		salary INT
	  )`)

	if err != nil {
		panic(err)
	}

	fmt.Println("Table employee created")
}
