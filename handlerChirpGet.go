package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"strconv"
)

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
