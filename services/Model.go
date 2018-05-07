package services

import (
	"net/http"
	"log"
)

type ResumeDetails struct {
	Id            		string `json:"id"`
	//candidate details will be moved to candidate collection  and address details will be moved to address collection(in process)
	CandidateFirstName 	string `json:"candidate_first_name"`
	CandidateMiddleName 	string `json:"candidate_middle_name"`
	CandidateLastName 	string `json:"candidate_last_name"`
	ContactNumber 		string `json:"contact_number"`
	EmailId 		string `json:"email_id"`
	Expertise 		string `json:"candidate_expertise"`
	CurrentLocation 	string `json:"current_location"`
	PrefferedLocation	string `json:"preffered_location"`
	CurrentCompany 		string `json:"current_company"`
	CurrentDesignation 	string `json:"current_designation"`
	NoticePeriod 		string `json:"notice_period"`
	Experience 		string `json:"experience"`
	CurrentCTC 		float64 `json:"current_ctc"`
	ExpectedCTC 		float64 `json:"expected_ctc"`
	ResFile 		string `json:"url"`
	InterviewStatus 	string `json:"interview_status"`
	ModeOfInterview 	string `json:"mode_of_interview"`
}

func ErrorWithJSON(w http.ResponseWriter, message string, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
   	w.Write([]byte(message))
	log.Println(w, message)
}

func ResponseWithJSON(w http.ResponseWriter, json []byte, code int) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(code)
	w.Write(json)
}
