package main

import (
	"log"
	"os"
	"text/template"
)

var tpl *template.Template

type sage struct {
	Name  string
	Motto string
}

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

func main() {
	buddha := sage{
		Name:  "Budda",
		Motto: "The belief of no belifes",
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", buddha)
	if err != nil {
		log.Fatalln(err)
	}
}
