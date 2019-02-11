package main

import (
	"log"
	"net/http"
	"strconv"
	"strings"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}
		cycles := 5
		for ind, data := range r.Form {
			if strings.EqualFold(ind, "cycles") {
				cycles, _ = strconv.Atoi(strings.Join(data, ""))
				// if err != nil {
				// 	log.Fatal(err)
				// }
			}
		}
		lissajous(w, cycles)
	})
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}
