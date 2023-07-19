package main

import (
	"fmt"
	"log"
	"net/http"
)

type apiconfig struct {
	fileserverHits int
}

func (cfg *apiconfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(writer, request)
	})
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
	//mux.Handle("/app/", http.StripPrefix("/app/", files))
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

func healthz(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)                    // 200
	w.Write([]byte(http.StatusText(http.StatusOK))) // ("ok")
}

func (cfg *apiconfig) metrics(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf("Hits:%v", cfg.fileserverHits)))
}
