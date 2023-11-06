package app

import (
	"encoding/json"
	"io"
	"net/http"
	"strings"

	"github.com/chora-io/server/auth"
	db "github.com/chora-io/server/db/client"
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

	// validate jwt token and return subject
	sub, err := auth.ValidateJWT(jsk, req.Token)
	if err != nil {
		respondError(rw, http.StatusUnauthorized, err.Error())
		return
	}

	// get user from database using subject (user id)
	user, err := dbr.SelectAuthUser(r.Context(), sub)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostAuthResponse(UserInfo{
		Id:        user.ID,
		Email:     user.Email.String,
		Address:   user.Address.String,
		Username:  user.Username.String,
		CreatedAt: user.CreatedAt,
	}))
}

func PostAuthEmail(jsk string, dbr db.Reader, dbw db.Writer, rw http.ResponseWriter, r *http.Request) {
	var req PostAuthEmailRequest

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

	if req.Email == "" {
		respondError(rw, http.StatusBadRequest, "email cannot be empty")
		return
	}

	// TODO: require access code
	//if req.AccessCode == "" {
	//	respondError(rw, http.StatusBadRequest, "access code cannot be empty")
	//	return
	//}

	// declare auth user
	var user db.AuthUser

	// check for existing auth user if token provided
	if req.Token != "" {

		// validate jwt token and return subject
		sub, err := auth.ValidateJWT(jsk, req.Token)
		if err != nil {
			respondError(rw, http.StatusUnauthorized, err.Error())
			return
		}

		// TODO: validate email access code

		// add or update email address for authenticated user
		err = dbw.UpdateAuthUserEmail(r.Context(), sub, req.Email)
		if err != nil {
			respondError(rw, http.StatusInternalServerError, err.Error())
			return
		}

		// get user from database using subject (user id)
		user, err = dbr.SelectAuthUser(r.Context(), sub)
		if err != nil {
			respondError(rw, http.StatusInternalServerError, err.Error())
			return
		}

	} else {

		// check for existing auth user and create if not found
		user, err = dbr.SelectAuthUserByEmail(r.Context(), req.Email)
		if err != nil {

			// insert auth user with email
			err = dbw.InsertAuthUserWithEmail(r.Context(), req.Email)
			if err != nil {
				respondError(rw, http.StatusInternalServerError, err.Error())
				return
			}

			// get new auth user by email
			user, err = dbr.SelectAuthUserByEmail(r.Context(), req.Email)
			if err != nil {
				respondError(rw, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	// generate new jwt token using user id
	token, err := auth.GenerateJWT(jsk, user.ID)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostAuthEmailResponse(token, UserInfo{
		Id:        user.ID,
		Email:     user.Email.String,
		Address:   user.Address.String,
		Username:  user.Username.String,
		CreatedAt: user.CreatedAt,
	}))
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

	if req.Address == "" {
		respondError(rw, http.StatusBadRequest, "address cannot be empty")
		return
	}

	if req.Signature == "" {
		respondError(rw, http.StatusBadRequest, "signature cannot be empty")
		return
	}

	// address must be a chora address
	if !strings.Contains(req.Address, "chora") {
		respondError(rw, http.StatusBadRequest, "address must be a chora address")
		return
	}

	// declare auth user
	var user db.AuthUser

	// check for existing auth user if token provided
	if req.Token != "" {

		// validate jwt token and return subject
		sub, err := auth.ValidateJWT(jsk, req.Token)
		if err != nil {
			respondError(rw, http.StatusUnauthorized, err.Error())
			return
		}

		// TODO: validate keplr account signature

		// add or update chora address for authenticated user
		err = dbw.UpdateAuthUserAddress(r.Context(), sub, req.Address)
		if err != nil {
			respondError(rw, http.StatusInternalServerError, err.Error())
			return
		}

		// get user from database using subject (user id)
		user, err = dbr.SelectAuthUser(r.Context(), sub)
		if err != nil {
			respondError(rw, http.StatusInternalServerError, err.Error())
			return
		}

	} else {

		// check for existing auth user if not found
		user, err = dbr.SelectAuthUserByAddress(r.Context(), req.Address)
		if err != nil {

			// insert auth user with address
			err = dbw.InsertAuthUserWithAddress(r.Context(), req.Address)
			if err != nil {
				respondError(rw, http.StatusInternalServerError, err.Error())
				return
			}

			// get new auth user by address
			user, err = dbr.SelectAuthUserByAddress(r.Context(), req.Address)
			if err != nil {
				respondError(rw, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	// generate new jwt token using user id
	token, err := auth.GenerateJWT(jsk, user.ID)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostAuthKeplrResponse(token, UserInfo{
		Id:        user.ID,
		Email:     user.Email.String,
		Address:   user.Address.String,
		Username:  user.Username.String,
		CreatedAt: user.CreatedAt,
	}))
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

	if req.Username == "" {
		respondError(rw, http.StatusBadRequest, "username cannot be empty")
		return
	}

	// TODO: require password
	//if req.Password == "" {
	//	respondError(rw, http.StatusBadRequest, "password cannot be empty")
	//	return
	//}

	// declare auth user
	var user db.AuthUser

	// TODO: hash password
	// hashedPassword := req.Password

	// check for existing auth user if token provided
	if req.Token != "" {

		// validate jwt token and return subject
		sub, err := auth.ValidateJWT(jsk, req.Token)
		if err != nil {
			respondError(rw, http.StatusUnauthorized, err.Error())
			return
		}

		// TODO: validate hashed password

		// add or update username for authenticated user
		err = dbw.UpdateAuthUserUsername(r.Context(), sub, req.Username)
		if err != nil {
			respondError(rw, http.StatusInternalServerError, err.Error())
			return
		}

		// get user from database using subject (user id)
		user, err = dbr.SelectAuthUser(r.Context(), sub)
		if err != nil {
			respondError(rw, http.StatusInternalServerError, err.Error())
			return
		}

	} else {

		// check for existing auth user and create if not found
		user, err = dbr.SelectAuthUserByUsername(r.Context(), req.Username)
		if err != nil {

			// insert auth user with username
			err = dbw.InsertAuthUserWithUsername(r.Context(), req.Username)
			if err != nil {
				respondError(rw, http.StatusInternalServerError, err.Error())
				return
			}

			// get new auth user by username
			user, err = dbr.SelectAuthUserByUsername(r.Context(), req.Username)
			if err != nil {
				respondError(rw, http.StatusInternalServerError, err.Error())
				return
			}
		}
	}

	// generate new jwt token using user id
	token, err := auth.GenerateJWT(jsk, user.ID)
	if err != nil {
		respondError(rw, http.StatusInternalServerError, err.Error())
		return
	}

	respondJSON(rw, http.StatusOK, NewPostAuthLoginResponse(token, UserInfo{
		Id:        user.ID,
		Email:     user.Email.String,
		Address:   user.Address.String,
		Username:  user.Username.String,
		CreatedAt: user.CreatedAt,
	}))
}
