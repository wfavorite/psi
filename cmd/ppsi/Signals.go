package main

import (
	"os"
	"os/signal"
	"syscall"

	"github.com/wfavorite/initq"
)

/* ------------------------------------------------------------------------ */

// Signals wraps the Signals channel.
type Signals struct {
	s     chan os.Signal `json:"-"`
	Stats map[string]int `json:"stats"`
}

/* ======================================================================== */

// RegisterSignals registers the signal s that we will either ignore or handle.
func (cd *CoreData) RegisterSignals() initq.ReqResult {

	if cd.EvtQ == nil {
		return initq.TryAgain
	}

	sigs := new(Signals)
	sigs.s = make(chan os.Signal, 1)
	sigs.Stats = make(map[string]int)

	// Those to ignore
	signal.Ignore(syscall.SIGHUP)

	// Those to handle
	signal.Notify(sigs.s, syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGXCPU, syscall.SIGUSR1)

	cd.Sigs = sigs

	return initq.Satisfied
}

/* ======================================================================== */

// HandleSignal is called with the signal that is received.
func (cd *CoreData) HandleSignal(sig os.Signal) {

	cd.Sigs.Stats[sig.String()]++

	switch sig {
	case syscall.SIGINT, syscall.SIGQUIT, syscall.SIGTERM:
		cd.NewShutdownEvent()
	case syscall.SIGXCPU:
		cd.Logr.Normal.Println("Shutting down on XCPU signal.")
		cd.NewShutdownEvent()
	case syscall.SIGUSR1:
		cd.NewObserveEvent()
	}

}
