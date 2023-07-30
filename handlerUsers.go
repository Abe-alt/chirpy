package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID    int    `json:"id"`
	Email string `json:"email"`
}

func (cfg *apiconfig) handlerUsers(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Email string `json:"email"`
	}
	decoder := json.NewDecoder(r.Body)
	parameter := parameters{}
	err := decoder.Decode(&parameter)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "can't decode parameters")
		return
	}
	user, err := cfg.DB.CreateNewUser(parameter.Email)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "can't create user")
		return
	}
	respondwithJson(w, 201, User{
		ID:    user.ID,
		Email: user.Email,
	})

}
