package main

import (
	"log"
	"net/http"
)

func main() {
	const filepathRoot = "."
	//const healthzPath = "/healthz"
	const imageFilePath = "./assets"
	const port = "8080"

	mux := http.NewServeMux()
	mux.Handle("/", http.FileServer(http.Dir(filepathRoot)))

	mux.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		w.Write([]byte("ok"))
	})
	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", imageFilePath, port)
	log.Fatal(srv.ListenAndServe())
}
