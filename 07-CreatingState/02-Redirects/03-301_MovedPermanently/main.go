package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/bar", bar)
	http.HandleFunc("/barred", barred)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func foo(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at foo", req.Method, "\n\n")
}

func bar(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at bar", req.Method)
	// w.Header().Set("Location", "/")
	// w.WriteHeader(http.StatusPermanentRedirect)
	// the above is same as
	http.Redirect(w, req, "/", http.StatusPermanentRedirect)
}

func barred(w http.ResponseWriter, req *http.Request) {
	fmt.Println("Your request method at barred", req.Method)
	tpl.ExecuteTemplate(w, "index.gohtml", nil)
}
