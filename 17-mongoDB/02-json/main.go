package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/webdev-training/17-mongoDB/02-json/models"
)

func main() {
	r := httprouter.New()
	r.GET("/", index)
	r.GET("/user/:id", getUser)
	http.ListenAndServe("localhost:8080", r)
}

func index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	s := `<!DOCTYPE html>
	<html lang="en">
	<head>
	<meta charset="UTF-8">
	<title>Index</title>
	</head>
	<body>
	<a href="/user/9872309847">GO TO: http://localhost:8080/user/9872309847</a>
	</body>
	</html>`

	w.Header().Set("Content-Tpye", "text/html; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(s))
}

func getUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// in future pull user out of db
	u := models.User{
		Name:   "James Bond",
		Gender: "Male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	// Marshal into JSON
	uj, _ := json.Marshal(u)

	// write content-type, statuscode, payload
	w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}
