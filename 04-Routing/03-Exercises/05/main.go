package main

import (
	"html/template"
	"log"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("templates/index.gohtml"))
}

func main() {
	f := http.StripPrefix("/public", http.FileServer(http.Dir("public")))
	http.Handle("/public/", f)
	http.HandleFunc("/", dogs)

	http.ListenAndServe(":8080", nil)
}

func dogs(w http.ResponseWriter, req *http.Request) {
	err := tpl.Execute(w, nil)
	if err != nil {
		log.Fatalln("Template couldn't execute", err)
	}
}
