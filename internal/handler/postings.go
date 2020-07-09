package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"olx-clone-server/internal/service"
	"olx-clone-server/internal/util"
)

type Posting struct {
	title string
}

// PostingHandler responds with Postings in plaintext
func Postings(w http.ResponseWriter, r *http.Request) {
	// GET method returns all postings in the db
	// POST method creates a new posting
	if r.Method == http.MethodGet {
		ps, err := service.AllPostings()
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
	} else {
		var i service.PostingInput
		err := util.DecodeJSONBody(w, r, &i)
		if err != nil {
			var mr *util.MalformedRequest
			if errors.As(err, &mr) {
				http.Error(w, mr.Msg, mr.Status)
			} else {
				// Default to 500 Internal Server Error
				log.Print(err.Error())
				http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
			}
			return
		}

		//Create the posting in the db with the given input
		err = service.CreatePosting(i)
		if err != nil {
			fmt.Errorf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
