package main

import (
	"encoding/json"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
	"strings"
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

func respondwithJson(w http.ResponseWriter, code int, payload interface{}) {
	dat, err := json.Marshal(payload)
	if err != nil {
		log.Printf("Failed to marshal JSON response : %v", payload)
		w.WriteHeader(500)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(dat)

}

func respondwithError(w http.ResponseWriter, code int, msg string) {
	if code > 499 {
		log.Println("Responding with 5xx error:", msg)
	}
	type errResponse struct {
		Error string `json:"error"`
	}

	respondwithJson(w, code, errResponse{
		Error: msg,
	})

}

func validate_chirp(w http.ResponseWriter, r *http.Request) {

	type params struct {
		Body string `json:"body"`
	}
	type validChirp struct {
		Valid bool `json:"valid"`
	}

	decoder := json.NewDecoder(r.Body)
	param := params{}
	err := decoder.Decode(&param)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}
	text := strings.Split(param.Body, " ")

	for i := 0; i < len(text); i++ {
		if text[i] == "bad" {
			text[i] = "#####"
		}

	}

	if len(text) > 140 {
		respondwithError(w, 400, "Chirp is too long")
	} else {

		respondwithJson(w, 200, text)
	}

}
