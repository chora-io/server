package client

import (
	"fmt"

	"github.com/rs/zerolog"
)

type gooseLogger struct {
	log zerolog.Logger
}

func (l gooseLogger) Print(v ...interface{}) {
	l.log.Print(v...)
}

func (l gooseLogger) Printf(format string, v ...interface{}) {
	l.log.Printf(format, v...)
}

func (l gooseLogger) Fatal(v ...interface{}) {
	l.log.Error().Msg(fmt.Sprint(v...))
}

func (l gooseLogger) Fatalf(format string, v ...interface{}) {
	l.log.Error().Msgf(format, v...)
}

func (l gooseLogger) Println(v ...interface{}) {
	l.log.Debug().Msg(fmt.Sprint(v...))
}
