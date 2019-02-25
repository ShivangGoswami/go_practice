package main

import (
	"log"
	"net/http"

	"./userservice"

	restful "github.com/emicklei/go-restful"
)

func main() {
	restful.Add(userservice.New())
	log.Fatal(http.ListenAndServe(":8000", nil))
}
