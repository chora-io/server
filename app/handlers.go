package app

import (
	"database/sql"
	"errors"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

	"github.com/choraio/server/db"
)

type GetRequestHandlerFunction func(dbr db.Reader, rw http.ResponseWriter, r *http.Request)

func GetContent(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 0, 64)
	if err != nil {
		respondError(
			rw,
			http.StatusBadRequest,
			fmt.Sprintf("invalid id: %s is not an integer", vars["id"]),
		)
		return
	}

	body, err := dbr.GetContent(r.Context(), id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(
				rw,
				http.StatusNotFound,
				fmt.Sprintf("content with id %d does not exist", id),
			)
		} else {
			respondError(rw, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondJSON(rw, http.StatusOK, NewGetContentResponse(id, body))
}

type PostRequestHandlerFunction func(dbw db.Writer, rw http.ResponseWriter, r *http.Request)

func PostContent(dbw db.Writer, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	body := vars["body"]

	id, err := dbw.PostContent(r.Context(), body)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostContentResponse(id, body))
}
