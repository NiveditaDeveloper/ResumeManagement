package services

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

//Find the candidate by Id
func FindById(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		vars := mux.Vars(r)
		id :=vars["id"]
		c := session.DB("Resume").C("ResumeManagement")
		var resDetails ResumeDetails
		err :=  c.Find(bson.M{"id":id}).One(&resDetails)
		if err != nil  {
			ErrorWithJSON(w, "Unable to get the candidate details :"+id, http.StatusInternalServerError)
			log.Println("Unable to get the candidate details : "+id, err)
			return
		}
		respBody, err := json.MarshalIndent(resDetails, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}
