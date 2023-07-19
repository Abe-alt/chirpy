package main

import (
	"log"
	"net/http"
)

func logsHandler(w http.ResponseWriter, r *http.Request) {
	// Here, you can write the log content or any other response you want.
	// For example, you can read a log file and write its content to the response.
	logContent := "This is a log entry.\nMore log lines..."
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(logContent))
}

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
}
