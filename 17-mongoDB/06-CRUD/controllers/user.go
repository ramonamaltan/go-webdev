package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/webdev-training/17-mongoDB/06-CRUD/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type UserController struct {
	session *mgo.Session
}

func NewUserController(s *mgo.Session) *UserController {
	return &UserController{s}
}

// methods are now capitalized to be exported
func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	// grab id
	id := p.ByName("id")

	// verify id is ObjectId hex representation otherwise status not found
	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	// ObjectIDHEx returns an ObjectId from the provided hex representation
	oid := bson.ObjectIdHex(id)

	u := models.User{}

	// fetch user
	if err := uc.session.DB("go-webdev-db").C("users").FindId(oid).One(&u); err != nil {
		w.WriteHeader(404)
		return
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

	// create bson Id
	u.Id = bson.NewObjectId()

	// store the user in mongodb
	uc.session.DB("go-webdev-db").C("users").Insert(u)

	// marshal/unmarshal for having JSON assigned to a variable
	uj, _ := json.Marshal(u)

	// write content type, status code and payload
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", uj)
}

func (uc UserController) DeleteUser(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id) {
		w.WriteHeader(404)
		return
	}

	oid := bson.ObjectIdHex(id)

	// delete user
	if err := uc.session.DB("go-webdev-db").C("users").RemoveId(oid); err != nil {
		w.WriteHeader(404)
		return
	}

	w.WriteHeader(http.StatusOK) // 200
	fmt.Fprint(w, "Deleted user", oid, "\n")
}
