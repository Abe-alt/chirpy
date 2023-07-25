package main

import (
	"github.com/go-chi/chi/v5"
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

	files := http.FileServer(http.Dir(filepathRoot))
	cfg := cfgApi.middlewareMetricsInc(files)

	r := chi.NewRouter()
	r.Handle("/app", http.StripPrefix("/app", cfg))
	r.Handle("/app/*", http.StripPrefix("/app", cfg))

	//r.Get("/healthz", healthz)
	//r.Get("/metrics", cfgApi.metrics)

	apiRouter := chi.NewRouter()
	apiRouter.Get("/healthz", healthz)
	//apiRouter.Get("/metrics", cfgApi.metrics)
	apiRouter.Post("/validate_chirp", validate_chirp)
	r.Mount("/api", apiRouter)

	metricsRouter := chi.NewRouter()
	metricsRouter.Get("/metrics", cfgApi.metrics)
	r.Mount("/admin", metricsRouter)

	corsMux := middlewareCors(r)

	// ************ Cde w/ mux ***************
	//mux := http.NewServeMux()
	//
	//mux.Handle("/app/", http.StripPrefix("/app/", cfg))
	//mux.HandleFunc("/healthz", healthz)
	//mux.HandleFunc("/metrics", cfgApi.metrics)
	//
	//corsMux := middlewareCors(mux)
	// **************************************

	srv := &http.Server{
		Addr:    ":" + port,
		Handler: corsMux,
	}

	log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
	log.Fatal(srv.ListenAndServe())
}
