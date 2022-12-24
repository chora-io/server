package app

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	"github.com/gorilla/mux"

	"github.com/choraio/server/db"
)

type GetRequestHandlerFunction func(dbr db.Reader, rw http.ResponseWriter, r *http.Request)

func GetData(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	id, err := strconv.ParseInt(vars["id"], 0, 32)
	if err != nil {
		respondError(
			rw,
			http.StatusBadRequest,
			fmt.Sprintf("invalid id: %s is not an integer", vars["id"]),
		)
		return
	}

	d, err := dbr.GetData(r.Context(), int32(id))
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			respondError(
				rw,
				http.StatusNotFound,
				fmt.Sprintf("data with id %d does not exist", id),
			)
		} else {
			respondError(rw, http.StatusInternalServerError, err.Error())
		}
		return
	}

	respondJSON(rw, http.StatusOK, NewGetDataResponse(int32(id), d.Canon, d.Context, d.Jsonld))
}

type PostRequestHandlerFunction func(dbw db.Writer, rw http.ResponseWriter, r *http.Request)

func PostData(dbw db.Writer, rw http.ResponseWriter, r *http.Request) {
	var d db.Datum

	bz, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(bz, &d)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	canon := strings.TrimSpace(d.Canon)
	context := strings.TrimSpace(d.Context)
	jsonld := compactJSONString(d.Jsonld)

	id, err := dbw.PostData(r.Context(), canon, context, jsonld)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostDataResponse(id, canon, context, jsonld))
}
