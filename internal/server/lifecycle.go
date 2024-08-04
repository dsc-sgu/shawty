package server

import (
	"context"

	"github.com/dsc-sgu/atcc/internal/database"
	"github.com/dsc-sgu/atcc/internal/log"
)

func onStartup(ctx context.Context) (err error) {
	log.S.Debug("Initializing application state")

	if database.C, err = database.Connect(); err != nil {
		log.S.Errorw("Failed to connect to the PostgreSQL", "error", err)
		return err
	}

	log.S.Debug("Initializing database schema")
	if err = database.C.InitSchema(ctx); err != nil {
		log.S.Errorw("Failed to initialize database schema", "error", err)
		return err
	}
	log.S.Debug("Database schema initialized")

	log.S.Debug("Application state initialized")
	return nil
}

func onShutdown() {
	database.C.Close()
}
