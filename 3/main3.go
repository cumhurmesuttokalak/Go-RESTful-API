package main

import (
	"fmt"
	"io"
	"net/http"
)

type ironman int

// ServeHTTP func ı özel bir func kafadan uydurulmuş bir func ismi değil.
func (x ironman) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. IRON!.")
}

type wolverine int

func (x wolverine) ServeHTTP(res http.ResponseWriter, r *http.Request) {
	io.WriteString(res, "Mr. WOLVERİNE!.")
}

func main() {
	var i ironman
	var w wolverine

	mux := http.NewServeMux()
	mux.Handle("/ironman", i)
	mux.Handle("/wolverine", w)

	http.ListenAndServe(":8080", mux)
	fmt.Println("3")
}
