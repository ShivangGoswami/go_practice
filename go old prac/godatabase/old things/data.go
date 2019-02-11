package main

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("goracle", "oracle://T1001289:T1001289@10.123.79.59:1521")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
}
