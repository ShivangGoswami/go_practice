package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Article struct {
	Title   string `json:"Title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

type Articles []Article

func allArticles(w http.ResponseWriter, r *http.Request) {
	articles := Articles{
		Article{Title: "Test Title", Desc: "Test Description", Content: "Hello World!"},
		Article{Title: "500 per month", Desc: "Netflix", Content: "Boss Baby"},
	}
	fmt.Println("Endpoint Hit:All Articles Endpoint")
	json.NewEncoder(w).Encode(articles)
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Homepage Endpoint hit")
}
func testPostArticles(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Test POST endpoint worked")
}
func handleRequests() {
	// myRouter := mux.NewRouter().StrictSlash(true)
	// myRouter.HandleFunc("/", homePage)
	// myRouter.HandleFunc("/articles", allArticles).Methods("GET")
	// myRouter.HandleFunc("/articles", testPostArticles).Methods("POST")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/articles", allArticles)
	log.Fatal(http.ListenAndServe(":1234", nil))
	//log.Fatal(http.ListenAndServe(":1234", myRouter))
}
func main() {
	handleRequests()
}
