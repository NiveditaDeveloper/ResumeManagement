package services

import (
	"gopkg.in/mgo.v2"
	"net/http"
	"gopkg.in/mgo.v2/bson"
	"log"
	"encoding/json"
	"github.com/gorilla/mux"
)

//Display all the values; values will be displayed in an UI grid structure, paginated as per the number of records (in process)
func DisplayResume(s *mgo.Session) func(w http.ResponseWriter, r *http.Request){
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		c := session.DB("Resume").C("ResumeManagement")
		var resDetails []ResumeDetails
		err := c.Find(bson.M{}).All(&resDetails)
		if err!=nil{
			ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
			log.Println("Failed to fetch resumes: ", err)
			return
		}
		respBody, err := json.MarshalIndent(resDetails, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		ResponseWithJSON(w, respBody, http.StatusOK)
	}
}

//Add a new candidate details
func AddResume (s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		var resDetails ResumeDetails
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&resDetails)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}
		c := session.DB("Resume").C("ResumeManagement")
		err = c.Insert(resDetails)
		if err!= nil{
			ErrorWithJSON(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Header().Set("Location", r.URL.Path+"/"+resDetails.Id)
		w.WriteHeader(http.StatusCreated)
	}
}

//Update an existing candidate details
func UpdateResume(s *mgo.Session) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		vars := mux.Vars(r)
		id :=vars["id"]
		var resDetails ResumeDetails
		decoder := json.NewDecoder(r.Body)
		err := decoder.Decode(&resDetails)
		if err != nil {
			ErrorWithJSON(w, "Incorrect body", http.StatusBadRequest)
			return
		}
		c := session.DB("Resume").C("ResumeManagement")
		err = c.Update(bson.M{"id": id}, &resDetails)
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed to update Candidate Details: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Candidate not found: "+id, http.StatusNotFound)
				return
			}
		}
		log.Println("Details Updated")
		w.WriteHeader(http.StatusOK)
	}
}

//Delete the candidate details by id
func DeleteResume(s *mgo.Session)  func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		session := s.Copy()
		defer session.Close()
		vars := mux.Vars(r)
		id :=vars["id"]
		c := session.DB("Resume").C("ResumeManagement")
		err := c.Remove(bson.M{"id": id})
		if err != nil {
			switch err {
			default:
				ErrorWithJSON(w, "Database error", http.StatusInternalServerError)
				log.Println("Failed to delete Candidate Details: ", err)
				return
			case mgo.ErrNotFound:
				ErrorWithJSON(w, "Candidate Details not found : "+id, http.StatusNotFound)
				return
			}
		}
		log.Println("Details Deleted")
		w.WriteHeader(http.StatusOK)
	}
}
