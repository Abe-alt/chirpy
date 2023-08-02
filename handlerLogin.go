package main

import (
	"encoding/json"
	"net/http"
	"time"
)

func (cfg *apiconfig) handlerLogin(w http.ResponseWriter, r *http.Request) {
	type parameter struct {
		Password         string `json:"password"`
		Email            string `json:"email"`
		ExpiresInSeconds int    `json:"expires_in_seconds"`
	}
	type response struct {
		User
		Token string `json:"token"`
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

	err = cfg.DB.CheckPasswordHash(user.HashedPassword, parameters.Password)
	if err != nil {
		respondwithError(w, http.StatusUnauthorized, "incorrect password")
		return
	}

	defaultExpiration := 60 * 60 * 24
	if parameters.ExpiresInSeconds == 0 {
		parameters.ExpiresInSeconds = defaultExpiration
	} else if parameters.ExpiresInSeconds > defaultExpiration {
		parameters.ExpiresInSeconds = defaultExpiration
	}

	token, err := cfg.DB.MakeJWT(user.ID, cfg.jwtSecret, time.Duration(parameters.ExpiresInSeconds)*time.Second)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't create JWT")
		return
	}

	respondwithJson(w, 200, response{
		User{
			ID:    user.ID,
			Email: user.Email,
		},
		token,
	})

}
