package main

import (
	"ResumeMgmt/services"
	"gopkg.in/mgo.v2"
	"github.com/gorilla/mux"
	"net/http"
)

func main() {
	session, err := mgo.Dial("localhost")
	if err != nil {
		panic(err)
	}
	defer session.Close()
	session.SetMode(mgo.Monotonic, true)
	ensureIndex(session)
	route := mux.NewRouter()
	route.HandleFunc(("/resumes"), services.DisplayResume(session)).Methods("GET")
	route.HandleFunc(("/resumes"), services.AddResume(session)).Methods("POST")
	route.HandleFunc(("/resumes/{id}"), services.FindById(session)).Methods("GET")
	route.HandleFunc(("/resumes/{id}"), services.UpdateResume(session)).Methods("PUT")
	route.HandleFunc(("/resumes/{id}"), services.DeleteResume(session)).Methods("DELETE")
	http.ListenAndServe("localhost:8080", route)
}

func ensureIndex(s *mgo.Session) {
	session := s.Copy()
	defer session.Close()
	c := session.DB("Resume").C("ResumeManagement")
	index := mgo.Index{
		Key:        []string{"id"},
		Unique:     true,
		DropDups:   true,
		Background: true,
		Sparse:     true,
	}
	err := c.EnsureIndex(index)
	if err != nil {
		panic(err)
	}
}
