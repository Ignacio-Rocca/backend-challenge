package handler

import (
	"encoding/json"
	"log"
	"net/http"
)

func setResponse(w http.ResponseWriter, statusCode int, body interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if body != nil {
		err := json.NewEncoder(w).Encode(body)
		if err != nil {
			log.Println(err.Error())
		}
	}
}
