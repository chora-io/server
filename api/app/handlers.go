package app

import (
	"net/http"

	db "github.com/choraio/server/db/client"
)

type AuthHandler func(jsk string, dbr db.Reader, dbw db.Writer, rw http.ResponseWriter, r *http.Request)

type GetHandler func(dbr db.Reader, rw http.ResponseWriter, r *http.Request)

type PostHandler func(dbw db.Writer, rw http.ResponseWriter, r *http.Request)
