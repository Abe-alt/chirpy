package main

import (
	"fmt"
	"log"
	"net/http"
)

type apiconfig struct {
	fileserverHits int
}

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)
		next.ServeHTTP(w, r)
	})
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

	filesWithLog := middlewareLog(files)

	cfg := cfgApi.middlewareMetricsInc(filesWithLog)

	//mux.Handle("/app/", http.StripPrefix("/app/", files))
	mux.Handle("/app/", http.StripPrefix("/app/", cfg))
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/metrics", cfgApi.metrics)

	// Register the logsHandler function for the "/logs" path
	//mux.HandleFunc("/logs", logsHandler)

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
func logsHandler(w http.ResponseWriter, r *http.Request) {
	// Here, you can write the log content or any other response you want.
	// For example, you can read a log file and write its content to the response.
	logContent := "This is a log entry.\nMore log lines..."
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(logContent))
}
