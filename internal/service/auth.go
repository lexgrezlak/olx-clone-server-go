package service

import (
	"context"
	"errors"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type User struct {
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Email string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

func getUserByEmail(email string) (*User, error) {
	rows, err := db.Pool.Query(context.Background(),
		"SELECT * FROM user WHERE email=$1 LIMIT 1", email)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var u User
	err = rows.Scan(&u.FirstName, &u.LastName, &u.Email, &u.PasswordHash)
	if err != nil {
		return nil, err
	}

	err = rows.Err()
	if err != nil {
		return nil, err
	}

	return &u, err
}

func ValidateUser(email, password string) (*User, error) {
	user, err := getUserByEmail(email)
	if err != nil {
		return nil, err
	}
	isPasswordValid := checkPasswordHash(user.PasswordHash, password)
	if !isPasswordValid {
		return nil, errors.New(http.StatusText(http.StatusUnauthorized))
	}

	return user, nil
}


func hashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	if err != nil {
		return "", err
	}
	return string(bytes), nil
}

func checkPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}