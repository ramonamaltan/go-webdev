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
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	err := tpl.ExecuteTemplate(os.Stdout, "vespa.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "two.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

	err = tpl.ExecuteTemplate(os.Stdout, "one.gohtml", nil)
	if err != nil {
		log.Fatalln(err)
	}

}
