package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-goracle/goracle"
)

func main() {
	db, err := sql.Open("goracle", "SYSTEM/1001289@localhost:1521/xe")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	var (
		id   int
		name string
	)
	rows, err := db.Query("select intval,strval from prg")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(id, name)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}
}
