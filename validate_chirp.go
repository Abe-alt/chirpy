package main

import (
	"net/http"
)

func validate_chirp(w http.ResponseWriter, r *http.Request) {
	type returnVals struct {
		Body string `json:"body"`
	}
	respBody := returnVals{
		Body: "Something went wrong",
	}

}
