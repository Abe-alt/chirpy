package database

import (
	"golang.org/x/crypto/bcrypt"
	"log"
)

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
