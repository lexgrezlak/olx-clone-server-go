package handler

import (
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

func handleTokenResponse(w http.ResponseWriter, email string) {
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which include the email and expiry time.
	claims := &Claims{
		Email:          email,
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
