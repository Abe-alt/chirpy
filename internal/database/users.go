package database

import "errors"

type User struct {
	ID       int    `json:"id"`
	Password int    `json:"password"`
	Email    string `json:"email"`
}

func (db *DB) CreateNewUser(email string, password int) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	id := len(dbStructure.Users) + 1
	user := User{
		ID:       id,
		Password: password,
		Email:    email,
	}
	dbStructure.Users[id] = user
	err = db.writeDB(dbStructure)
	if err != nil {
		return User{}, err
	}

	return user, nil

}

func (db *DB) GetUser(id int) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}

	user, ok := dbStructure.Users[id]
	if !ok {
		return User{}, errors.New("user does not exist")
	}

	return user, nil
}
