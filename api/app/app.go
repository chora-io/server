package app

import (
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"

	db "github.com/choraio/server/db/client"
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
func Initialize(cfg Config, log zerolog.Logger) *App {
	db, err := db.NewDatabase(cfg.DatabaseUrl, log)
	if err != nil {
		panic(err)
	}

	app := &App{
		env: cfg.ServerEnv,
		dbr: db.Reader(),
		dbw: db.Writer(),
		rtr: mux.NewRouter(),
		aos: cfg.ApiAllowedOrigins,
		log: log,
	}

	// authenticate
	app.get("/auth", app.handleGetRequest(Auth))

	// data requests
	app.get("/data/{iri}", app.handleGetRequest(GetData))
	app.post("/data/", app.handlePostRequest(PostData))

	// indexer requests
	app.get("/idx/{chain_id}/proposal/{proposal_id}", app.handleGetRequest(GetIdxGroupProposal))
	app.get("/idx/{chain_id}/proposals/{group_id}", app.handleGetRequest(GetIdxGroupProposals))

	return app
}

// Run blocks the current thread of execution and serves the API.
func (a *App) Run(host string) {

	// add handler for static app index request
	a.rtr.HandleFunc("/", a.handleIndexRequest())

	// add allowed origins for get and post requests
	origins := handlers.AllowedOrigins(strings.Split(a.aos, ","))

	// set cors handler with allowed origins
	handler := handlers.CORS(origins)(a.rtr)

	// start listening and serving requests
	if err := http.ListenAndServe(host, handler); err != nil {
		log.Fatal(err)
	}
}

func (a *App) get(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.rtr.HandleFunc(path, f).Methods("GET")
}

func (a *App) post(path string, f func(w http.ResponseWriter, r *http.Request)) {
	a.rtr.HandleFunc(path, f).Methods("POST")
}

func (a *App) handleIndexRequest() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "./html/index.html")
	}
}

func (a *App) handleGetRequest(handler GetHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.dbr, w, r)
	}
}

func (a *App) handlePostRequest(handler PostHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.dbw, w, r)
	}
}
