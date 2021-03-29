package main

import (
	"io"
	"net/http"
)

// type hotdog int
// func (d hotdog) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "dog, dog, dog")
// }

func d(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "dog, dog, dog")
}

// type hotcat int
// func (c hotcat) ServeHTTP(w http.ResponseWriter, req *http.Request) {
// 	io.WriteString(w, "cat, cat, cat")
// }

func c(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "cat, cat, cat")
}

func main() {
	// var d hotdog
	// var c hotcat
	// http.Handle("/dog/", d)
	// http.Handle("/cat/", c)

	// instead of http Handle use http Handle func and no handler needed with ServeHTTP method but a func(Responsewriter, *Request)
	http.HandleFunc("/dog/", d)
	http.HandleFunc("/cat/", c)

	// use DefaultServeMux -> pass in nil instead of mux
	http.ListenAndServe(":8080", nil)
}
