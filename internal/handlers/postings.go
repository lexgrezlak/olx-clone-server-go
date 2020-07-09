package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"olx-clone-server/internal/services"
	"olx-clone-server/internal/utils"
)

type Posting struct {
	title string
}

// PostingHandler responds with Postings in plaintext
func Postings(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	// GET method returns all postings in the db
	// POST method creates a new posting
	switch r.Method {
	case http.MethodGet:
		ps, err := services.AllPostings()
		if err != nil {
			fmt.Errorf("%v", err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
		w.WriteHeader(http.StatusOK)
		if payload, err := json.Marshal(ps); err == nil {
			w.Write(payload)
		}
		break
	case http.MethodPost:
		var i services.PostingInput
		err := utils.DecodeJSONBody(w, r, &i)
		if err != nil {
			var mr *utils.MalformedRequest
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
		err = services.CreatePosting(i)
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
