package handler

import (
	"errors"
	"fmt"
	"log"
	"net/http"
	"olx-clone-server/internal/service"
	"olx-clone-server/internal/util"
)

func CreatePosting(datastore PostingDatastore) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var i service.CreatePostingInput
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
		err = datastore.CreatePosting(r.Context(), i)
		if err != nil {
			fmt.Errorf("%v", err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		w.WriteHeader(http.StatusCreated)
	}
}
