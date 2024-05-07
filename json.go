package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func responseWithError(w http.ResponseWriter, code int, msg string) {

	if code > 499 {
		fmt.Println("Responding with 5XX err ", msg)
	}

	type errReponse struct {
		Error string `json:"error"`
	}

	responseWithJSON(w, code, errReponse{Error: msg})

}

func responseWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	// marshal payload to JSON object
	data, err := json.Marshal(payload)

	if err != nil {
		log.Printf("Failed to marshal json reponse %v ", payload)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(data)

}
