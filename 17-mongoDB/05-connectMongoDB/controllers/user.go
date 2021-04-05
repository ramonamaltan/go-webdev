package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/webdev-training/17-mongoDB/05-connectMongoDB/models"
	"gopkg.in/mgo.v2"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// methods are now capitalized to be exported
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{
		Name:   "James Bond",
		Gender: "male",
		Age:    32,
		Id:     p.ByName("id"),
	}

	uj, _ := json.Marshal(u)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) CreateUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	u := models.User{}

	// encode/decode for sending/receiving JSON to/from a stream
	json.NewDecoder(r.Body).Decode(&u)

	// change Id
	u.Id = "007"

	// marshal/unmarshal for having JSON assigned to a variable
	uj, _ := json.Marshal(u)

	// write content type, status code and payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// TODO: write code to delete user
	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Write code to delete user\n")
}
