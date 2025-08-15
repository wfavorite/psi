package main

import (
	"log"
	"os"

	"github.com/wfavorite/initq"
)

/* ------------------------------------------------------------------------ */

// Logger defines the multiple (verbosity) destinations for log messages.
type Logger struct {
	Normal  *log.Logger
	Verbose *log.Logger
}

/* ======================================================================== */

// StartLogger sets up the two logger destinations.
func (cd *CoreData) StartLogger() initq.ReqResult {

	if cd.Logr != nil {
		return initq.Satisfied
	}

	logr := new(Logger)

	logr.Normal = log.New(os.Stdout, "", log.LstdFlags)
	logr.Verbose = log.New(os.Stdout, "", log.LstdFlags)

	cd.Logr = logr
	cd.Logr.Verbose.Println("Logger defined.")
	return initq.Satisfied
}
