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
func singlerowerror() {
	var name string
	err := db.QueryRow("select strval from prg where intval=:p1", 90).Scan(&name)
	if err != nil {
		if err == sql.ErrNoRows {
			log.Println("no rows fetched!!")
			//log.Println(name == "")
		} else {
			log.Fatal(err)
		}
	}
}
func null1() {
	rows, err := db.Query("select strval from prg")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var s sql.NullString
		err := rows.Scan(&s)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(s)
		if s.Valid {
			log.Println(s.String)
		} else {
			log.Println("null value")
		}
	}
}
func null2() {
	rows, err := db.Query("select q3,q1,q2 from pro")
	defer rows.Close()
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var s1 sql.NullFloat64
		var s2 sql.NullInt64
		var s3 sql.NullString
		err := rows.Scan(&s1, &s2, &s3)
		if err != nil {
			log.Fatal(err)
		}
		if s1.Valid && s2.Valid && s3.String != "" {
			log.Println(s1.Float64, s2.Int64, s3.String)
		} else {
			log.Println("null values")
		}
	}
	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}
}
func main() {
	connect()
	defer db.Close()
	//singlerowerror()
	null2()
}
