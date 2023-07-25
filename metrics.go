package main

import (
	"fmt"
	"net/http"
)

func (cfg *apiconfig) middlewareMetricsInc(next http.Handler) http.Handler {
	return http.HandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
		cfg.fileserverHits++
		next.ServeHTTP(writer, request)
	})
}

func (cfg *apiconfig) metrics(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add("Content-Type", "text/plain; charset=utf-8")
	w.Header().Add("Content-Type", "text/html")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(fmt.Sprintf(
		`<html>

	<body>
	<h1>Welcome, Chirpy Admin</h1>
	<p>Chirpy has been visited %d times!</p>
	</body>

	</html>
`, cfg.fileserverHits)))

	//w.Write([]byte(fmt.Sprintf("Hits:%v", cfg.fileserverHits)))

}
