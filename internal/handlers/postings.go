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
func Postings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// GET method returns all postings in the db
	// POST method creates a new posting
	switch r.Method {
	case http.MethodGet:
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
		break
		// Creates a posting
	case http.MethodPost:
		// Store the request body in i
		i := new(models.PostingInput)
		err := json.NewDecoder(r.Body).Decode(i)
		if err != nil {
			fmt.Errorf("%v", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		// Create the posting in the db
		err = models.CreatePosting(*i)
		if err != nil {
			fmt.Errorf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
		break
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		break
	}
}
