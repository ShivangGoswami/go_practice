package main

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	_ "github.com/go-goracle/goracle"
)

type googleplaystore struct {
	app           sql.NullString
	category      sql.NullString
	rating        sql.NullFloat64
	reviews       sql.NullInt64
	size          sql.NullString
	installs      sql.NullString
	typeApp       sql.NullString
	price         sql.NullInt64
	contentRating sql.NullString
	genres        sql.NullString
	lastUpdated   sql.NullString
	currentVer    sql.NullString
	androidVer    sql.NullString
}

type jsonplaystore struct {
	app           string  `json:"app"`
	category      string  `json:"category"`
	rating        float64 `json:"rating"`
	reviews       int64   `json:"reviews"`
	size          string  `json:"size"`
	installs      string  `json:"installs"`
	typeApp       string  `json:"type"`
	price         int64   `json:"price"`
	contentRating string  `json:"content_rating"`
	genres        string  `json:"genres"`
	lastUpdated   string  `json:"last_updated"`
	currentVer    string  `json:"current_version"`
	androidVer    string  `json:"Android_version"`
}

type rows []googleplaystore

func checkerr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
func getallrows() rows {
	var datarows rows
	db, err := sql.Open("goracle", "SYSTEM/1001289@localhost:1521/xe")
	err = db.Ping()
	checkerr(err)
	rows, err := db.Query("select * from googleplaystore")
	checkerr(err)
	defer rows.Close()
	for rows.Next() {
		var ele googleplaystore
		err := rows.Scan(&ele.app, &ele.category, &ele.rating, &ele.reviews, &ele.size, &ele.installs, &ele.typeApp, &ele.price, &ele.contentRating, &ele.genres, &ele.lastUpdated, &ele.currentVer, &ele.androidVer)
		checkerr(err)
		datarows = append(datarows, ele)
		//fmt.Println(datarows)
	}
	err = rows.Err()
	checkerr(err)
	return datarows
}

type jsonrows []jsonplaystore

func getjsondata(data rows) jsonrows {
	var json jsonrows
	for _, ele := range data {
		var jsonele jsonplaystore
		if ele.app.Valid {
			jsonele.app = ele.app.String
		} else {
			jsonele.app = ""
		}
		if ele.category.Valid {
			jsonele.category = ele.category.String
		} else {
			jsonele.category = ""
		}
		if ele.rating.Valid {
			jsonele.rating = ele.rating.Float64
		} else {
			jsonele.rating = 0
		}
		if ele.reviews.Valid {
			jsonele.reviews = ele.reviews.Int64
		} else {
			jsonele.reviews = 0
		}
		if ele.size.Valid {
			jsonele.size = ele.size.String
		} else {
			jsonele.size = ""
		}
		if ele.installs.Valid {
			jsonele.installs = ele.installs.String
		} else {
			jsonele.installs = ""
		}
		if ele.typeApp.Valid {
			jsonele.typeApp = ele.typeApp.String
		} else {
			jsonele.typeApp = ""
		}
		if ele.price.Valid {
			jsonele.price = ele.price.Int64
		} else {
			jsonele.price = 0
		}
		if ele.contentRating.Valid {
			jsonele.contentRating = ele.contentRating.String
		} else {
			jsonele.contentRating = ""
		}
		if ele.genres.Valid {
			jsonele.genres = ele.genres.String
		} else {
			jsonele.genres = ""
		}
		if ele.lastUpdated.Valid {
			jsonele.lastUpdated = ele.lastUpdated.String
		} else {
			jsonele.lastUpdated = ""
		}
		if ele.currentVer.Valid {
			jsonele.currentVer = ele.currentVer.String
		} else {
			jsonele.currentVer = ""
		}
		if ele.androidVer.Valid {
			jsonele.androidVer = ele.androidVer.String
		} else {
			jsonele.androidVer = ""
		}
		json = append(json, jsonele)
	}
	return json
}
func allrows(w http.ResponseWriter, r *http.Request) {
	data := getallrows()
	jdata := getjsondata(data)
	jdata = jdata[0:3]
	json.NewEncoder(w).Encode(jdata)
}

func handleRequest() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/getgooglestore", allrows).Methods("GET")
	log.Fatal(http.ListenAndServe(":1234", myRouter))
}
func main() {
	handleRequest()
}
