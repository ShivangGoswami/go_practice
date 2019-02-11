package main

import (
	"database/sql"
	"log"
	"os"

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
func execstmt() {
	stmt, err := db.Prepare("INSERT INTO prg VALUES(:p1,:p2)")
	if err != nil {
		log.Fatal(err)
	}
	res, err := stmt.Exec("789", "Ali")
	if err != nil {
		log.Fatal(err)
	}
	// lastID, err := res.LastInsertId()
	// if err != nil {
	// 	log.Fatal(err)
	// }
	rowCnt, err := res.RowsAffected()
	if err != nil {
		log.Fatal(err)
	}
	//log.Printf("ID = %d ,affected = %d\n", lastID, rowCnt)
	log.Printf("affected = %d\n", rowCnt)
}
func transaction() {
	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare("INSERT INTO prg VALUES(:p1,:p2)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()
	for i := 0; i < 10; i++ {
		_, err := stmt.Exec(i, os.Args[1])
		if err != nil {
			log.Fatal(err)
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
	}
	//stmt.Close()
}
func main() {
	connect()
	defer db.Close()
	//execstmt()
	transaction()
}
