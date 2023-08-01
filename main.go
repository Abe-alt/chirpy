package main

import (
	"chirpy/internal/database"
	"flag"
	"github.com/go-chi/chi/v5"
	"log"
	"net/http"
)

type apiconfig struct {
	fileserverHits int
	DB             *database.DB
}

func main() {
	const filepathRoot = "."
	const port = "8080"

	db, err := database.NewDB("database.json")
	if err != nil {
		log.Fatal(err)
	}
	dbg := flag.Bool("debug", false, "Enable debug mode")
	flag.Parse()
	if dbg != nil && *dbg {
		err := db.ResetDB()
		if err != nil {
			log.Fatal(err)
		}

		cfgApi := apiconfig{
			fileserverHits: 0,
			DB:             db,
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

		apiRouter.Post("/chirps", cfgApi.handlerChirpsCreate)
		apiRouter.Get("/chirps", cfgApi.handlerChirpsRetrieve)
		apiRouter.Get("/chirps/{id}", cfgApi.handlerChirpRetrieve)
		apiRouter.Post("/users", cfgApi.handlerUsers)
		apiRouter.Post("/login", cfgApi.handlerLogin)
		r.Mount("/api", apiRouter)

		metricsRouter := chi.NewRouter()
		metricsRouter.Get("/metrics", cfgApi.metrics)
		r.Mount("/admin", metricsRouter)

		corsMux := middlewareCors(r)

		srv := &http.Server{
			Addr:    ":" + port,
			Handler: corsMux,
		}

		log.Printf("Serving files from %s on port: %s\n", filepathRoot, port)
		log.Fatal(srv.ListenAndServe())
	}

	//func chirps(w http.ResponseWriter, r *http.Request) {
	//
	//	type params struct {
	//		Body string `json:"body"`
	//		//Cleaned_body string `json:"cleaned_Body"`
	//	}
	//	type clean struct {
	//		Cleaned_body string `json:"cleaned_Body"`
	//	}
	//	type validChirp struct {
	//		Valid bool `json:"valid"`
	//	}
	//
	//	decoder := json.NewDecoder(r.Body)
	//	param := params{}
	//	err := decoder.Decode(&param)
	//	if err != nil {
	//		respondwithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
	//		return
	//	}
	//
	//	if len(param.Body) > 140 {
	//		respondwithError(w, 400, "Chirp is too long")
	//	} else {
	//		respondwithJson(w, 200, clean{
	//			Cleaned_body: responsCleaned(param.Body),
	//		})
	//		//respondwithJson(w, 200, responsCleaned(param.Body))
	//
	//	}
	//
	//}
	//
	//func responsCleaned(text string) string {
	//
	//	slice := strings.Split(strings.ToLower(text), " ")
	//	for i := 0; i < len(slice); i++ {
	//		if slice[i] == "kerfuffle" || slice[i] == "sharbert" || slice[i] == "fornax" {
	//			slice[i] = "****"
	//		}
	//	}
	//	return strings.Join(slice, " ")
	//}
}
