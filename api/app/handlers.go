package app

import (
	"net/http"

	db "github.com/choraio/server/db/client"
)

type GetHandler func(dbr db.Reader, rw http.ResponseWriter, r *http.Request)
