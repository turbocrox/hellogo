package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, message string) {
	if code > 499 {
		log.Println("Server error:", message)
	}
	type errResponse struct {
		Error string `json:"error"` /// convert  into  a  error json//
	}
	respondWithJSON(w, code, errResponse{Error: message})
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Println("Error marshalling JSON:", err)
		w.WriteHeader(500)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)
}

////  this file  is  made  to  covert  to  json  file and  alll it help  to    convert  and  write  in  it  ///
