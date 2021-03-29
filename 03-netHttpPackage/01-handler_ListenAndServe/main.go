package main

import (
	"fmt"
	"net/http"
)

type hotdog int

// hotdog is now also a handler
// it implicitly implements the Handler interface
// any type with the method ServeHTTP(Responsewriter, *Request) is a Handler
func (m hotdog) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Any code")
}

func main() {
	var d hotdog
	// ListenAndServe starts an HTTP server with a given address and handler
	http.ListenAndServe(":8080", d)
}
