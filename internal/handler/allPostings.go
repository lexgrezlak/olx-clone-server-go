package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Posting struct {
	title string
}

func AllPostings(datastore PostingDatastore) http.HandlerFunc{
	return func(w http.ResponseWriter, r *http.Request) {
		// GET method returns all postings in the db
		// POST method creates a new posting
			ps, err := datastore.GetAllPostings(r.Context())
			if err != nil {
				fmt.Errorf("%v", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
			}
			w.WriteHeader(http.StatusOK)
			if payload, err := json.Marshal(ps); err == nil {
				w.Write(payload)
			}
			// Else http.MethodPost, it's specified in the router that
			// only GET and POST are allowed
	}
}