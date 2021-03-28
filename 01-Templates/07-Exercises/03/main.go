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

type ingredient struct {
	Name   string
	Amount float64
	Unit   string
}

type breakfast struct {
	Name        string
	Ingredients []ingredient
}

type lunch struct {
	Name        string
	Ingredients []ingredient
}

type dinner struct {
	Name        string
	Ingredients []ingredient
}

type menu struct {
	Breakfast breakfast
	Lunch     lunch
	Dinner    dinner
}

func main() {
	m := menu{
		Breakfast: breakfast{
			Name: "porridge",
			Ingredients: []ingredient{
				ingredient{
					Name:   "milk",
					Amount: 0.2,
					Unit:   "liters",
				},
				ingredient{
					Name:   "oats",
					Amount: 40,
					Unit:   "gramms",
				},
			},
		},
		Lunch: lunch{
			Name: "salad",
			Ingredients: []ingredient{
				ingredient{
					Name:   "dressing",
					Amount: 0.02,
					Unit:   "liters",
				},
				ingredient{
					Name:   "salad",
					Amount: 400,
					Unit:   "gramms",
				},
			},
		},
		Dinner: dinner{
			Name: "salad",
			Ingredients: []ingredient{
				ingredient{
					Name:   "dressing",
					Amount: 0.02,
					Unit:   "liters",
				},
				ingredient{
					Name:   "salad",
					Amount: 400,
					Unit:   "gramms",
				},
			},
		},
	}

	err := tpl.ExecuteTemplate(os.Stdout, "tpl.gohtml", m)
	if err != nil {
		log.Fatalln(err)
	}
}
