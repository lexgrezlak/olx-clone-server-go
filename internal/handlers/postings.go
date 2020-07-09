package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"olx-clone-server/internal/models"
)

type Posting struct {
	title string
}

// PostingHandler responds with Postings in plaintext
func GetAllPostings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	ps, err := models.AllPostings()
	if err != nil {
		fmt.Errorf("%v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	if payload, err := json.Marshal(ps); err == nil {
		w.Write(payload)
	}
}