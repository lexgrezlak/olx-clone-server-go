package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"olx-clone-server/internal/service"
)

type Posting struct {
	title string
}

func AllPostings(datastore service.PostingDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ps, err := datastore.GetAllPostings()
		if err != nil {
			fmt.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		if payload, err := json.Marshal(ps); err == nil {
			w.Write(payload)
		}
	}
}
