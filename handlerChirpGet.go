package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

func (cfg *apiconfig) handlerChirpRetrieve(w http.ResponseWriter, r *http.Request) {

	uID := chi.URLParam(r, "id")
	dbChirps, err := cfg.DB.GetChirps()
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't retrieve chirps")
	}
	chirps := []Chirp{}
	for _, dbChirp := range dbChirps {
		chirps = append(chirps, Chirp{
			ID:   dbChirp.ID,
			Body: dbChirp.Body,
		})
	}
	intUID, _ := strconv.Atoi(uID)
	for i := 0; i < len(chirps); i++ {
		//for _, chirp := range chirps {
		if chirps[i].ID == intUID {
			respondwithJson(w, http.StatusOK, chirps[i].Body)
		}
	}
}
