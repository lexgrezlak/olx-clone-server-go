package service

import (
	"errors"
	"fmt"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

type SignUpInput struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"password"`
}

type User struct {
	Id           string `json:"id"`
	FirstName    string `json:"firstName"`
	LastName     string `json:"lastName"`
	Email        string `json:"email"`
	PasswordHash string `json:"passwordHash"`
}

func (api *api) CreateUser(i SignUpInput) error {
	ph, err := hashPassword(i.Password)
	if err != nil {
		return err
	}

	_, err = api.db.Exec(
		`INSERT INTO public.user ("firstName", "lastName", "email", "passwordHash") VALUES ($1, $2, $3, $4)`,
		i.FirstName, i.LastName, i.Email, ph)

	if err != nil {
		return err
	}
	return nil
}

func (api *api) ValidateUser(email, password string) (*User, error) {
	user, err := api.GetUserByEmail(email)
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
		fmt.Errorf("an error occurred while hashing password: %s", err)
		return "", err
	}
	return string(bytes), nil
}

func checkPasswordHash(password, passwordHash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(passwordHash), []byte(password))
	return err == nil
}
