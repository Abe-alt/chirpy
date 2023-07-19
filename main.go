package main

import (
	"log"
	"net/http"
)

type apiconfig struct {
	fileserverHits int
}

func main() {
	const filepathRoot = "."
	const port = "8080"
	cfgApi := apiconfig{
		fileserverHits: 0,
	}

	mux := http.NewServeMux()
	files := http.FileServer(http.Dir(filepathRoot))
	cfg := cfgApi.middlewareMetricsInc(files)

	mux.Handle("/app/", http.StripPrefix("/app/", cfg))
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/metrics", cfgApi.metrics)

	corsMux := middlewareCors(mux)

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
