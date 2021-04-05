package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/webdev-training/17-mongoDB/04-controllers/controllers"
)

func main() {
	r := httprouter.New()
	uc := controllers.NewUserController()
	r.GET("/user/:id", uc.GetUser)
	r.POST("/user", uc.CreateUser)
	r.DELETE("/user/:id", uc.DeleteUser)
	http.ListenAndServe("localhost:8080", r)
}
