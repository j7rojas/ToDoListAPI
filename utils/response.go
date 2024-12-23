package utils

import (
	"encoding/json"
	"log"
	"net/http"
)

// JSONResponse function that handlers the conversation of a struct to a json and the return code
// two things that will be done for every handler function
//
// Parameters
// status - Status code you wish to return
// data - The data you are trying to encode into json
// w - requires Response Write to send back response
//
// This function has no return
func JSONResponse(w http.ResponseWriter, status int, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(data); err != nil {
		log.Fatalf("Failed to encode JSON response: %v", err)
	}
}
