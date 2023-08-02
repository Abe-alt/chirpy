package database

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"time"
)

func (db *DB) HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", nil
	}
	return string(hash), nil
}

func (db *DB) CheckPasswordHash(hash string, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}

// MakeJWT -
func (db *DB) MakeJWT(userID int, tokenSecret string, expiresIn time.Duration) (string, error) {
	signingKey := []byte(tokenSecret)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Issuer:    "chirpy",
		IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
		ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(expiresIn)),
		Subject:   fmt.Sprintf("%d", userID),
	})
	return token.SignedString(signingKey)
}

//func (db *DB) ValidateJWT(tokenString, tokenSecret string) (string, error) {
//	//claims := jwt.RegisteredClaims{
//	//	Issuer:    "chirpy",
//	//	Subject:   string(),
//	//	Audience:  nil,
//	//	ExpiresAt: jwt.NewNumericDate(time.Now().UTC().Add(2)),
//	//	IssuedAt:  jwt.NewNumericDate(time.Now().UTC()),
//	//}
//	//
//	//jwt.NewWithClaims(jwt.SigningMethodES256, claims)
//
//	claimsStruct := jwt.RegisteredClaims{}
//	token, err := jwt.ParseWithClaims(
//		tokenString,
//		&claimsStruct,
//		func(token *jwt.Token) (interface{}, error) { return []byte(tokenSecret), nil },
//	)
//	if err != nil {
//		return "", err
//	}
//
//	userIDString, err := token.Claims.GetSubject()
//	if err != nil {
//		return "", err
//	}
//
//	expiresAt, err := token.Claims.GetExpirationTime()
//	if err != nil {
//		return "", err
//	}
//
//	if expiresAt.Before(time.Now().UTC()) {
//		return "", errors.New("JWT is expired")
//	}
//
//	return userIDString, nil
//}
