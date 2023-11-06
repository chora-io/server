package client

import (
	"database/sql"
	"io"

	"github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	sqldblogger "github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"

	"github.com/chora-io/server/db/migrations"
)

//go:generate sqlc generate

// Database defines an interface that gives access to Reader and Writer interfaces.
type Database interface {
	Reader() Reader
	Writer() Writer

	io.Closer
}

type db struct {
	db *sql.DB
}

// NewDatabase wraps a postgres database connection and runs goose migrations.
func NewDatabase(dsn string, log zerolog.Logger) (Database, error) {
	sqlDb := sqldblogger.OpenDriver(
		dsn,
		pq.Driver{},
		zerologadapter.New(log),
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
	)

	err := sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	// set logger for goose migrations
	goose.SetLogger(gooseLogger{log})

	// set goose dialect for migrations
	err = goose.SetDialect("postgres")
	if err != nil {
		return nil, err
	}

	// set migrations base directory
	goose.SetBaseFS(migrations.Migrations)

	// run goose migrations
	err = goose.Up(sqlDb, ".")
	if err != nil {
		return nil, err
	}

	d := &db{db: sqlDb}

	return d, nil
}

// Close shuts down the database client.
func (d *db) Close() error {
	return d.db.Close()
}

// Reader returns an interface to read from the database.
func (d *db) Reader() Reader {
	return &reader{q: New(d.db)}
}

// Writer returns an interface to write to the database.
func (d *db) Writer() Writer {
	return &writer{q: New(d.db)}
}
