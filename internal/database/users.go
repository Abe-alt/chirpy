package database

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"log"
)

type User struct {
	ID       int    `json:"id"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

func (db *DB) CreateNewUser(email, password string) (User, error) {

	if _, err := db.GetUserByEmail(email); !errors.Is(err, errors.New("user not exists")) {
		return User{}, errors.New("user already exists")
	}
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	id := len(dbStructure.Users) + 1
	user := User{
		ID:       id,
		Password: db.HashPassword(password),
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

func (db *DB) HashPassword(password string) string {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		log.Printf("couldn't encrypt the password")
	}
	return string(hash)
}

func (db *DB) CheckPasswordHash(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

func (db *DB) GetUserByEmail(email string) (User, error) {
	dbStructure, err := db.loadDB()
	if err != nil {
		return User{}, err
	}
	for _, user := range dbStructure.Users {
		if user.Email == email {
			return user, nil
		}
	}
	return User{}, errors.New("user not exists")
}
