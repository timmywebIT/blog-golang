package main

import (
	"database/sql"
	"log"
)

func main() {

	db, error := sql.Open("sqlite3", "./data.db")
	if error != nil {
		log.Fatal(error)
	}

	defer db.Close()

}
