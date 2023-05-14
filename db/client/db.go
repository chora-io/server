package client

import (
	"database/sql"
	"fmt"
	"io"

	"github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"
	"github.com/simukti/sqldb-logger"
	"github.com/simukti/sqldb-logger/logadapter/zerologadapter"

	"github.com/choraio/server/db/migrations"
)

//go:generate sqlc generate

// Database defines an interface which gives access to a Reader and Writer.
type Database interface {
	Reader() Reader
	Writer() Writer

	io.Closer
}

type db struct {
	db *sql.DB
}

// NewDatabase wraps a postgres database connection and also runs any needed migrations on startup.
func NewDatabase(postgresUrl string, logger zerolog.Logger) (Database, error) {
	loggerAdapter := zerologadapter.New(logger)
	sqlDb := sqldblogger.OpenDriver(
		postgresUrl,
		pq.Driver{},
		loggerAdapter,
		sqldblogger.WithQueryerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithExecerLevel(sqldblogger.LevelDebug),
		sqldblogger.WithPreparerLevel(sqldblogger.LevelDebug),
	)

	err := sqlDb.Ping()
	if err != nil {
		return nil, err
	}

	if err := goose.SetDialect("postgres"); err != nil {
		return nil, err
	}

	goose.SetBaseFS(migrations.Migrations)
	err = goose.Up(sqlDb, "migrations")

	goose.SetLogger(gooseLogger{logger})
	if err != nil {
		return nil, err
	}

	db := db{db: sqlDb}
	return &db, nil
}

// Writer returns an interface to write to the database.
func (d *db) Writer() Writer {
	return &writer{q: New(d.db)}
}

// Reader returns an interface to read from the database.
func (d *db) Reader() Reader {
	return &reader{q: New(d.db)}
}

func (d *db) Close() error {
	return d.db.Close()
}

type gooseLogger struct {
	logger zerolog.Logger
}

func (g gooseLogger) Print(v ...interface{}) {
	g.logger.Print(v...)
}

func (g gooseLogger) Printf(format string, v ...interface{}) {
	g.logger.Printf(format, v...)
}

func (g gooseLogger) Fatal(v ...interface{}) {
	g.logger.Error().Msg(fmt.Sprint(v...))
}

func (g gooseLogger) Fatalf(format string, v ...interface{}) {
	g.logger.Error().Msgf(format, v...)
}

func (g gooseLogger) Println(v ...interface{}) {
	g.logger.Debug().Msg(fmt.Sprint(v...))
}
