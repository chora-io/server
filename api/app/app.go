package app

import (
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	_ "golang.org/x/crypto/blake2b"

	"github.com/chora-io/server/api/html"

	db "github.com/chora-io/server/db/client"
)

type App struct {
	env string
	dbr db.Reader
	dbw db.Writer
	jsk string
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
		jsk: cfg.JwtSecretKey,
		rtr: mux.NewRouter(),
		aos: cfg.ApiAllowedOrigins,
		log: log,
	}

	app.get("/", app.handleIndexRequest())

	// auth requests
	app.post("/auth", app.handleAuthRequest(PostAuth))
	app.post("/auth/email", app.handleAuthRequest(PostAuthEmail))
	app.post("/auth/keplr", app.handleAuthRequest(PostAuthKeplr))
	app.post("/auth/login", app.handleAuthRequest(PostAuthLogin))

	// data requests
	app.get("/data/{iri}", app.handleGetRequest(GetData))
	app.post("/data", app.handlePostRequest(PostData))

	// indexer requests
	app.get("/idx/{chain_id}/group-proposal/{proposal_id}", app.handleGetRequest(GetIdxGroupProposal))
	app.get("/idx/{chain_id}/group-proposals/{group_id}", app.handleGetRequest(GetIdxGroupProposals))
	app.get("/idx/{chain_id}/group-vote/{proposal_id}/{voter}", app.handleGetRequest(GetIdxGroupVote))
	app.get("/idx/{chain_id}/group-votes/{proposal_id}", app.handleGetRequest(GetIdxGroupVotes))

	return app
}

// Run blocks the current thread of execution and serves the API.
func (a *App) Run(host string) {

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
	bz, err := html.Html.ReadFile("index.html")
	if err != nil {
		log.Fatal(err)
	}

	return func(w http.ResponseWriter, r *http.Request) {
		_, err = io.WriteString(w, string(bz))
		if err != nil {
			log.Fatal(err)
		}
	}
}

func (a *App) handleAuthRequest(handler AuthHandler) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(a.jsk, a.dbr, a.dbw, w, r)
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
