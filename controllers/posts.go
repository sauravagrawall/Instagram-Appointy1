package controllers

import (
	"fmt"
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/sauravagrawall/Instagram-Appointy/models"
	"net/http"
)

type PostsController struct{
	session *mgo.Session
}

func NewPostsController(s *mgo.Session) *PostsController{
	return &PostsController{s}
}

func (pc PostsController) GetPost (w http.ResponseWriter, r *http.Request, p httprouter.Params){
	id := p.ByName("id")

	if !bson.IsObjectIdHex(id){
		w.WriteHeader(http.StatusNotFound)
	}

	oid := bson.ObjectIdHex(id)

	ps := models.Posts{}

	if err := pc.session.DB("Instagram-Appointy").C("posts").FindId(oid).One(&ps); err != nil{
		w.WriteHeader(404)
		return
	}

	pj, err := json.Marshal(ps)
	if err!=nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "%s\n", pj)


}

func (pc PostsController) CreatePost (w http.ResponseWriter, r *http.Request, _ httprouter.Params){
	ps := models.Posts{}

	json.NewDecoder(r.Body).Decode(&ps)
	ps.Id = bson.NewObjectId()
	pc.session.DB("Instagram-Appointy").C("posts").Insert(ps)

	pj, err := json.Marshal(ps)
	
	if err !=nil{
		fmt.Println(err)
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "%s\n", pj)
}