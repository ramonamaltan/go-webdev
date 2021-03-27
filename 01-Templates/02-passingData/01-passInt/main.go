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
	// pass in one piece of data into template
	// aggregate data structure (struct)
	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", 42)
	if err != nil {
		log.Fatalln(err)
	}
}
