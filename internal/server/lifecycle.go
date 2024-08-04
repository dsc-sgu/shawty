package server

import (
	"github.com/dsc-sgu/atcc/internal/db"
	"github.com/dsc-sgu/atcc/internal/log"
)

func onStartup() error {
	log.S.Debugw("Initializing application state")

	var err error
	db.C, err = db.Connect()
	if err != nil {
		log.S.Errorw("Failed to connect to the PostgreSQL", "error", err)
		return err
	}

	log.S.Debugw("Application state initialized")
	return nil
}

func onShutdown() {
	db.C.Close()
}
