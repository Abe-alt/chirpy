package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"sort"
	"strconv"
)

func (cfg *apiconfig) handlerChirpsRetrieve(w http.ResponseWriter, r *http.Request) {
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
		return
	}

	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:   dbChirp.ID,
			Body: dbChirp.Body,
		})
	}

	sort.Slice(chirps, func(i, j int) bool {
		return chirps[i].ID < chirps[j].ID
	})

	respondwithJson(w, http.StatusOK, chirps)
}

func (cfg *apiconfig) handlerChirpRetrieve(w http.ResponseWriter, r *http.Request) {
	uID := chi.URLParam(r, "id")
	intUID, err := strconv.Atoi(uID)
	if err != nil {
		respondwithError(w, http.StatusNotFound, "Wrong Chirp ID")
		return
	}

	dbChirp, err := cfg.DB.GetChirp(intUID)
	if err != nil {
		respondwithError(w, 404, "Couldn't retrieve the chirp")
		return
	}

	respondwithJson(w, http.StatusOK, Chirp{
		ID:   dbChirp.ID,
		Body: dbChirp.Body,
	})

	// ***** own code to retrieve a clean chirp msg *******
	//	for i := 0; i < len(chirps); i++ {
	//		//for _, chirp := range chirps {
	//		if chirps[i].ID == intUID {
	//			respondwithJson(w, http.StatusOK, chirps[i].Body)
	//		}
	//	}
	//	w.WriteHeader(404)
}
