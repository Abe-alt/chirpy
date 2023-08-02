package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	ID int `json:"id"`
	//Password string `json:"password"`
	Email string `json:"email"`
}

func (cfg *apiconfig) handlerUsers(w http.ResponseWriter, r *http.Request) {

	type parameters struct {
		Password string `json:"password"`
		Email    string `json:"email"`
	}
	type response struct {
		User
	}
	decoder := json.NewDecoder(r.Body)
	parameter := parameters{}
	err := decoder.Decode(&parameter)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "can't decode parameters")
		return
	}

	hashedPassword, err := cfg.DB.HashPassword(parameter.Password)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "Couldn't hash password")
		return
	}

	user, err := cfg.DB.CreateNewUser(parameter.Email, hashedPassword)
	if err != nil {
		respondwithError(w, http.StatusInternalServerError, "can't create user")
		return
	}
	respondwithJson(w, 201, response{
		User: User{
			ID: user.ID,
			//Password: user.Password,
			Email: user.Email,
		},
	})

}
