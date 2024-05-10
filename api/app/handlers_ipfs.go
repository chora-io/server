package app

import (
	"encoding/json"
	"io"
	"net/http"

	db "github.com/chora-io/server/db/client"
	"github.com/gorilla/mux"
)

func GetIpfs(dbr db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid := vars["cid"]

	// TODO: get content from local ipfs node
	var content = ""

	respondJSON(rw, http.StatusOK, NewGetIpfsResponse(cid, content))
}

func PostIpfs(dbw db.Writer, rw http.ResponseWriter, r *http.Request) {
	var req PostIpfsRequest

	bz, err := io.ReadAll(r.Body)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(bz, &req)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	content := compactJSONString(req.Content)

	// TODO: pin content to local ipfs node
	var cid = ""

	respondJSON(rw, http.StatusOK, NewPostIpfsResponse(cid, content))
}
