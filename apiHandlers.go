package main

import (
	"net/http"
	"fmt"
	"encoding/json"
)

const responseFailureStatus = "404 not found"
const responseFailureMessage = "0 result found."

type response struct {
	Status string `json:"status"`
	Message string `json:"message"`
	Payload interface{} `json:"payload"`
}

func handle404(w http.ResponseWriter, r *http.Request) {

	jsonBytes, err := json.Marshal(response{Status:responseFailureStatus,
		Message: responseFailureMessage,
		Payload: make([]string, 0),
	})

	if err != nil {
		logError(err, "handle404, json.Marshal", "")
	}

	fmt.Fprintf(w, string(jsonBytes))
}




func homeHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "home page...")
}