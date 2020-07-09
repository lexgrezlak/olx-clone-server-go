package handler

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"log"
	"net/http"
	"olx-clone-server/internal/service"
	"olx-clone-server/internal/util"
)

// Create the JWT key used to create the signature
var jwtKey = []byte("secret_key")

// Create a struct to read the email and the password
// from the request body
type signInInput struct {
	Email string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Email string `json:"email"`
	jwt.StandardClaims
}

func SignIn(w http.ResponseWriter, r *http.Request) {
	var i signInInput
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
		return
	}

	// Verify if the credentials are correct.
	u, err := service.ValidateUser(i.Email, i.Password)
	if err != nil {
		msg := "Wrong credentials"
		http.Error(w, msg, http.StatusUnauthorized)
		return
	}

	handleTokenResponse(w, u.Email)
}

