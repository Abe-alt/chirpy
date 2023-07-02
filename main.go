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
	const port = "8080"
	const filepathRoot = "."
	// Create a new http.ServeMux
	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))
	//Wrap that mux in a custom middleware function that adds CORS headers to the response (see the tip below on how to do that).
	corsMux := middlewareCors(mux)

	//Create a new http.Server and use the corsMux as the handler
	srv := &http.Server{Addr: ":" + port, Handler: corsMux}
	//&http.Server{} creates a new http.Server struct and returns a pointer to that struct.

	//Use the server's ListenAndServe method to start the server
	log.Printf("serving files from %s on port %s:", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
	//if err := http.ListenAndServe(":8080", corsMux); err != nil {
	//	log.Fatal(err)
	//}

}