package handler

import (
	"encoding/json"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"olx-clone-server/internal/service"
	"olx-clone-server/internal/util"
	"time"
)

// Create the JWT key used to create the signature
var jwtKey := []byte("secret_key")

// Create a struct to read the email and the password
// from the request body
type SignInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json: "email"`
	jwt.StandardClaims
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var i SignInInput
	err := util.DecodeJSONBody(w, r, &i)

	// DecodeJSONBody handles all the errors.
	if err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
	}

	// Verify if the credentials are correct.
	user, err := service.ValidateUser(i.Email, i.Password)
	if err != nil {
		msg := "Wrong credentials"
		http.Error(w, msg, http.StatusUnauthorized)
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which include the email and expiry time.
	claims := &Claims{
		Email:          i.Email,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed in unix milliseconds.
			ExpiresAt: expirationTime.Unix(),
		},
	}

	// Sign the token.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string.
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
	}

	http.SetCookie(w, &http.Cookie{
		Name:       "token",
		Value:      tokenString,
		Expires:    expirationTime,
		Secure:     false,
		HttpOnly:   true,
		SameSite:   0,
	})
}