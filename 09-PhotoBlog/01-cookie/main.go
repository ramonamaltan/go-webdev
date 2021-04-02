package main

import (
	"html/template"
	"net/http"

	"github.com/google/uuid"
)

type user struct {
	UserName  string
	Password  string
	FirstName string
	LastName  string
}

var tpl *template.Template
var dbUsers = map[string]user{}

func init() {
	tpl = template.Must(template.ParseGlob("templates/*"))
}

func main() {
	http.HandleFunc("/", index)
	http.Handle("/favicon.ico", http.NotFoundHandler())
	http.ListenAndServe(":8080", nil)
}

func index(w http.ResponseWriter, req *http.Request) {
	c := getCookie(w, req)
	tpl.ExecuteTemplate(w, "index.gohtml", c.Value)
}

func getCookie(w http.ResponseWriter, req *http.Request) *http.Cookie {
	c, err := req.Cookie("session")
	if err != nil {
		sID := uuid.New()
		c = &http.Cookie{
			Name:  "user",
			Value: sID.String(),
		}
		http.SetCookie(w, c)
	}
	return c
}
