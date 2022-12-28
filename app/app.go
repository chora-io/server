package app

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	"github.com/choraio/server/db"
)

type App struct {
	env string
	dbr db.Reader
	dbw db.Writer
	rtr *mux.Router
	aos string
	log zerolog.Logger
}

// Initialize initializes the app.
func Initialize(cfg Config, dbr db.Reader, dbw db.Writer, log zerolog.Logger) *App {
	app := &App{
		env: cfg.AppEnv,
		dbr: dbr,
		dbw: dbw,
		rtr: mux.NewRouter(),
		aos: cfg.AppAllowedOrigins,
		log: log,
	}

	// get requests
	app.Get("/data/{iri}", app.handleGetRequest(GetData))

	// post requests
	app.Post("/data", app.handlePostRequest(PostData))

	return app
}

// Run blocks the current thread of execution and serves the API.
func (a *App) Run(host string) {
	aos := handlers.AllowedOrigins(strings.Split(a.aos, ","))
	log.Fatal(http.ListenAndServe(host, handlers.CORS(aos)(a.rtr)))
}

// Get wraps the router for GET method
func (a *App) Get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.rtr.HandleFunc(path, f).Methods("GET")
}

// Post wraps the router for POST method
func (a *App) Post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.rtr.HandleFunc(path, f).Methods("POST")
}

func (a *App) handleGetRequest(handler GetRequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.dbr, w, r)
	}
}

func (a *App) handlePostRequest(handler PostRequestHandlerFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.dbw, w, r)
	}
}
