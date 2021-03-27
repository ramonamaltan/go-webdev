package main

import (
	"log"
	"os"
	"text/template"
)

// declare here to use throughout all functions
var tpl *template.Template

func init() {
	// must does error checking
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", `Release self-focus; embrace other-focus.`)
	if err != nil {
		log.Fatalln(err)
	}
}
