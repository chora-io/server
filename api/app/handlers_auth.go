package app

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/choraio/server/auth"
	db "github.com/choraio/server/db/client"
)

func PostAuth(jsk string, dbr db.Reader, dbw db.Writer, rw http.ResponseWriter, r *http.Request) {
	var req PostAuthRequest

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

	// TODO: verify address is associated with token

	// verify auth user exists
	_, err = dbr.SelectAuthUser(r.Context(), req.Address)
	if err != nil {
		respondError(rw, http.StatusNotFound, err.Error())
		return
	}

	// validate jwt token
	err = auth.ValidateJWT(jsk, req.Token)
	if err != nil {
		respondError(rw, http.StatusUnauthorized, err.Error())
		return
	}

	// update last authenticated
	err = dbw.UpdateAuthUserLastAuthenticated(r.Context(), req.Address)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostAuthResponse(true))
}

func PostAuthKeplr(jsk string, dbr db.Reader, dbw db.Writer, rw http.ResponseWriter, r *http.Request) {
	var req PostAuthKeplrRequest

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

	// address must be a chora address
	if !strings.Contains(req.Address, "chora") {
		respondError(rw, http.StatusBadRequest, "invalid address")
		return
	}

	// generate jwt token using address
	token, err := auth.GenerateJWT(jsk, req.Address, req.Signature)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	// check if auth user exists
	_, err = dbr.SelectAuthUser(r.Context(), req.Address)
	if err != nil {

		// insert auth user
		err = dbw.InsertAuthUser(r.Context(), req.Address)
		if err != nil {
			respondError(rw, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// update last authenticated
	err = dbw.UpdateAuthUserLastAuthenticated(r.Context(), req.Address)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostAuthKeplrResponse(token, true))
}

func PostAuthLogin(jsk string, dbr db.Reader, dbw db.Writer, rw http.ResponseWriter, r *http.Request) {
	var req PostAuthLoginRequest

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

	// TODO: create or access custodial account
	address := "[custodial account address]"
	signature := "[custodial account signature]"

	// generate jwt token using address
	token, err := auth.GenerateJWT(jsk, address, signature)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	// check if auth user exists
	_, err = dbr.SelectAuthUser(r.Context(), address)
	if err != nil {

		// insert auth user
		err = dbw.InsertAuthUser(r.Context(), address)
		if err != nil {
			respondError(rw, http.StatusInternalServerError, err.Error())
			return
		}
	}

	// update last authenticated
	err = dbw.UpdateAuthUserLastAuthenticated(r.Context(), address)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostAuthLoginResponse(token, true))
}
