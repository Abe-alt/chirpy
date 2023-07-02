package main

import (
	"log"
	"net/http"
)

func middlewareCors(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS, PUT, DELETE")
		w.Header().Set("Access-Control-Allow-Headers", "*")
		if r.Method == "OPTIONS" {
			w.WriteHeader(http.StatusOK)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func main() {
	// Create a new http.ServeMux
	mux := http.NewServeMux()

	//Wrap that mux in a custom middleware function that adds CORS headers to the response (see the tip below on how to do that).
	corsMux := middlewareCors(mux)

	//Create a new http.Server and use the corsMux as the handler
	srv := http.Server{Addr: ":8080", Handler: corsMux}

	//Use the server's ListenAndServe method to start the server
	log.Printf("serving on port 8080:")
	log.Fatal(srv.ListenAndServe())
	//if err := http.ListenAndServe(":8080", corsMux); err != nil {
	//	log.Fatal(err)
	//
	//}
}
