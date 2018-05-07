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
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed to find Candidate: ", err)
			return
		}
		if resDetails.Id == "" {
			ErrorWithJSON(w, "Id not found", http.StatusNotFound)
			return
		}
		respBody, err := json.MarshalIndent(resDetails, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}
