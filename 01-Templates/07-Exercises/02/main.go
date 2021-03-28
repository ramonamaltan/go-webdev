// 1. Create a data structure to pass to a template which
// * contains information about California hotels including Name, Address, City, Zip, Region
// * region can be: Southern, Central, Northern
// * can hold an unlimited number of hotels

package main

import (
	"html/template"
	"log"
	"os"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseFiles("tpl.gohtml"))
}

// type hotel struct {
// 	Name, Address, City, Zip, Region string
// }

type hotel struct {
	Name    string
	Address string
	City    string
	Zip     string
	Region  string
}

type hotels []hotel

func main() {
	h := hotels{
		hotel{
			"PalmSprings",
			"Palm Valley",
			"Palmerama",
			"0379Palm",
			"Southern",
		},
		hotel{
			"Hotel Mama",
			"Koloniestr",
			"Berlin",
			"13359",
			"Northern",
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", h)
	if err != nil {
		log.Fatalln(err)
	}
}
