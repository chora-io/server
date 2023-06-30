package app

import (
	"net/http"

	"github.com/gorilla/mux"

	db "github.com/choraio/server/db/client"
)

func Auth(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	address := vars["address"]
	signature := vars["signature"]

	// TODO: create or update account and generate token
	token := address + "_" + signature

	respondJSON(rw, http.StatusOK, NewAuthResponse(token))
}
