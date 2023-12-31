package main

import (
	"encoding/json"
	"net/http"
)

type Chirp struct {
	ID   int    `json:"id"`
	Body string `json:"body"`
}

func (cfg *apiconfig) handlerChirpsCreate(w http.ResponseWriter, r *http.Request) {
	type parameters struct {
		Body string `json:"body"`
	}

	decoder := json.NewDecoder(r.Body)
	params := parameters{}
	err := decoder.Decode(&params)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't decode parameters")
		return
	}

	chirp, err := cfg.DB.CreateChirp(params.Body)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't create chirp")
		return
	}

	respondwithJson(w, http.StatusCreated, Chirp{
		ID:   chirp.ID,
		Body: chirp.Body,
	})
}

//func validateChirp(body string) (string, error) {
//	const maxChirpLength = 140
//	if len(body) > maxChirpLength {
//		return "", errors.New("Chirp is too long")
//	}
//
//	badWords := map[string]struct{}{
//		"kerfuffle": {},
//		"sharbert":  {},
//		"fornax":    {},
//	}
//	cleaned := getCleanedBody(body, badWords)
//	return cleaned, nil
//}
//
//func getCleanedBody(body string, badWords map[string]struct{}) string {
//	words := strings.Split(body, " ")
//	for i, word := range words {
//		loweredWord := strings.ToLower(word)
//		if _, ok := badWords[loweredWord]; ok {
//			words[i] = "****"
//		}
//	}
//	cleaned := strings.Join(words, " ")
//	return cleaned
//}
