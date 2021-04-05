package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/webdev-training/17-mongoDB/06-CRUD/controllers"
	"gopkg.in/mgo.v2"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController(getSession())
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}

func getSession() *mgo.Session {
	// connect to local mongo by default 27017
	s, err := mgo.Dial("mongodb://localhost")

	// check if connection error/ is mongo running?
	if err != nil {
		panic(err)
	}
	return s
}
