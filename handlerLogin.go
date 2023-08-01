package main

import (
	"encoding/json"
	"net/http"
)

func (cfg *apiconfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	parameters := parameter{}
	err := decoder.Decode(&parameters)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "couldn't decode parameters")
		return
	}
	user, err := cfg.DB.GetUserByEmail(parameters.Email)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "couldn't get user")
		return
	}

	err = cfg.DB.CheckPasswordHash(user.Password, parameters.Password)
	if err != nil {
		respondwithError(w, http.StatusUnauthorized, "incorrect password")
		return
	}

	respondwithJson(w, 200, User{
		ID:    user.ID,
		Email: user.Email,
	})

}
