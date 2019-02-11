package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-goracle/goracle"
)

var db *sql.DB

func connect() {
	var err error
	db, err = sql.Open("goracle", "SYSTEM/1001289@localhost:1521/xe")
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}

}
func fetch() {
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

func parameter() {
	var (
		id   int
		name string
	)
	stmt, err := db.Prepare("select intval,strval from prg where intval=:param1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	rows, err := stmt.Query(123)
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
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
func singlerow1() {
	var name string
	err := db.QueryRow("select strval from prg where intval=:p1", 123).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}
func singlerow2() {
	stmt, err := db.Prepare("select strval from prg where intval=:p1")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	var name string
	err = stmt.QueryRow(123).Scan(&name)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(name)
}

func main() {
	connect()
	defer db.Close()
	//fetch()
	//parameter()
	//singlerow1()
	singlerow2()
}
