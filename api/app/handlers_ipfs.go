package app

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/gorilla/mux"
	shell "github.com/ipfs/go-ipfs-api"

	db "github.com/chora-io/server/db/client"
)

func GetIpfs(_ db.Reader, rw http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	cid := vars["cid"]

	sh := shell.NewShell("http://127.0.0.1:5001")

	rdr, err := sh.Cat(cid)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	buf := new(bytes.Buffer)
	_, err = buf.ReadFrom(rdr)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}
	content := buf.String()

	respondJSON(rw, http.StatusOK, NewGetIpfsResponse(cid, content))
}

func PostIpfs(_ db.Writer, rw http.ResponseWriter, r *http.Request) {
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

	sh := shell.NewShell("http://127.0.0.1:5001")

	content := compactJSONString(req.Content)

	cid, err := sh.Add(strings.NewReader(content))
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostIpfsResponse(cid, req.Content))
}
