package context

import (
	"os"
	"time"

	"github.com/choraio/server/idx/config"
	"github.com/rs/zerolog"

	db "github.com/choraio/server/db/client"
)

// Context is the context.
type Context struct {
	config.Config

	BackoffDuration   time.Duration
	BackoffMaxRetries uint64

	db db.Database
}

// NewContext creates a new context.
func NewContext(cfg config.Config) (Context, error) {
	ctx := Context{Config: cfg}
	log := zerolog.New(os.Stdout)

	ctx.BackoffDuration = 1 * time.Second
	ctx.BackoffMaxRetries = 2

	var err error
	ctx.db, err = db.NewDatabase(cfg.DatabaseUrl, log)
	if err != nil {
		return Context{}, err
	}

	return ctx, nil
}

// Close closes the context.
func (b Context) Close() error {
	if b.db != nil {
		err := b.db.Close()
		if err != nil {
			return err
		}
	}

	return nil
}
