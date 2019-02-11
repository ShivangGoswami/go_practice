package main

import (
	"database/sql"
	"log"

	_ "github.com/go-goracle/goracle"
)

func main() {
	db, err := sql.Open("goracle", "SYSTEM/1001289@localhost:1521/xe")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var (
		id   int
		name string
	)
	rows, err := db.Query("select * from prg")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		log.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
