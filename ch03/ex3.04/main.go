package main

import (
	"log"
	"net/http"

	"github.com/go-book-gopl/ch03/ex3.04/plotsine"
)

func main() {
	handler := func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "image/svg+xml")
		plotsine.Plotter(w)
	}

	http.HandleFunc("/sine", handler)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
	return
}
