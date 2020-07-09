package handler

import (
	"errors"
	"log"
	"net/http"
	"olx-clone-server/internal/common"
	"olx-clone-server/internal/service"
	"olx-clone-server/internal/util"
)



func SignUp(w http.ResponseWriter, r *http.Request) {
	var i common.SignUpInput

	// DecodeJSONBody handles the other errors.
	err := util.DecodeJSONBody(w, r, &i)
	if err != nil {
		var mr *util.MalformedRequest
		if errors.As(err, &mr) {
			http.Error(w, mr.Msg, mr.Status)
		} else {
			log.Print(err)
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}

	err = service.CreateUser(i)
	if err != nil {
		msg := "User with this email already exists"
		http.Error(w, msg, http.StatusBadRequest)
		return
	}

	handleTokenResponse(w, i.Email)
	w.WriteHeader(http.StatusCreated)
}