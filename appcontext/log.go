package appcontext

import (
	"context"
	"os"

	log "github.com/sirupsen/logrus"
)

func SetupLog(ctx context.Context, logLevel string, jsonOutput bool) {
	level, err := log.ParseLevel(logLevel)
	if err != nil {
		level = log.InfoLevel
		log.WithError(err).
			Warning("Unable to ParseLevel, default to info")
	}
	if jsonOutput {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetLevel(level)
	log.SetOutput(os.Stdout)
}
